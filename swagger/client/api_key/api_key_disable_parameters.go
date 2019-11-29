// Code generated by go-swagger; DO NOT EDIT.

package api_key

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

// NewAPIKeyDisableParams creates a new APIKeyDisableParams object
// with the default values initialized.
func NewAPIKeyDisableParams() *APIKeyDisableParams {
	var ()
	return &APIKeyDisableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIKeyDisableParamsWithTimeout creates a new APIKeyDisableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIKeyDisableParamsWithTimeout(timeout time.Duration) *APIKeyDisableParams {
	var ()
	return &APIKeyDisableParams{

		timeout: timeout,
	}
}

// NewAPIKeyDisableParamsWithContext creates a new APIKeyDisableParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIKeyDisableParamsWithContext(ctx context.Context) *APIKeyDisableParams {
	var ()
	return &APIKeyDisableParams{

		Context: ctx,
	}
}

// NewAPIKeyDisableParamsWithHTTPClient creates a new APIKeyDisableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIKeyDisableParamsWithHTTPClient(client *http.Client) *APIKeyDisableParams {
	var ()
	return &APIKeyDisableParams{
		HTTPClient: client,
	}
}

/*APIKeyDisableParams contains all the parameters to send to the API endpoint
for the API key disable operation typically these are written to a http.Request
*/
type APIKeyDisableParams struct {

	/*APIKeyID
	  API Key ID (public component).

	*/
	APIKeyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the API key disable params
func (o *APIKeyDisableParams) WithTimeout(timeout time.Duration) *APIKeyDisableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the API key disable params
func (o *APIKeyDisableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the API key disable params
func (o *APIKeyDisableParams) WithContext(ctx context.Context) *APIKeyDisableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the API key disable params
func (o *APIKeyDisableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the API key disable params
func (o *APIKeyDisableParams) WithHTTPClient(client *http.Client) *APIKeyDisableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the API key disable params
func (o *APIKeyDisableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKeyID adds the aPIKeyID to the API key disable params
func (o *APIKeyDisableParams) WithAPIKeyID(aPIKeyID string) *APIKeyDisableParams {
	o.SetAPIKeyID(aPIKeyID)
	return o
}

// SetAPIKeyID adds the apiKeyId to the API key disable params
func (o *APIKeyDisableParams) SetAPIKeyID(aPIKeyID string) {
	o.APIKeyID = aPIKeyID
}

// WriteToRequest writes these params to a swagger request
func (o *APIKeyDisableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param apiKeyID
	frAPIKeyID := o.APIKeyID
	fAPIKeyID := frAPIKeyID
	if fAPIKeyID != "" {
		if err := r.SetFormParam("apiKeyID", fAPIKeyID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
