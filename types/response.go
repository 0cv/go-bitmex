package types

import (
	"encoding/json"
	"time"

	"github.com/adampointer/go-bitmex/swagger/models"
)

// CompositeResponse is what all incoming API messages are unmarshaled to
type CompositeResponse struct {
	SubscriptionResponse
	AutoCancelResponse
	ErrorResponse
	SuccessResponse
	InfoResponse
	Request *Command `json:"request"`
}

// IsAutoCancelResponse returns true if the message is an AutoCancelResponse
func (c *CompositeResponse) IsAutoCancelResponse() bool {
	return c.Request != nil && c.Request.Op == CommandOpCancelAllAfter
}

// ToAutoCancelResponse returns the AutoCancelResponse from the CompositeResponse
func (c *CompositeResponse) ToAutoCancelResponse() *AutoCancelResponse {
	return &AutoCancelResponse{
		Now:        c.Now,
		CancelTime: c.CancelTime,
	}
}

// IsErrorResponse returns true if the message is an ErrorResponse
func (c *CompositeResponse) IsErrorResponse() bool {
	return len(c.Error) > 0
}

// ToErrorResponse returns the AutoCancelResponse from the CompositeResponse
func (c *CompositeResponse) ToErrorResponse() *ErrorResponse {
	return &ErrorResponse{Error: c.Error}
}

// IsSuccessResponse returns true if the message is a SuccessResponse
func (c *CompositeResponse) IsSuccessResponse() bool {
	return c.Success != nil
}

// ToSuccessResponse returns the SuccessResponse from the CompositeResponse
func (c *CompositeResponse) ToSuccessResponse() *SuccessResponse {
	return &SuccessResponse{
		Subscribe: c.Subscribe,
		Success:   c.Success,
	}
}

// IsInfoResponse returns true if the message is an InfoResponse
func (c *CompositeResponse) IsInfoResponse() bool {
	return len(c.Info) > 0
}

// ToInfoResponse returns the InfoResponse from the CompositeResponse
func (c *CompositeResponse) ToInfoResponse() *InfoResponse {
	return &InfoResponse{
		Info:      c.Info,
		Version:   c.Version,
		Timestamp: c.Timestamp,
		Docs:      c.Docs,
		Limit:     c.Limit,
	}
}

// IsSubscriptionResponse returns true if the message is a SubscriptionResponse
func (c *CompositeResponse) IsSubscriptionResponse() bool {
	return len(c.Table) > 0
}

// ToSubscriptionResponse returns the SubscriptionResponse from the CompositeResponse
func (c *CompositeResponse) ToSubscriptionResponse() *SubscriptionResponse {
	return &SubscriptionResponse{
		Table:       c.Table,
		Keys:        c.Keys,
		Types:       c.Types,
		ForeignKeys: c.ForeignKeys,
		Attributes:  c.Attributes,
		Action:      c.Action,
		Data:        c.Data,
		Filter:      c.Filter,
	}
}

// SubscriptionResponse is a response to a subscription command
type SubscriptionResponse struct {
	Table       string              `json:"table"`
	Keys        []string            `json:"keys"`
	Types       map[string]string   `json:"types"`
	ForeignKeys map[string]string   `json:"foreignKeys"`
	Attributes  map[string]string   `json:"attributes"`
	Action      string              `json:"action"`
	Data        json.RawMessage     `json:"data"`
	Filter      *SubscriptionFilter `json:"filter"`
	Request     *Command            `json:"-"`
}

// WithRequest allows you to attach the original command to the response
func (s *SubscriptionResponse) WithRequest(r *Command) *SubscriptionResponse {
	s.Request = r
	return s
}

