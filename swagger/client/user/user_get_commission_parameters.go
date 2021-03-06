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

// NewUserGetCommissionParams creates a new UserGetCommissionParams object
// with the default values initialized.
func NewUserGetCommissionParams() *UserGetCommissionParams {

	return &UserGetCommissionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserGetCommissionParamsWithTimeout creates a new UserGetCommissionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserGetCommissionParamsWithTimeout(timeout time.Duration) *UserGetCommissionParams {

	return &UserGetCommissionParams{

		timeout: timeout,
	}
}

// NewUserGetCommissionParamsWithContext creates a new UserGetCommissionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserGetCommissionParamsWithContext(ctx context.Context) *UserGetCommissionParams {

	return &UserGetCommissionParams{

		Context: ctx,
	}
}

// NewUserGetCommissionParamsWithHTTPClient creates a new UserGetCommissionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserGetCommissionParamsWithHTTPClient(client *http.Client) *UserGetCommissionParams {

	return &UserGetCommissionParams{
		HTTPClient: client,
	}
}

/*UserGetCommissionParams contains all the parameters to send to the API endpoint
for the user get commission operation typically these are written to a http.Request
*/
type UserGetCommissionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user get commission params
func (o *UserGetCommissionParams) WithTimeout(timeout time.Duration) *UserGetCommissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user get commission params
func (o *UserGetCommissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user get commission params
func (o *UserGetCommissionParams) WithContext(ctx context.Context) *UserGetCommissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user get commission params
func (o *UserGetCommissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user get commission params
func (o *UserGetCommissionParams) WithHTTPClient(client *http.Client) *UserGetCommissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user get commission params
func (o *UserGetCommissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UserGetCommissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
