// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	alternate "github.com/thetreep/go-swagger/examples/external-types/fred"
)

// MyCustomArray This generate an array type in models, based on the external type.
//
// []alternate.MyAlternateType
//
// The validation method of the external type is called by the generated array.
//
// swagger:model MyCustomArray
type MyCustomArray []alternate.MyAlternateType

// Validate validates this my custom array
func (m MyCustomArray) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if err := m[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName(strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName(strconv.Itoa(i))
			}
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this my custom array based on context it is used
func (m MyCustomArray) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
