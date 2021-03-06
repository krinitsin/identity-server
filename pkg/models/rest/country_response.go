// Code generated by go-swagger; DO NOT EDIT.

package rest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CountryResponse country response
//
// swagger:model CountryResponse
type CountryResponse struct {

	// Country
	Country string `json:"Country,omitempty"`
}

// Validate validates this country response
func (m *CountryResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this country response based on context it is used
func (m *CountryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CountryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CountryResponse) UnmarshalBinary(b []byte) error {
	var res CountryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
