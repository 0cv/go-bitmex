package types

// Event is a topic which can be subscribed to on the local event bus
type Event string

// String implements Stringer
func (e Event) String() string {
	return string(e)
}

const (
	// EventWebsocketError is a topic for connection errors
	EventWebsocketError Event = "websocket:error"
	// EventWebsocketConnected is fired when the websocket connection is established
	EventWebsocketConnected Event = "websocket:connected"
	// EventWebsocketDisconnected is fired when the websocket connection is terminated
	EventWebsocketDisconnected Event = "websocket:disconnected"
	// EventWebsocketMessage is a topic for incoming message
	EventWebsocketMessage Event = "websocket:message"
	// EventApiResponseAutoCancel is a response to a CommandOpCancelAllAfter command
	EventApiResponseAutoCancel Event = "api:response:autocancel"
	// EventApiResponseError is a topic for API errors as opposed to low level websocket errors
	EventApiResponseError Event = "api:response:error"
	// EventApiResponseInfo is a topic for info messages
	EventApiResponseInfo Event = "api:response:info"
	// EventApiResponseSubscription is a topic where all subscribed messages are published
	EventApiResponseSubscription Event = "api:response:subscription"
	// EventApiResponseSuccess is a response to a CommandOpSubscribe command
	EventApiResponseSuccess Event = "api:response:success"
)
