// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUserGetDepositAddressParams creates a new UserGetDepositAddressParams object
// with the default values initialized.
func NewUserGetDepositAddressParams() *UserGetDepositAddressParams {
	var (
		currencyDefault = string("XBt")
	)
	return &UserGetDepositAddressParams{
		Currency: &currencyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUserGetDepositAddressParamsWithTimeout creates a new UserGetDepositAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserGetDepositAddressParamsWithTimeout(timeout time.Duration) *UserGetDepositAddressParams {
	var (
		currencyDefault = string("XBt")
	)
	return &UserGetDepositAddressParams{
		Currency: &currencyDefault,

		timeout: timeout,
	}
}

// NewUserGetDepositAddressParamsWithContext creates a new UserGetDepositAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserGetDepositAddressParamsWithContext(ctx context.Context) *UserGetDepositAddressParams {
	var (
		currencyDefault = string("XBt")
	)
	return &UserGetDepositAddressParams{
		Currency: &currencyDefault,

		Context: ctx,
	}
}

// NewUserGetDepositAddressParamsWithHTTPClient creates a new UserGetDepositAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserGetDepositAddressParamsWithHTTPClient(client *http.Client) *UserGetDepositAddressParams {
	var (
		currencyDefault = string("XBt")
	)
	return &UserGetDepositAddressParams{
		Currency:   &currencyDefault,
		HTTPClient: client,
	}
}

/*UserGetDepositAddressParams contains all the parameters to send to the API endpoint
for the user get deposit address operation typically these are written to a http.Request
*/
type UserGetDepositAddressParams struct {

	/*Currency*/
	Currency *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user get deposit address params
func (o *UserGetDepositAddressParams) WithTimeout(timeout time.Duration) *UserGetDepositAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user get deposit address params
func (o *UserGetDepositAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user get deposit address params
func (o *UserGetDepositAddressParams) WithContext(ctx context.Context) *UserGetDepositAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user get deposit address params
func (o *UserGetDepositAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user get deposit address params
func (o *UserGetDepositAddressParams) WithHTTPClient(client *http.Client) *UserGetDepositAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user get deposit address params
func (o *UserGetDepositAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrency adds the currency to the user get deposit address params
func (o *UserGetDepositAddressParams) WithCurrency(currency *string) *UserGetDepositAddressParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the user get deposit address params
func (o *UserGetDepositAddressParams) SetCurrency(currency *string) {
	o.Currency = currency
}

// WriteToRequest writes these params to a swagger request
func (o *UserGetDepositAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Currency != nil {

		// query param currency
		var qrCurrency string
		if o.Currency != nil {
			qrCurrency = *o.Currency
		}
		qCurrency := qrCurrency
		if qCurrency != "" {
			if err := r.SetQueryParam("currency", qCurrency); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
