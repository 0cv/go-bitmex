package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/adampointer/go-bitmex"
	"github.com/adampointer/go-bitmex/book"
	log "github.com/sirupsen/logrus"
)

func main() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	ws := bitmex.NewWebsocketClient(&bitmex.WebsocketClientConfig{
		ApiKey:    os.Getenv("BITMEX_API_KEY"),
		ApiSecret: os.Getenv("BITMEX_API_SECRET"),
		URL:       "wss://www.bitmex.com/realtime",
	})

	l2, err := ws.L2OrderBook("XBTUSD", 0.25)
	if err != nil {
		log.WithError(err).Fatal("error creating l2 order book")
	}

	err = l2.SubscribeToTop5Quotes(false, func(q *book.QuoteData) {
		log.Infof("Best Bid %.2f (%d) Best Ask %.2f (%d) MidPrice %.2f", q.TopBid, q.TopBids[0].Size, q.TopAsk, q.TopAsks[0].Size, q.MidPrice)
	})
	if err != nil {
		log.WithError(err).Fatal("error creating subscription")
	}

	ws.Start()

	// Trap kill signals
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-terminate
}
