// Code generated by go-swagger; DO NOT EDIT.

package instrument

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewInstrumentGetParams creates a new InstrumentGetParams object
// with the default values initialized.
func NewInstrumentGetParams() *InstrumentGetParams {
	var (
		countDefault   = int32(100)
		reverseDefault = bool(false)
		startDefault   = int32(0)
	)
	return &InstrumentGetParams{
		Count:   &countDefault,
		Reverse: &reverseDefault,
		Start:   &startDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewInstrumentGetParamsWithTimeout creates a new InstrumentGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInstrumentGetParamsWithTimeout(timeout time.Duration) *InstrumentGetParams {
	var (
		countDefault   = int32(100)
		reverseDefault = bool(false)
		startDefault   = int32(0)
	)
	return &InstrumentGetParams{
		Count:   &countDefault,
		Reverse: &reverseDefault,
		Start:   &startDefault,

		timeout: timeout,
	}
}

// NewInstrumentGetParamsWithContext creates a new InstrumentGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewInstrumentGetParamsWithContext(ctx context.Context) *InstrumentGetParams {
	var (
		countDefault   = int32(100)
		reverseDefault = bool(false)
		startDefault   = int32(0)
	)
	return &InstrumentGetParams{
		Count:   &countDefault,
		Reverse: &reverseDefault,
		Start:   &startDefault,

		Context: ctx,
	}
}

// NewInstrumentGetParamsWithHTTPClient creates a new InstrumentGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInstrumentGetParamsWithHTTPClient(client *http.Client) *InstrumentGetParams {
	var (
		countDefault   = int32(100)
		reverseDefault = bool(false)
		startDefault   = int32(0)
	)
	return &InstrumentGetParams{
		Count:      &countDefault,
		Reverse:    &reverseDefault,
		Start:      &startDefault,
		HTTPClient: client,
	}
}

/*InstrumentGetParams contains all the parameters to send to the API endpoint
for the instrument get operation typically these are written to a http.Request
*/
type InstrumentGetParams struct {

	/*Columns
	  Array of column names to fetch. If omitted, will return all columns.

	Note that this method will always return item keys, even when not specified, so you may receive more columns that you expect.

	*/
	Columns *string
	/*Count
	  Number of results to fetch.

	*/
	Count *int32
	/*EndTime
	  Ending date filter for results.

	*/
	EndTime *strfmt.DateTime
	/*Filter
	  Generic table filter. Send JSON key/value pairs, such as `{"key": "value"}`. You can key on individual fields, and do more advanced querying on timestamps. See the [Timestamp Docs](https://www.bitmex.com/app/restAPI#Timestamp-Filters) for more details.

	*/
	Filter *string
	/*Reverse
	  If true, will sort results newest first.

	*/
	Reverse *bool
	/*Start
	  Starting point for results.

	*/
	Start *int32
	/*StartTime
	  Starting date filter for results.

	*/
	StartTime *strfmt.DateTime
	/*Symbol
	  Instrument symbol. Send a bare series (e.g. XBU) to get data for the nearest expiring contract in that series.

	You can also send a timeframe, e.g. `XBU:monthly`. Timeframes are `daily`, `weekly`, `monthly`, `quarterly`, and `biquarterly`.

	*/
	Symbol *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the instrument get params
func (o *InstrumentGetParams) WithTimeout(timeout time.Duration) *InstrumentGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the instrument get params
func (o *InstrumentGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the instrument get params
func (o *InstrumentGetParams) WithContext(ctx context.Context) *InstrumentGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the instrument get params
func (o *InstrumentGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the instrument get params
func (o *InstrumentGetParams) WithHTTPClient(client *http.Client) *InstrumentGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the instrument get params
func (o *InstrumentGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithColumns adds the columns to the instrument get params
func (o *InstrumentGetParams) WithColumns(columns *string) *InstrumentGetParams {
	o.SetColumns(columns)
	return o
}

// SetColumns adds the columns to the instrument get params
func (o *InstrumentGetParams) SetColumns(columns *string) {
	o.Columns = columns
}

// WithCount adds the count to the instrument get params
func (o *InstrumentGetParams) WithCount(count *int32) *InstrumentGetParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the instrument get params
func (o *InstrumentGetParams) SetCount(count *int32) {
	o.Count = count
}

// WithEndTime adds the endTime to the instrument get params
func (o *InstrumentGetParams) WithEndTime(endTime *strfmt.DateTime) *InstrumentGetParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the instrument get params
func (o *InstrumentGetParams) SetEndTime(endTime *strfmt.DateTime) {
	o.EndTime = endTime
}

// WithFilter adds the filter to the instrument get params
func (o *InstrumentGetParams) WithFilter(filter *string) *InstrumentGetParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the instrument get params
func (o *InstrumentGetParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithReverse adds the reverse to the instrument get params
func (o *InstrumentGetParams) WithReverse(reverse *bool) *InstrumentGetParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the instrument get params
func (o *InstrumentGetParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithStart adds the start to the instrument get params
func (o *InstrumentGetParams) WithStart(start *int32) *InstrumentGetParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the instrument get params
func (o *InstrumentGetParams) SetStart(start *int32) {
	o.Start = start
}

// WithStartTime adds the startTime to the instrument get params
func (o *InstrumentGetParams) WithStartTime(startTime *strfmt.DateTime) *InstrumentGetParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the instrument get params
func (o *InstrumentGetParams) SetStartTime(startTime *strfmt.DateTime) {
	o.StartTime = startTime
}

// WithSymbol adds the symbol to the instrument get params
func (o *InstrumentGetParams) WithSymbol(symbol *string) *InstrumentGetParams {
	o.SetSymbol(symbol)
	return o
}

// SetSymbol adds the symbol to the instrument get params
func (o *InstrumentGetParams) SetSymbol(symbol *string) {
	o.Symbol = symbol
}

// WriteToRequest writes these params to a swagger request
func (o *InstrumentGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Columns != nil {

		// query param columns
		var qrColumns string
		if o.Columns != nil {
			qrColumns = *o.Columns
		}
		qColumns := qrColumns
		if qColumns != "" {
			if err := r.SetQueryParam("columns", qColumns); err != nil {
				return err
			}
		}

	}

	if o.Count != nil {

		// query param count
		var qrCount int32
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.EndTime != nil {

		// query param endTime
		var qrEndTime strfmt.DateTime
		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := qrEndTime.String()
		if qEndTime != "" {
			if err := r.SetQueryParam("endTime", qEndTime); err != nil {
				return err
			}
		}

	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool
		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {
			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart int32
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt32(qrStart)
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if o.StartTime != nil {

		// query param startTime
		var qrStartTime strfmt.DateTime
		if o.StartTime != nil {
			qrStartTime = *o.StartTime
		}
		qStartTime := qrStartTime.String()
		if qStartTime != "" {
			if err := r.SetQueryParam("startTime", qStartTime); err != nil {
				return err
			}
		}

	}

	if o.Symbol != nil {

		// query param symbol
		var qrSymbol string
		if o.Symbol != nil {
			qrSymbol = *o.Symbol
		}
		qSymbol := qrSymbol
		if qSymbol != "" {
			if err := r.SetQueryParam("symbol", qSymbol); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
