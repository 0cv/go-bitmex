package bitmex

import (
	"bytes"
	"encoding"
	"errors"
	"fmt"
	"github.com/go-openapi/runtime"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/adampointer/go-bitmex/swagger/client"
	"github.com/avast/retry-go"
	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

var defaultRetryOpts = []retry.Option{retry.Delay(500 * time.Millisecond), retry.Attempts(10)}

func customTextConsumer(defaultTextConsumer runtime.Consumer) runtime.Consumer {
	return runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		if d, ok := data.(encoding.TextUnmarshaler); ok {
			return defaultTextConsumer.Consume(reader, d)
		}
		return nil
	})
}

// NewBitmexClient creates a new REST client
func NewBitmexClient(config *ClientConfig) *client.BitMEX {
	transportConfig := client.DefaultTransportConfig().
		WithHost(config.HostUrl.Host).
		WithBasePath(client.DefaultBasePath).
		WithSchemes([]string{config.HostUrl.Scheme})
	httpclient := newHttpClient(config)
	transport := rc.NewWithClient(transportConfig.Host, transportConfig.BasePath, transportConfig.Schemes, httpclient)
	transport.Consumers[runtime.HTMLMime] = customTextConsumer(transport.Consumers[runtime.HTMLMime])
	transport.Consumers[runtime.TextMime] = customTextConsumer(transport.Consumers[runtime.TextMime])
	transport.Debug = config.VerboseRequestLogging
	return client.New(transport, strfmt.Default)
}

// ClientConfig holds configuration data for the REST client
// Rather than using this directly you should generally use the NewClientConfig
// function and the builder functions
type ClientConfig struct {
	HostUrl               *url.URL
	underlyingTransport   http.RoundTripper
	ApiKey, ApiSecret     string
	VerboseRequestLogging bool
	RetryOpts             []retry.Option
}

// NewClientConfig returns a *ClientConfig with the default transport set and the default retry options
// Default retry is exponential backoff, 10 attempts with an initial delay of 500 milliseconds
func NewClientConfig() *ClientConfig {
	return &ClientConfig{
		underlyingTransport: http.DefaultTransport,
		RetryOpts:           defaultRetryOpts,
	}
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

// WithRetryOptions sets the request retry options, replacing the defaults
func (c *ClientConfig) WithRetryOptions(opts ...retry.Option) *ClientConfig {
	c.RetryOpts = opts
	return c
}

// WithRetryOption appends a retry option to the defaults
func (c *ClientConfig) WithRetryOption(opt retry.Option) *ClientConfig {
	c.RetryOpts = append(c.RetryOpts, opt)
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
	var res *http.Response
	err := retry.Do(func() error {
		var err error
		res, err = t.underlyingTransport.RoundTrip(req)
		if err != nil {
			return err
		}
		if res.StatusCode == 429 {
			log.Fatal("rate limiting - shutting down")
			return retry.Unrecoverable(errors.New("rate limiting"))
		}
		if res.StatusCode == 503 {
			return errors.New("http status 503")
		}
		return nil
	}, t.config.RetryOpts...)

	log.WithFields(log.Fields{
		"limit":     res.Header.Get("x-ratelimit-limit"),
		"remaining": res.Header.Get("x-ratelimit-remaining"),
		"reset":     res.Header.Get("x-ratelimit-reset"),
	}).Debugf("rate limits")

	return res, err
}

func newHttpClient(config *ClientConfig) *http.Client {
	transport := &transport{underlyingTransport: config.underlyingTransport, config: config}
	h := &http.Client{Transport: transport}
	return h
}