// AffiliateData attempts to unmarshal the SubscriptionResponse to an Affiliate
func (s *SubscriptionResponse) AffiliateData() ([]*models.Affiliate, error) {
	var out []*models.Affiliate
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// AnnouncementData attempts to unmarshal the SubscriptionResponse to an Announcement
func (s *SubscriptionResponse) AnnouncementData() ([]*models.Announcement, error) {
	var out []*models.Announcement
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ChatData attempts to unmarshal the SubscriptionResponse to a Chat
func (s *SubscriptionResponse) ChatData() ([]*models.Chat, error) {
	var out []*models.Chat
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectedUsersData attempts to unmarshal the SubscriptionResponse to a ConnectedUsers
func (s *SubscriptionResponse) ConnectedUsersData() ([]*models.ConnectedUsers, error) {
	var out []*models.ConnectedUsers
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutionData attempts to unmarshal the SubscriptionResponse to an Execution
func (s *SubscriptionResponse) ExecutionData() ([]*models.Execution, error) {
	var out []*models.Execution
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// FundingData attempts to unmarshal the SubscriptionResponse to a Funding
func (s *SubscriptionResponse) FundingData() ([]*models.Funding, error) {
	var out []*models.Funding
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// InstrumentData attempts to unmarshal the SubscriptionResponse to an Instrument
func (s *SubscriptionResponse) InstrumentData() ([]*models.Instrument, error) {
	var out []*models.Instrument
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// InsuranceData attempts to unmarshal the SubscriptionResponse to an Insurance
func (s *SubscriptionResponse) InsuranceData() ([]*models.Insurance, error) {
	var out []*models.Insurance
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// LiquidationData attempts to unmarshal the SubscriptionResponse to a Liquidation
func (s *SubscriptionResponse) LiquidationData() ([]*models.Liquidation, error) {
	var out []*models.Liquidation
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// MarginData attempts to unmarshal the SubscriptionResponse to a Margin
func (s *SubscriptionResponse) MarginData() ([]*models.Margin, error) {
	var out []*models.Margin
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OrderData attempts to unmarshal the SubscriptionResponse to an Order
func (s *SubscriptionResponse) OrderData() ([]*models.Order, error) {
	var out []*models.Order
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// OrderBookL2Data attempts to unmarshal the SubscriptionResponse to an OrderBookL2
func (s *SubscriptionResponse) OrderBookL2Data() ([]*models.OrderBookL2, error) {
	var out []*models.OrderBookL2
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// PositionData attempts to unmarshal the SubscriptionResponse to a Position
func (s *SubscriptionResponse) PositionData() ([]*models.Position, error) {
	var out []*models.Position
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationData attempts to unmarshal the SubscriptionResponse to a Notification
func (s *SubscriptionResponse) NotificationData() ([]*models.Notification, error) {
	var out []*models.Notification
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// QuoteData attempts to unmarshal the SubscriptionResponse to a Quote
func (s *SubscriptionResponse) QuoteData() ([]*models.Quote, error) {
	var out []*models.Quote
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SettlementData attempts to unmarshal the SubscriptionResponse to a Settlement
func (s *SubscriptionResponse) SettlementData() ([]*models.Settlement, error) {
	var out []*models.Settlement
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// TradeData attempts to unmarshal the SubscriptionResponse to a Trade
func (s *SubscriptionResponse) TradeData() ([]*models.Trade, error) {
	var out []*models.Trade
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// TradeBinData attempts to unmarshal the SubscriptionResponse to a TradeBin
func (s *SubscriptionResponse) TradeBinData() ([]*models.TradeBin, error) {
	var out []*models.TradeBin
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionData attempts to unmarshal the SubscriptionResponse to a Transaction
func (s *SubscriptionResponse) TransactionData() ([]*models.Transaction, error) {
	var out []*models.Transaction
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletData attempts to unmarshal the SubscriptionResponse to a Wallet
func (s *SubscriptionResponse) WalletData() ([]*models.Wallet, error) {
	var out []*models.Wallet
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionFilter allows you to add a filter to a subscription request
type SubscriptionFilter struct {
	Symbol  string `json:"symbol"`
	Account int64  `json:"account"`
}

// AutoCancelResponse is a response to CommandOpCancelAllAfter
type AutoCancelResponse struct {
	Now        time.Time `json:"now"`
	CancelTime time.Time `json:"cancelTime"`
	Request    *Command  `json:"-"`
}

// WithRequest allows you to attach the original command to the response
func (a *AutoCancelResponse) WithRequest(r *Command) *AutoCancelResponse {
	a.Request = r
	return a
}

// ErrorResponse is an API error message
type ErrorResponse struct {
	Error   string   `json:"error"`
	Request *Command `json:"-"`
}

// WithRequest allows you to attach the original command to the response
func (e *ErrorResponse) WithRequest(r *Command) *ErrorResponse {
	e.Request = r
	return e
}

// SuccessResponse is a success message responding to a subscribe command
type SuccessResponse struct {
	Subscribe string   `json:"subscribe"`
	Success   *bool    `json:"success"`
	Request   *Command `json:"-"`
}

// WithRequest allows you to attach the original command to the response
func (s *SuccessResponse) WithRequest(r *Command) *SuccessResponse {
	s.Request = r
	return s
}

// InfoResponse is an information message sent by the remote when you first connect
type InfoResponse struct {
	Info      string     `json:"info"`
	Version   string     `json:"version"`
	Timestamp time.Time  `json:"timestamp"`
	Docs      string     `json:"docs"`
	Limit     *LimitInfo `json:"limit"`
}

// LimitInfo informs about rate limits
type LimitInfo struct {
	Remaining int `json:"remaining"`
}
