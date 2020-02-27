package book

import (
	"fmt"
	"math"
	"sort"
	"sync"

	"github.com/adampointer/go-bitmex/swagger/models"
	"github.com/adampointer/go-bitmex/types"
	evbus "github.com/asaskevich/EventBus"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type QuoteData struct {
	TopAsk, TopBid, MidPrice, Spread float64
	TopAsks, TopBids                 []*models.OrderBookL2
}

type orderBookLevels []*models.OrderBookL2

func (o orderBookLevels) Len() int {
	return len(o)
}

func (o orderBookLevels) Less(i, j int) bool {
	return o[i].Price < o[j].Price || math.IsNaN(o[i].Price) && math.IsNaN(o[j].Price)
}

func (o orderBookLevels) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

const topQuoteTopic = "quote"

type QuoteHandler func(d *QuoteData)

type Book struct {
	lock                     *sync.RWMutex
	Asks, Bids               map[int64]*models.OrderBookL2
	initialised, dirty       bool
	bus                      evbus.Bus
	topAsk, topBid, tickSize float64
	log                      *log.Logger
	topAsks, topBids         []*models.OrderBookL2
	rawQuoteHandler          QuoteHandler
}

func NewBook(bus evbus.Bus, ts float64, l *log.Logger) *Book {
	return &Book{
		Asks:     make(map[int64]*models.OrderBookL2),
		Bids:     make(map[int64]*models.OrderBookL2),
		lock:     &sync.RWMutex{},
		bus:      bus,
		tickSize: ts,
		log:      l,
	}
}

func (b *Book) UpdateHandler(r *types.SubscriptionResponse) error {
	b.dirty = false
	data, err := r.OrderBookL2Data()
	if err != nil {
		return errors.Wrap(err, "error decoding l2 book data")
	}

	var oldTopAsk, oldTopBid float64

	b.lock.Lock()
	defer b.lock.Unlock()

	oldTopAsk = b.topAsk
	oldTopBid = b.topBid

	switch r.Action {
	case "partial":
		err = b.init(data)
	case "update":
		err = b.update(data)
	case "delete":
		err = b.delete(data)
	case "insert":
		err = b.create(data)
	default:
		err = fmt.Errorf("unknown action received '%s'", r.Action)
	}
	if err != nil {
		return err
	}
	var notify bool
	if b.dirty {
		notify = b.sortPrices()
	}

	// Sanity checking
	if b.topBid == 0 && b.topAsk == 0 {
		return nil
	}
	if b.topBid == 0 && b.topAsk != 0 {
		b.topBid = b.topAsk - b.tickSize
	}
	if b.topAsk == 0 && b.topBid != 0 {
		b.topAsk = b.topBid + b.tickSize
	}
	if b.topBid >= b.topAsk {
		if b.topBid == oldTopBid {
			b.topBid = b.topAsk - b.tickSize
		} else if b.topAsk == oldTopAsk {
			b.topAsk = b.topBid + b.tickSize
		}
	}

	// Notify
	q := &QuoteData{
		TopAsk:   b.topAsk,
		TopBid:   b.topBid,
		MidPrice: (b.topAsk + b.topBid) / 2,
		TopAsks:  b.topAsks,
		TopBids:  b.topBids,
		Spread:   b.topAsk - b.topBid,
	}
	if notify {
		b.rawQuoteHandler(q)
	}

	if b.topAsk != oldTopAsk || b.topBid != oldTopBid && b.bus != nil {
		b.bus.Publish(topQuoteTopic, q)
	}

	return nil
}

func (b *Book) TopAsk() float64 {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.topAsk
}

func (b *Book) TopBid() float64 {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.topBid
}

// SubscribeToTopQuotes subscribes to a topic which is updated when the top bid or ask price changes
func (b *Book) SubscribeToTopQuotes(once bool, f QuoteHandler) error {
	b.log.Debug("subscribe quotes")
	if once {
		return b.bus.SubscribeOnceAsync(topQuoteTopic, f)
	}
	return b.bus.SubscribeAsync(topQuoteTopic, f, true)
}

func (b *Book) UnsubscribeFromTopQuotes(f QuoteHandler) error {
	b.log.Debug("unsubscribe quotes")
	return b.bus.Unsubscribe(topQuoteTopic, f)
}

// SubscribeToTop5Quotes subscribes to a topic which is updated when there is a change
// in price or size of the top 5 bids or asks
func (b *Book) SubscribeToTop5Quotes(_ bool, f QuoteHandler) error {
	b.rawQuoteHandler = f
	return nil
}

func (b *Book) UnsubscribeFromTop5Quotes(_ QuoteHandler) error {
	return nil
}

func (b *Book) init(table []*models.OrderBookL2) error {
	if err := b.addRows(table); err != nil {
		return err
	}
	b.initialised = true
	return nil
}

func (b *Book) create(table []*models.OrderBookL2) error {
	if !b.initialised {
		return nil
	}
	if err := b.addRows(table); err != nil {
		return err
	}
	return nil
}

func (b *Book) update(table []*models.OrderBookL2) error {
	if !b.initialised {
		return nil
	}
	if err := b.updateRows(table); err != nil {
		return err
	}
	return nil
}

func (b *Book) delete(table []*models.OrderBookL2) error {
	if !b.initialised {
		return nil
	}
	if err := b.deleteRows(table); err != nil {
		return err
	}
	return nil
}

func (b *Book) addRows(table []*models.OrderBookL2) error {
	b.dirty = true
	for _, row := range table {
		switch *row.Side {
		case "Buy":
			b.Bids[*row.ID] = row
		case "Sell":
			b.Asks[*row.ID] = row
		default:
			return fmt.Errorf("unknown side parameter in row '%s", *row.Side)
		}
	}
	return nil
}

func (b *Book) deleteRows(table []*models.OrderBookL2) error {
	b.dirty = true
	for _, row := range table {
		switch *row.Side {
		case "Buy":
			delete(b.Bids, *row.ID)
		case "Sell":
			delete(b.Asks, *row.ID)
		}
	}
	return nil
}

func (b *Book) updateRows(table []*models.OrderBookL2) error {
	for _, row := range table {
		var model *models.OrderBookL2
		switch *row.Side {
		case "Buy":
			if m, ok := b.Bids[*row.ID]; ok {
				model = m
			}
		case "Sell":
			if m, ok := b.Asks[*row.ID]; ok {
				model = m
			}
		}
		if model == nil {
			b.log.Warnf("update for unknown row: %d %s", *row.ID, *row.Side)
			if err := b.addRows([]*models.OrderBookL2{row}); err != nil {
				b.log.WithError(err).Error("error inserting row")
			}
			continue
		}
		if row.Price != model.Price && row.Price != 0 {
			b.dirty = true
			model.Price = row.Price
		}
		if row.Size != model.Size && row.Size != 0 {
			model.Size = row.Size
		}
	}
	return nil
}

func (b *Book) sortPrices() (dirty bool) {
	if len(b.Asks) == 0 || len(b.Bids) == 0 {
		return
	}
	levels := make(orderBookLevels, len(b.Asks)+len(b.Bids))
	var i int
	for _, j := range b.Bids {
		levels[i] = j
		i++
	}
	for _, j := range b.Asks {
		levels[i] = j
		i++
	}
	sort.Sort(levels)

	topBids := levels[len(b.Bids)-5 : len(b.Bids)]
	topBids[0], topBids[1], topBids[3], topBids[4] = topBids[4], topBids[3], topBids[1], topBids[0]
	topAsks := levels[len(b.Bids) : len(b.Bids)+5]

	b.topBid = topBids[0].Price
	b.topAsk = topAsks[0].Price

	if !isLevelEqual(b.topBids, topBids) || !isLevelEqual(b.topAsks, topAsks) {
		dirty = true
	}

	b.topBids = makeDeepCopy(topBids)
	b.topAsks = makeDeepCopy(topAsks)
	return
}

func isLevelEqual(a, b []*models.OrderBookL2) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i].Size != b[i].Size {
			return false
		}
		if a[i].Price != b[i].Price {
			return false
		}
		if *a[i].ID != *b[i].ID {
			return false
		}
		if *a[i].Side != *b[i].Side {
			return false
		}
	}
	return true
}

func makeDeepCopy(a []*models.OrderBookL2) []*models.OrderBookL2 {
	b := make([]*models.OrderBookL2, len(a))
	for i, j := range a {
		id := *j.ID
		side := *j.Side
		symbol := *j.Symbol
		b[i] = &models.OrderBookL2{
			ID:     &id,
			Price:  j.Price,
			Side:   &side,
			Size:   j.Size,
			Symbol: &symbol,
		}
	}
	return b
}
