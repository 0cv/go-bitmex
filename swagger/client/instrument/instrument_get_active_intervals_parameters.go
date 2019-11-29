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

	strfmt "github.com/go-openapi/strfmt"
)

// NewInstrumentGetActiveIntervalsParams creates a new InstrumentGetActiveIntervalsParams object
// with the default values initialized.
func NewInstrumentGetActiveIntervalsParams() *InstrumentGetActiveIntervalsParams {

	return &InstrumentGetActiveIntervalsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInstrumentGetActiveIntervalsParamsWithTimeout creates a new InstrumentGetActiveIntervalsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInstrumentGetActiveIntervalsParamsWithTimeout(timeout time.Duration) *InstrumentGetActiveIntervalsParams {

	return &InstrumentGetActiveIntervalsParams{

		timeout: timeout,
	}
}

// NewInstrumentGetActiveIntervalsParamsWithContext creates a new InstrumentGetActiveIntervalsParams object
// with the default values initialized, and the ability to set a context for a request
func NewInstrumentGetActiveIntervalsParamsWithContext(ctx context.Context) *InstrumentGetActiveIntervalsParams {

	return &InstrumentGetActiveIntervalsParams{

		Context: ctx,
	}
}

// NewInstrumentGetActiveIntervalsParamsWithHTTPClient creates a new InstrumentGetActiveIntervalsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInstrumentGetActiveIntervalsParamsWithHTTPClient(client *http.Client) *InstrumentGetActiveIntervalsParams {

	return &InstrumentGetActiveIntervalsParams{
		HTTPClient: client,
	}
}

/*InstrumentGetActiveIntervalsParams contains all the parameters to send to the API endpoint
for the instrument get active intervals operation typically these are written to a http.Request
*/
type InstrumentGetActiveIntervalsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the instrument get active intervals params
func (o *InstrumentGetActiveIntervalsParams) WithTimeout(timeout time.Duration) *InstrumentGetActiveIntervalsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the instrument get active intervals params
func (o *InstrumentGetActiveIntervalsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the instrument get active intervals params
func (o *InstrumentGetActiveIntervalsParams) WithContext(ctx context.Context) *InstrumentGetActiveIntervalsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the instrument get active intervals params
func (o *InstrumentGetActiveIntervalsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the instrument get active intervals params
func (o *InstrumentGetActiveIntervalsParams) WithHTTPClient(client *http.Client) *InstrumentGetActiveIntervalsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the instrument get active intervals params
func (o *InstrumentGetActiveIntervalsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *InstrumentGetActiveIntervalsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
