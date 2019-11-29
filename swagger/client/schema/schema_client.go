// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new schema API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for schema API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SchemaGet gets model schemata for data objects returned by this API
*/
func (a *Client) SchemaGet(params *SchemaGetParams) (*SchemaGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Schema.get",
		Method:             "GET",
		PathPattern:        "/schema",
		ProducesMediaTypes: []string{"application/javascript", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SchemaGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaGetOK), nil

}

/*
SchemaWebsocketHelp returns help text and subject list for websocket usage
*/
func (a *Client) SchemaWebsocketHelp(params *SchemaWebsocketHelpParams) (*SchemaWebsocketHelpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaWebsocketHelpParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Schema.websocketHelp",
		Method:             "GET",
		PathPattern:        "/schema/websocketHelp",
		ProducesMediaTypes: []string{"application/javascript", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SchemaWebsocketHelpReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaWebsocketHelpOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
