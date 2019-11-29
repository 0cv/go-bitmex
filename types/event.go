package types

type Event string

func (e Event) String() string {
	return string(e)
}

const (
	EventWebsocketError          Event = "websocket:error"
	EventWebsocketConnected      Event = "websocket:connected"
	EventWebsocketDisconnected   Event = "websocket:disconnected"
	EventWebsocketMessage        Event = "websocket:message"
	EventApiResponseAutoCancel   Event = "api:response:autocancel"
	EventApiResponseError        Event = "api:response:error"
	EventApiResponseInfo         Event = "api:response:info"
	EventApiResponseSubscription Event = "api:response:subscription"
	EventApiResponseSuccess      Event = "api:response:success"
)
