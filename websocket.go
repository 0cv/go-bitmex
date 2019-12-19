package bitmex

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/adampointer/go-bitmex/types"
	evbus "github.com/asaskevich/EventBus"
	"github.com/desertbit/timer"
	"github.com/sacOO7/gowebsocket"
	log "github.com/sirupsen/logrus"
	"go.uber.org/multierr"
)

const heartbeatRate = 5 * time.Second

// EventHandler is a function called in response to websocket events
type EventHandler func(obj interface{})

// APITopicHandler is a function called in response to a API subscription event being received
type APITopicHandler func(response *types.SubscriptionResponse)

// WebSocketClientConfig holds configuration data for the websocket client
type WebsocketClientConfig struct {
	ApiKey, ApiSecret, URL string
}

// NewWebsocketClientConfig returns a new config struct
func NewWebsocketClientConfig() *WebsocketClientConfig {
	return &WebsocketClientConfig{}
}

// WithURL sets the url to use e.g. wss://testnet.bitmex.com/realtime
func (c *WebsocketClientConfig) WithURL(u string) *WebsocketClientConfig {
	c.URL = u
	return c
}

// WithAuth sets the credentials and is optional if you are exclusively using public endpoints
func (c *WebsocketClientConfig) WithAuth(apiKey, apiSecret string) *WebsocketClientConfig {
	c.ApiKey = apiKey
	c.ApiSecret = apiSecret
	return c
}

// WebsocketClient is a BitMEX websocket client
type WebsocketClient struct {
	socket        gowebsocket.Socket
	bus           evbus.Bus
	topics        types.SubscriptionTopics
	conMutex      sync.Mutex
	connected     bool
	config        *WebsocketClientConfig
	heartbeatLock sync.Mutex
	heartbeat     *timer.Timer
}

// NewWebsocketClient returns a new BitMEX websocket client
func NewWebsocketClient(cfg *WebsocketClientConfig) *WebsocketClient {
	return &WebsocketClient{
		socket: gowebsocket.New(cfg.URL),
		bus:    evbus.New(),
		config: cfg,
	}
}

// SetConfig allows you to set the configuration after struct initialisation
func (w *WebsocketClient) SetConfig(cfg *WebsocketClientConfig) {
	w.config = cfg
}

// Start initialises the event bus and connects the websocket
func (w *WebsocketClient) Start() {
	w.setupWebsocketSubscribers()
	w.setupWebsocketPublishers()
	w.socket.Connect()
}

// Restart the websocket connection
func (w *WebsocketClient) Restart() {
	w.socket.Connect()
}

// Shutdown closes the websocket and unsubscribes listeners
func (w *WebsocketClient) Shutdown() {
	w.destroyWebsocketSubscribers()
	defer func() {
		if err := recover(); err != nil {
			log.Warn(err)
		}
	}()
	w.socket.Close()
}

// SendCommand sends command json to the remote
func (w *WebsocketClient) SendCommand(command *types.Command) error {
	msg, err := json.Marshal(command)
	if err != nil {
		return fmt.Errorf("error marshalling command to json: %s", err)
	}
	log.Tracef("===x %s", string(msg))
	w.socket.SendText(string(msg))
	return nil
}

// SubscribeToEvents allows you to register a handler to listen to websocket events
func (w *WebsocketClient) SubscribeToEvents(evt types.Event, handler EventHandler) error {
	return w.bus.SubscribeAsync(evt.String(), handler, false)
}

// SubscribeToOneEvent allows you to register a handler to listen to a single websocket event
// it will be unsubscribed after the first instance automatically
func (w *WebsocketClient) SubscribeToOneEvent(evt types.Event, handler EventHandler) error {
	return w.bus.SubscribeOnceAsync(evt.String(), handler)
}

// UnsubscribeFromEvents unsubscribes a previously subscribed handler
func (w *WebsocketClient) UnsubscribeFromEvents(evt types.Event, handler EventHandler) error {
	return w.bus.Unsubscribe(evt.String(), handler)
}

// SubscribeToApiTopic allows you to subscribe to a BitMEX API topic without message ordering guarantees
func (w *WebsocketClient) SubscribeToApiTopic(topic types.SubscriptionTopic, handler APITopicHandler) error {
	if w.connected {
		err := w.SendCommand(&types.Command{
			Op:   types.CommandOpSubscribe,
			Args: types.CommandArgs{topic.String()},
		})
		if err != nil {
			return err
		}
	} else {
		w.topics = append(w.topics, topic)
	}
	return w.bus.SubscribeAsync(fmt.Sprintf("%s:%s", types.EventApiResponseSubscription, topic.Topic()), handler, true)
}

// UnsubscribeFromApiTopic allows you to subscribe a handler from a previousl subscribed BitMEX API topic
func (w *WebsocketClient) UnsubscribeFromApiTopic(topic types.SubscriptionTopic, handler APITopicHandler) error {
	return w.bus.Unsubscribe(fmt.Sprintf("%s:%s", types.EventApiResponseSubscription, topic.Topic()), handler)
}

func (w *WebsocketClient) setupWebsocketPublishers() {
	w.socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		w.bus.Publish(types.EventWebsocketError.String(), err)
	}
	w.socket.OnConnected = func(_ gowebsocket.Socket) {
		w.bus.Publish(types.EventWebsocketConnected.String(), struct{}{})
	}
	w.socket.OnDisconnected = func(_ error, socket gowebsocket.Socket) {
		w.bus.Publish(types.EventWebsocketDisconnected.String(), struct{}{})
	}
	w.socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		w.heartbeatLock.Lock()
		if w.heartbeat != nil {
			w.heartbeat.Reset(2 * heartbeatRate)
		}
		w.heartbeatLock.Unlock()
		w.bus.Publish(types.EventWebsocketMessage.String(), message)
	}
}

