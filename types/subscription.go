package types

import "fmt"

type SubscriptionTopic struct {
	topic  string
	filter *SubscriptionFilter
}

func newSubscriptionTopic(topic string) SubscriptionTopic {
	return SubscriptionTopic{topic: topic}
}

func (s SubscriptionTopic) String() string {
	if s.filter != nil {
		if len(s.filter.Symbol) > 0 {
			return fmt.Sprintf("%s:%s", s.topic, s.filter.Symbol)
		}
		if s.filter.Account > 0 {
			return fmt.Sprintf("%s:%d", s.topic, s.filter.Account)
		}
	}
	return s.topic
}

func (s SubscriptionTopic) Topic() string {
	return s.topic
}

func (s SubscriptionTopic) WithInstrument(instrument string) SubscriptionTopic {
	s.filter = &SubscriptionFilter{Symbol: instrument}
	return s
}

func (s SubscriptionTopic) WithAccount(account int64) SubscriptionTopic {
	s.filter = &SubscriptionFilter{Account: account}
	return s
}

var (
	// Public
	SubscriptionTopicAnnouncement        = newSubscriptionTopic("announcement")
	SubscriptionTopicChat                = newSubscriptionTopic("chat")
	SubscriptionTopicConnected           = newSubscriptionTopic("connected")
	SubscriptionTopicFunding             = newSubscriptionTopic("funding")
	SubscriptionTopicInstrument          = newSubscriptionTopic("instrument")
	SubscriptionTopicInsurance           = newSubscriptionTopic("insurance")
	SubscriptionTopicLiquidation         = newSubscriptionTopic("liquidation")
	SubscriptionTopicOrderBookL2_25      = newSubscriptionTopic("orderBookL2_25")
	SubscriptionTopicOrderBookL2         = newSubscriptionTopic("orderBookL2")
	SubscriptionTopicOrderBook10         = newSubscriptionTopic("orderBook10")
	SubscriptionTopicPublicNotifications = newSubscriptionTopic("publicNotifications")
	SubscriptionTopicQuote               = newSubscriptionTopic("quote")
	SubscriptionTopicQuote1m             = newSubscriptionTopic("quoteBin1m")
	SubscriptionTopicQuote5m             = newSubscriptionTopic("quoteBin5m")
	SubscriptionTopicQuote1h             = newSubscriptionTopic("quoteBin1h")
	SubscriptionTopicQuote1d             = newSubscriptionTopic("quoteBin1d")
	SubscriptionTopicSettlement          = newSubscriptionTopic("settlement")
	SubscriptionTopicTrade               = newSubscriptionTopic("trade")
	SubscriptionTopicTrade1m             = newSubscriptionTopic("tradeBin1m")
	SubscriptionTopicTrade5m             = newSubscriptionTopic("tradeBin5m")
	SubscriptionTopicTrade1h             = newSubscriptionTopic("tradeBin1h")
	SubscriptionTopicTrade1d             = newSubscriptionTopic("tradeBin1d")

	// Private
	SubscriptionTopicAffiliate           = newSubscriptionTopic("affiliate")
	SubscriptionTopicExecution           = newSubscriptionTopic("execution")
	SubscriptionTopicOrder               = newSubscriptionTopic("order")
	SubscriptionTopicMargin              = newSubscriptionTopic("margin")
	SubscriptionTopicPosition            = newSubscriptionTopic("position")
	SubscriptionTopicPrivateNotification = newSubscriptionTopic("privateNotifications")
	SubscriptionTopicTransact            = newSubscriptionTopic("transact")
	SubscriptionTopicWallet              = newSubscriptionTopic("wallet")
)

type SubscriptionTopics []SubscriptionTopic

func (s SubscriptionTopics) Args() CommandArgs {
	out := make([]interface{}, len(s))
	for i, j := range s {
		out[i] = j.String()
	}
	return out
}
