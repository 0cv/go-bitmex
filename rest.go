package bitmex

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/adampointer/go-bitmex/swagger/client"
	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

// NewBitmexClient creates a new REST client
func NewBitmexClient(config *ClientConfig) *client.BitMEX {
	transportConfig := client.DefaultTransportConfig().
		WithHost(config.HostUrl.Host).
		WithBasePath(client.DefaultBasePath).
		WithSchemes([]string{config.HostUrl.Scheme})
	httpclient := newHttpClient(config)
	transport := rc.NewWithClient(transportConfig.Host, transportConfig.BasePath, transportConfig.Schemes, httpclient)
	transport.Debug = config.Debug
	return client.New(transport, strfmt.Default)
}

// ClientConfig holds configuration data for the REST client
// Rather than using this directly you should generally use the NewClientConfig
// function and the builder functions
type ClientConfig struct {
	HostUrl             *url.URL
	underlyingTransport http.RoundTripper
	ApiKey, ApiSecret   string
	Debug               bool
}

// NewClientConfig returns a *ClientConfig with the default transport set
func NewClientConfig() *ClientConfig {
	return &ClientConfig{underlyingTransport: http.DefaultTransport}
}

// WithURL sets the url to use e.g. https://testnet.bitmex.com
func (c *ClientConfig) WithURL(u string) *ClientConfig {
	hostUrl, err := url.Parse(u)
	if err != nil {
		log.Fatalf("cannot parse url: %s", err)
	}
	c.HostUrl = hostUrl
	return c
}

// WithAuth sets the credentials and is optional if you are exclusively using public endpoints
func (c *ClientConfig) WithAuth(apiKey, apiSecret string) *ClientConfig {
	c.ApiKey = apiKey
	c.ApiSecret = apiSecret
	return c
}

// WithTransport allows you to override the underlying transport used by the custom RoundTripper
func (c *ClientConfig) WithTransport(t http.RoundTripper) *ClientConfig {
	c.underlyingTransport = t
	return c
}

type transport struct {
	config              *ClientConfig
	underlyingTransport http.RoundTripper
}

// RoundTrip implements http.RoundTripper for transport
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(t.config.ApiKey) != 0 {
		path := req.URL.Path
		if len(req.URL.Query()) > 0 {
			path = fmt.Sprintf("%s?%s", path, req.URL.RawQuery)
		}
		var body []byte
		var err error
		if req.Body != nil {
			body, err = ioutil.ReadAll(req.Body)
			if err != nil {
				return nil, err
			}
			req.Body = ioutil.NopCloser(bytes.NewReader(body))
		}
		params := &sigParams{
			method:  req.Method,
			path:    path,
			secret:  t.config.ApiSecret,
			body:    string(body),
			expires: expiryTime(),
		}
		sig, err := calculateSignature(params)
		if err != nil {
			return nil, err
		}
		req.Header.Add("api-expires", params.expiryString())
		req.Header.Add("api-key", t.config.ApiKey)
		req.Header.Add("api-signature", sig)
	}
	res, err := t.underlyingTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == 429 {
		log.Fatal("rate limiting - shutting down")
	}
	return res, nil
}

func newHttpClient(config *ClientConfig) *http.Client {
	transport := &transport{underlyingTransport: config.underlyingTransport, config: config}
	h := &http.Client{Transport: transport}
	return h
}