func (w *WebsocketClient) setupWebsocketSubscribers() {
	err := multierr.Combine(
		w.bus.SubscribeAsync(types.EventWebsocketError.String(), w.websocketErrorHandler, false),
		w.bus.SubscribeAsync(types.EventWebsocketDisconnected.String(), w.websocketDisconnectHandler, false),
		w.bus.SubscribeAsync(types.EventWebsocketConnected.String(), w.websocketConnectHandler, false),
		w.bus.SubscribeAsync(types.EventWebsocketMessage.String(), w.websocketMessageHandler, false),
		w.bus.SubscribeAsync(types.EventApiResponseSubscription.String(), w.apiTopicHandler, false),
		w.bus.SubscribeAsync(types.EventApiResponseSuccess.String(), w.apiTopicHandler, false),
	)
	if err != nil {
		log.Fatalf("unable to setup subscribers: %s", err)
	}
}

func (w *WebsocketClient) destroyWebsocketSubscribers() {
	err := multierr.Combine(
		w.bus.Unsubscribe(types.EventWebsocketError.String(), w.websocketErrorHandler),
		w.bus.Unsubscribe(types.EventWebsocketDisconnected.String(), w.websocketDisconnectHandler),
		w.bus.Unsubscribe(types.EventWebsocketConnected.String(), w.websocketConnectHandler),
		w.bus.Unsubscribe(types.EventWebsocketMessage.String(), w.websocketMessageHandler),
		w.bus.Unsubscribe(types.EventApiResponseSubscription.String(), w.apiTopicHandler),
		w.bus.Unsubscribe(types.EventApiResponseSuccess.String(), w.apiTopicHandler),
	)
	if err != nil {
		log.Fatalf("unable to setup subscribers: %s", err)
	}
}

func (w *WebsocketClient) websocketErrorHandler(err error) {
	log.Errorf("websocket error: %s", err)
	w.disconnectAndRetry()
}

func (w *WebsocketClient) websocketConnectHandler(_ struct{}) {
	log.Info("connected")
	if len(w.config.ApiKey) != 0 {
		cmd, err := websocketAuthCommand(w.config.ApiKey, w.config.ApiSecret)
		if err != nil {
			log.Fatal(err)
		}
		log.Info("sent auth command")
		err = w.SendCommand(cmd)
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(w.topics) > 0 {
		err := w.SendCommand(&types.Command{
			Op:   types.CommandOpSubscribe,
			Args: w.topics.Args(),
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	w.conMutex.Lock()
	w.connected = true
	w.conMutex.Unlock()
	w.startHeartbeat()
}

func (w *WebsocketClient) websocketDisconnectHandler(_ struct{}) {
	w.disconnectAndRetry()
}

func (w *WebsocketClient) disconnectAndRetry() {
	log.Warn("disconnected")
	w.heartbeat.Stop()
	w.conMutex.Lock()
	w.connected = false
	w.conMutex.Unlock()
	log.Info("retrying websocket connection in 30s")
	time.Sleep(30 * time.Second)
	w.Restart()
}

func (w *WebsocketClient) websocketMessageHandler(msg string) {
	log.Tracef("x=== %s", msg)
	if msg == "pong" {
		return
	}
	raw := &types.CompositeResponse{}
	if err := json.Unmarshal([]byte(msg), raw); err != nil {
		log.Errorf("error decoding message: %s", err)
		return
	}
	switch {
	case raw.IsAutoCancelResponse():
		w.bus.Publish(types.EventApiResponseAutoCancel.String(), raw.ToAutoCancelResponse().WithRequest(raw.Request))
	case raw.IsErrorResponse():
		w.bus.Publish(types.EventApiResponseError.String(), raw.ToErrorResponse().WithRequest(raw.Request))
	case raw.IsInfoResponse():
		w.bus.Publish(types.EventApiResponseInfo.String(), raw.ToInfoResponse())
	case raw.IsSubscriptionResponse():
		w.bus.Publish(types.EventApiResponseSubscription.String(), raw.ToSubscriptionResponse().WithRequest(raw.Request))
	case raw.IsSuccessResponse():
		w.bus.Publish(types.EventApiResponseSuccess.String(), raw.ToSuccessResponse().WithRequest(raw.Request))
	default:
		log.Errorf("unknown message: %s", msg)
	}
}

func (w *WebsocketClient) apiTopicHandler(obj interface{}) {
	switch o := obj.(type) {
	case *types.SuccessResponse:
		fields := log.Fields{
			"args": o.Request.Args,
		}
		if *o.Success {
			log.WithFields(fields).Infof("%s operation successful", o.Request.Op)
		} else {
			log.WithFields(fields).Infof("%s operation failed'", o.Request.Op)
		}
	case *types.SubscriptionResponse:
		w.bus.Publish(fmt.Sprintf("%s:%s", types.EventApiResponseSubscription, o.Table), o)
	default:
		log.Errorf("error casting apiTopic object")
	}
}

func (w *WebsocketClient) startHeartbeat() {
	log.Info("starting heartbeat")
	w.heartbeatLock.Lock()
	w.heartbeat = timer.NewTimer(2 * heartbeatRate)
	w.heartbeatLock.Unlock()
	ticker := time.NewTicker(heartbeatRate)
	go func() {
		<-w.heartbeat.C
		ticker.Stop()
		log.Warn("no response to heartbeat, restarting websocket")
		w.Restart()
	}()
	go func() {
		for {
			<-ticker.C
			w.socket.SendText("ping")
		}
	}()
}
