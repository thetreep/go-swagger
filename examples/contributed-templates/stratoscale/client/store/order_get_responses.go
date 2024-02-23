// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thetreep/go-swagger/examples/contributed-templates/stratoscale/models"
)

// OrderGetReader is a Reader for the OrderGet structure.
type OrderGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrderGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrderGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrderGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /store/order/{orderId}] OrderGet", response, response.Code())
	}
}

// NewOrderGetOK creates a OrderGetOK with default headers values
func NewOrderGetOK() *OrderGetOK {
	return &OrderGetOK{}
}

/*
OrderGetOK describes a response with status code 200, with default header values.

successful operation
*/
type OrderGetOK struct {
	Payload *models.Order
}

// IsSuccess returns true when this order get o k response has a 2xx status code
func (o *OrderGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this order get o k response has a 3xx status code
func (o *OrderGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this order get o k response has a 4xx status code
func (o *OrderGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this order get o k response has a 5xx status code
func (o *OrderGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this order get o k response a status code equal to that given
func (o *OrderGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the order get o k response
func (o *OrderGetOK) Code() int {
	return 200
}

func (o *OrderGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /store/order/{orderId}][%d] orderGetOK %s", 200, payload)
}

func (o *OrderGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /store/order/{orderId}][%d] orderGetOK %s", 200, payload)
}

func (o *OrderGetOK) GetPayload() *models.Order {
	return o.Payload
}

func (o *OrderGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderGetBadRequest creates a OrderGetBadRequest with default headers values
func NewOrderGetBadRequest() *OrderGetBadRequest {
	return &OrderGetBadRequest{}
}

/*
OrderGetBadRequest describes a response with status code 400, with default header values.

Invalid ID supplied
*/
type OrderGetBadRequest struct {
}

// IsSuccess returns true when this order get bad request response has a 2xx status code
func (o *OrderGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this order get bad request response has a 3xx status code
func (o *OrderGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this order get bad request response has a 4xx status code
func (o *OrderGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this order get bad request response has a 5xx status code
func (o *OrderGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this order get bad request response a status code equal to that given
func (o *OrderGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the order get bad request response
func (o *OrderGetBadRequest) Code() int {
	return 400
}

func (o *OrderGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /store/order/{orderId}][%d] orderGetBadRequest", 400)
}

func (o *OrderGetBadRequest) String() string {
	return fmt.Sprintf("[GET /store/order/{orderId}][%d] orderGetBadRequest", 400)
}

func (o *OrderGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrderGetNotFound creates a OrderGetNotFound with default headers values
func NewOrderGetNotFound() *OrderGetNotFound {
	return &OrderGetNotFound{}
}

/*
OrderGetNotFound describes a response with status code 404, with default header values.

Order not found
*/
type OrderGetNotFound struct {
}

// IsSuccess returns true when this order get not found response has a 2xx status code
func (o *OrderGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this order get not found response has a 3xx status code
func (o *OrderGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this order get not found response has a 4xx status code
func (o *OrderGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this order get not found response has a 5xx status code
func (o *OrderGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this order get not found response a status code equal to that given
func (o *OrderGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the order get not found response
func (o *OrderGetNotFound) Code() int {
	return 404
}

func (o *OrderGetNotFound) Error() string {
	return fmt.Sprintf("[GET /store/order/{orderId}][%d] orderGetNotFound", 404)
}

func (o *OrderGetNotFound) String() string {
	return fmt.Sprintf("[GET /store/order/{orderId}][%d] orderGetNotFound", 404)
}

func (o *OrderGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
