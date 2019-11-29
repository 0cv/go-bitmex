package types

import (
	"encoding/json"
	"time"

	"github.com/adampointer/go-bitmex/swagger/models"
)

type CompositeResponse struct {
	SubscriptionResponse
	AutoCancelResponse
	ErrorResponse
	SuccessResponse
	InfoResponse
	Request *Command `json:"request"`
}

func (c *CompositeResponse) IsAutoCancelResponse() bool {
	return c.Request != nil && c.Request.Op == CommandOpCancelAllAfter
}

func (c *CompositeResponse) ToAutoCancelResponse() *AutoCancelResponse {
	return &AutoCancelResponse{
		Now:        c.Now,
		CancelTime: c.CancelTime,
	}
}

func (c *CompositeResponse) IsErrorResponse() bool {
	return len(c.Error) > 0
}

func (c *CompositeResponse) ToErrorResponse() *ErrorResponse {
	return &ErrorResponse{Error: c.Error}
}

func (c *CompositeResponse) IsSuccessResponse() bool {
	return c.Success != nil
}

func (c *CompositeResponse) ToSuccessResponse() *SuccessResponse {
	return &SuccessResponse{
		Subscribe: c.Subscribe,
		Success:   c.Success,
	}
}

func (c *CompositeResponse) IsInfoResponse() bool {
	return len(c.Info) > 0
}

func (c *CompositeResponse) ToInfoResponse() *InfoResponse {
	return &InfoResponse{
		Info:      c.Info,
		Version:   c.Version,
		Timestamp: c.Timestamp,
		Docs:      c.Docs,
		Limit:     c.Limit,
	}
}

func (c *CompositeResponse) IsSubscriptionResponse() bool {
	return len(c.Table) > 0
}

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

func (s *SubscriptionResponse) WithRequest(r *Command) *SubscriptionResponse {
	s.Request = r
	return s
}

func (s *SubscriptionResponse) AffiliateData() ([]*models.Affiliate, error) {
	var out []*models.Affiliate
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) AnnouncementData() ([]*models.Announcement, error) {
	var out []*models.Announcement
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) ChatData() ([]*models.Chat, error) {
	var out []*models.Chat
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) ConnectedUsersData() ([]*models.ConnectedUsers, error) {
	var out []*models.ConnectedUsers
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) ExecutionData() ([]*models.Execution, error) {
	var out []*models.Execution
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) FundingData() ([]*models.Funding, error) {
	var out []*models.Funding
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) InstrumentData() ([]*models.Instrument, error) {
	var out []*models.Instrument
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) InsuranceData() ([]*models.Insurance, error) {
	var out []*models.Insurance
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) LiquidationData() ([]*models.Liquidation, error) {
	var out []*models.Liquidation
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) MarginData() ([]*models.Margin, error) {
	var out []*models.Margin
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) OrderData() ([]*models.Order, error) {
	var out []*models.Order
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) OrderBookL2Data() ([]*models.OrderBookL2, error) {
	var out []*models.OrderBookL2
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) PositionData() ([]*models.Position, error) {
	var out []*models.Position
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) NotificationData() ([]*models.Notification, error) {
	var out []*models.Notification
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) QuoteData() ([]*models.Quote, error) {
	var out []*models.Quote
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) SettlementData() ([]*models.Settlement, error) {
	var out []*models.Settlement
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) TradeData() ([]*models.Trade, error) {
	var out []*models.Trade
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) TradeBinData() ([]*models.TradeBin, error) {
	var out []*models.TradeBin
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) TransactionData() ([]*models.Transaction, error) {
	var out []*models.Transaction
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SubscriptionResponse) WalletData() ([]*models.Wallet, error) {
	var out []*models.Wallet
	if err := json.Unmarshal(s.Data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type SubscriptionFilter struct {
	Symbol  string `json:"symbol"`
	Account int64  `json:"account"`
}

type AutoCancelResponse struct {
	Now        time.Time `json:"now"`
	CancelTime time.Time `json:"cancelTime"`
	Request    *Command  `json:"-"`
}

func (a *AutoCancelResponse) WithRequest(r *Command) *AutoCancelResponse {
	a.Request = r
	return a
}

type ErrorResponse struct {
	Error   string   `json:"error"`
	Request *Command `json:"-"`
}

func (e *ErrorResponse) WithRequest(r *Command) *ErrorResponse {
	e.Request = r
	return e
}

type SuccessResponse struct {
	Subscribe string   `json:"subscribe"`
	Success   *bool    `json:"success"`
	Request   *Command `json:"-"`
}

func (s *SuccessResponse) WithRequest(r *Command) *SuccessResponse {
	s.Request = r
	return s
}

type InfoResponse struct {
	Info      string     `json:"info"`
	Version   string     `json:"version"`
	Timestamp time.Time  `json:"timestamp"`
	Docs      string     `json:"docs"`
	Limit     *LimitInfo `json:"limit"`
}

type LimitInfo struct {
	Remaining int `json:"remaining"`
}
