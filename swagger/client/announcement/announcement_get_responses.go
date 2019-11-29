// Code generated by go-swagger; DO NOT EDIT.

package announcement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/adampointer/go-bitmex/swagger/models"
)

// AnnouncementGetReader is a Reader for the AnnouncementGet structure.
type AnnouncementGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AnnouncementGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAnnouncementGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAnnouncementGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAnnouncementGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAnnouncementGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAnnouncementGetOK creates a AnnouncementGetOK with default headers values
func NewAnnouncementGetOK() *AnnouncementGetOK {
	return &AnnouncementGetOK{}
}

/*AnnouncementGetOK handles this case with default header values.

Request was successful
*/
type AnnouncementGetOK struct {
	Payload []*models.Announcement
}

func (o *AnnouncementGetOK) Error() string {
	return fmt.Sprintf("[GET /announcement][%d] announcementGetOK  %+v", 200, o.Payload)
}

func (o *AnnouncementGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnnouncementGetBadRequest creates a AnnouncementGetBadRequest with default headers values
func NewAnnouncementGetBadRequest() *AnnouncementGetBadRequest {
	return &AnnouncementGetBadRequest{}
}

/*AnnouncementGetBadRequest handles this case with default header values.

Parameter Error
*/
type AnnouncementGetBadRequest struct {
	Payload *models.Error
}

func (o *AnnouncementGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /announcement][%d] announcementGetBadRequest  %+v", 400, o.Payload)
}

func (o *AnnouncementGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnnouncementGetUnauthorized creates a AnnouncementGetUnauthorized with default headers values
func NewAnnouncementGetUnauthorized() *AnnouncementGetUnauthorized {
	return &AnnouncementGetUnauthorized{}
}

/*AnnouncementGetUnauthorized handles this case with default header values.

Unauthorized
*/
type AnnouncementGetUnauthorized struct {
	Payload *models.Error
}

func (o *AnnouncementGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /announcement][%d] announcementGetUnauthorized  %+v", 401, o.Payload)
}

func (o *AnnouncementGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnnouncementGetNotFound creates a AnnouncementGetNotFound with default headers values
func NewAnnouncementGetNotFound() *AnnouncementGetNotFound {
	return &AnnouncementGetNotFound{}
}

/*AnnouncementGetNotFound handles this case with default header values.

Not Found
*/
type AnnouncementGetNotFound struct {
	Payload *models.Error
}

func (o *AnnouncementGetNotFound) Error() string {
	return fmt.Sprintf("[GET /announcement][%d] announcementGetNotFound  %+v", 404, o.Payload)
}

func (o *AnnouncementGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
