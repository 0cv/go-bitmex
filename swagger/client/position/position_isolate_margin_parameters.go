// Code generated by go-swagger; DO NOT EDIT.

package position

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

// NewPositionIsolateMarginParams creates a new PositionIsolateMarginParams object
// with the default values initialized.
func NewPositionIsolateMarginParams() *PositionIsolateMarginParams {
	var (
		enabledDefault = bool(true)
	)
	return &PositionIsolateMarginParams{
		Enabled: &enabledDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPositionIsolateMarginParamsWithTimeout creates a new PositionIsolateMarginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPositionIsolateMarginParamsWithTimeout(timeout time.Duration) *PositionIsolateMarginParams {
	var (
		enabledDefault = bool(true)
	)
	return &PositionIsolateMarginParams{
		Enabled: &enabledDefault,

		timeout: timeout,
	}
}

// NewPositionIsolateMarginParamsWithContext creates a new PositionIsolateMarginParams object
// with the default values initialized, and the ability to set a context for a request
func NewPositionIsolateMarginParamsWithContext(ctx context.Context) *PositionIsolateMarginParams {
	var (
		enabledDefault = bool(true)
	)
	return &PositionIsolateMarginParams{
		Enabled: &enabledDefault,

		Context: ctx,
	}
}

// NewPositionIsolateMarginParamsWithHTTPClient creates a new PositionIsolateMarginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPositionIsolateMarginParamsWithHTTPClient(client *http.Client) *PositionIsolateMarginParams {
	var (
		enabledDefault = bool(true)
	)
	return &PositionIsolateMarginParams{
		Enabled:    &enabledDefault,
		HTTPClient: client,
	}
}

/*PositionIsolateMarginParams contains all the parameters to send to the API endpoint
for the position isolate margin operation typically these are written to a http.Request
*/
type PositionIsolateMarginParams struct {

	/*Enabled
	  True for isolated margin, false for cross margin.

	*/
	Enabled *bool
	/*Symbol
	  Position symbol to isolate.

	*/
	Symbol string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the position isolate margin params
func (o *PositionIsolateMarginParams) WithTimeout(timeout time.Duration) *PositionIsolateMarginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the position isolate margin params
func (o *PositionIsolateMarginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the position isolate margin params
func (o *PositionIsolateMarginParams) WithContext(ctx context.Context) *PositionIsolateMarginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the position isolate margin params
func (o *PositionIsolateMarginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the position isolate margin params
func (o *PositionIsolateMarginParams) WithHTTPClient(client *http.Client) *PositionIsolateMarginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the position isolate margin params
func (o *PositionIsolateMarginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabled adds the enabled to the position isolate margin params
func (o *PositionIsolateMarginParams) WithEnabled(enabled *bool) *PositionIsolateMarginParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the position isolate margin params
func (o *PositionIsolateMarginParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithSymbol adds the symbol to the position isolate margin params
func (o *PositionIsolateMarginParams) WithSymbol(symbol string) *PositionIsolateMarginParams {
	o.SetSymbol(symbol)
	return o
}

// SetSymbol adds the symbol to the position isolate margin params
func (o *PositionIsolateMarginParams) SetSymbol(symbol string) {
	o.Symbol = symbol
}

// WriteToRequest writes these params to a swagger request
func (o *PositionIsolateMarginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Enabled != nil {

		// form param enabled
		var frEnabled bool
		if o.Enabled != nil {
			frEnabled = *o.Enabled
		}
		fEnabled := swag.FormatBool(frEnabled)
		if fEnabled != "" {
			if err := r.SetFormParam("enabled", fEnabled); err != nil {
				return err
			}
		}

	}

	// form param symbol
	frSymbol := o.Symbol
	fSymbol := frSymbol
	if fSymbol != "" {
		if err := r.SetFormParam("symbol", fSymbol); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
