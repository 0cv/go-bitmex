// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Chat Trollbox Data
// swagger:model Chat
type Chat struct {

	// channel ID
	ChannelID float64 `json:"channelID,omitempty"`

	// date
	// Required: true
	// Format: date-time
	Date *strfmt.DateTime `json:"date"`

	// from bot
	FromBot *bool `json:"fromBot,omitempty"`

	// html
	// Required: true
	HTML *string `json:"html"`

	// id
	ID int32 `json:"id,omitempty"`

	// message
	// Required: true
	Message *string `json:"message"`

	// user
	// Required: true
	User *string `json:"user"`
}

// Validate validates this chat
func (m *Chat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTML(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Chat) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Chat) validateHTML(formats strfmt.Registry) error {

	if err := validate.Required("html", "body", m.HTML); err != nil {
		return err
	}

	return nil
}

func (m *Chat) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *Chat) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Chat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Chat) UnmarshalBinary(b []byte) error {
	var res Chat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
