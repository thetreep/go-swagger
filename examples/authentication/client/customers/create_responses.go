// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thetreep/go-swagger/examples/authentication/models"
)

// CreateReader is a Reader for the Create structure.
type CreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCreated creates a CreateCreated with default headers values
func NewCreateCreated() *CreateCreated {
	return &CreateCreated{}
}

/*
CreateCreated describes a response with status code 201, with default header values.

created
*/
type CreateCreated struct {
	Payload *models.Customer
}

// IsSuccess returns true when this create created response has a 2xx status code
func (o *CreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create created response has a 3xx status code
func (o *CreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create created response has a 4xx status code
func (o *CreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create created response has a 5xx status code
func (o *CreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create created response a status code equal to that given
func (o *CreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create created response
func (o *CreateCreated) Code() int {
	return 201
}

func (o *CreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /customers][%d] createCreated %s", 201, payload)
}

func (o *CreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /customers][%d] createCreated %s", 201, payload)
}

func (o *CreateCreated) GetPayload() *models.Customer {
	return o.Payload
}

func (o *CreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDefault creates a CreateDefault with default headers values
func NewCreateDefault(code int) *CreateDefault {
	return &CreateDefault{
		_statusCode: code,
	}
}

/*
CreateDefault describes a response with status code -1, with default header values.

error
*/
type CreateDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create default response has a 2xx status code
func (o *CreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create default response has a 3xx status code
func (o *CreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create default response has a 4xx status code
func (o *CreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create default response has a 5xx status code
func (o *CreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create default response a status code equal to that given
func (o *CreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create default response
func (o *CreateDefault) Code() int {
	return o._statusCode
}

func (o *CreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /customers][%d] create default %s", o._statusCode, payload)
}

func (o *CreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /customers][%d] create default %s", o._statusCode, payload)
}

func (o *CreateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
