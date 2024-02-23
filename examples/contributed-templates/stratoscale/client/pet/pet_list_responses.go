// Code generated by go-swagger; DO NOT EDIT.

package pet

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

// PetListReader is a Reader for the PetList structure.
type PetListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PetListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPetListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPetListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pet] PetList", response, response.Code())
	}
}

// NewPetListOK creates a PetListOK with default headers values
func NewPetListOK() *PetListOK {
	return &PetListOK{}
}

/*
PetListOK describes a response with status code 200, with default header values.

successful operation
*/
type PetListOK struct {
	Payload []*models.Pet
}

// IsSuccess returns true when this pet list o k response has a 2xx status code
func (o *PetListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pet list o k response has a 3xx status code
func (o *PetListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pet list o k response has a 4xx status code
func (o *PetListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pet list o k response has a 5xx status code
func (o *PetListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pet list o k response a status code equal to that given
func (o *PetListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pet list o k response
func (o *PetListOK) Code() int {
	return 200
}

func (o *PetListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pet][%d] petListOK %s", 200, payload)
}

func (o *PetListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pet][%d] petListOK %s", 200, payload)
}

func (o *PetListOK) GetPayload() []*models.Pet {
	return o.Payload
}

func (o *PetListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPetListBadRequest creates a PetListBadRequest with default headers values
func NewPetListBadRequest() *PetListBadRequest {
	return &PetListBadRequest{}
}

/*
PetListBadRequest describes a response with status code 400, with default header values.

Invalid status value
*/
type PetListBadRequest struct {
}

// IsSuccess returns true when this pet list bad request response has a 2xx status code
func (o *PetListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pet list bad request response has a 3xx status code
func (o *PetListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pet list bad request response has a 4xx status code
func (o *PetListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pet list bad request response has a 5xx status code
func (o *PetListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pet list bad request response a status code equal to that given
func (o *PetListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pet list bad request response
func (o *PetListBadRequest) Code() int {
	return 400
}

func (o *PetListBadRequest) Error() string {
	return fmt.Sprintf("[GET /pet][%d] petListBadRequest", 400)
}

func (o *PetListBadRequest) String() string {
	return fmt.Sprintf("[GET /pet][%d] petListBadRequest", 400)
}

func (o *PetListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
