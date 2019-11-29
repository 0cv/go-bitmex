[![Build Status](https://travis-ci.com/adampointer/go-bitmex.svg?branch=master)](https://travis-ci.com/adampointer/go-bitmex)  [![Go Report Card](https://goreportcard.com/badge/github.com/adampointer/go-bitmex)](https://goreportcard.com/report/github.com/adampointer/go-bitmex) 

# go-bitmex

Event driven BitMEX client library for Golang which provides a full access to both the REST and websocket APIs.

Authentication is handled automatically for you by both clients. If there is credentials passed it will authenticate.
Credentials are optional and if you are only using public endpoints, then there is no need to supply them.

## Further Documentation

Read the tests [rest_test.go](https://github.com/adampointer/go-bitmex/blob/master/rest_test.go) and [websocket_test.go](https://github.com/adampointer/go-bitmex/blob/master/websocket_test.go)

GoDoc package documentation is [here](https://godoc.org/github.com/adampointer/go-bitmex)

## Using the REST client

This is a basic go-swagger client generated from the official swagger documents provided by BitMEX with a custom
`http.RoundTripper` to calculate and inject the authorisation signature into the headers.

### Example

```go
package main

import (
    "fmt"

    "github.com/adampointer/go-bitmex"
    "github.com/adampointer/go-bitmex/swagger/client/trade"
    "github.com/wacul/ptr"
)

func main() {
    cfg := bitmex.NewClientConfig().
        WithURL("https://testnet.bitmex.com").
        WithAuth("api_key", "api_secret") // Not required for public endpoints
    restClient := bitmex.NewBitmexClient(cfg)

    params := trade.NewTradeGetBucketedParams()
    params.SetSymbol(ptr.String("XBTUSD"))
    params.SetBinSize(ptr.String("1d"))
    params.SetReverse(ptr.Bool(true))

    res, err := restClient.Trade.TradeGetBucketed(params)
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}
```

## WebSocket Client

The websocket implementation is slightly more involved. The core of it is an event driven websocket library. Events/mesages 
are then shipped out to an event bus where clients can subscribe to specific topics. 

There are two types of subscriptions available. Websocket events, such as connect, disconnect etc. and API subscriptions
where parsed messages are pushed. API subscriptions can be filtered by instrument.

Both kind of subscription can be added/removed at anytime either before `Start()` is called or after.

### Websocket Events

```go
err := wsClient.SubscribeToEvents(types.EventWebsocketConnected, func(interface{}) {
    fmt.Println("connected")
})
```

### Subscriptions

```go
// Subscribe to all orders
wsClient.SubscribeToApiTopic(types.SubscriptionTopicOrder, func(res *types.SubscriptionResponse) {
    // Decode the response into an OrderData model
    data, err := res.OrderData()
    if err != nil {
        log.Errorf("error getting order data: %s", err)
	}
    fmt.Println(data)
})
// API subscription with instrument filter
wsClient.SubscribeToApiTopic(types.SubscriptionTopicOrderBookL2.WithInstrument("XRPZ19"), func(res *types.SubscriptionResponse) {
    data, err := res.OrderBookL2Data()
    if err != nil {
        log.Errorf("error getting orderbook data: %s", err)
    }
    fmt.Println(data)
})
```

### Full Example

```go
package main

import (
    "log"

    "github.com/adampointer/go-bitmex"
    "github.com/adampointer/go-bitmex/types"
)

func main() {
    cfg := bitmex.NewWebsocketClientConfig().
        WithURL("wss://testnet.bitmex.com/realtime").
        WithAuth("api_key", "api_secret") // Not required for public endpoints
    wsClient := bitmex.NewWebsocketClient(cfg)

    err := wsClient.SubscribeToApiTopic(types.SubscriptionTopicOrderBookL2.WithInstrument("XRPZ19"), func(res *types.SubscriptionResponse) {
        data, err := res.OrderBookL2Data()
        if err != nil {
            log.Fatalf("error getting orderbook data: %s", err)
        }
        log.Println(data)
    })
    if err != nil {
        panic(err)
    }
    wsClient.Start()
}
```

### Heartbeat and Reconnection

Due to an issue in the underlying websocket library, the disconnect handler may not get called if the connection is silently
dropped by the remote. Heartbeat is implemented and pings the remote every 5 seconds. A 10 second timer is reset on every message 
received. If the timer expires, the library will automatically call `Restart()`. On a successful reconnect, any configured
API subscription will be re-subscribed.
