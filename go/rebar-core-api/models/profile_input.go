package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ProfileInput profile input
// swagger:model profile-input
type ProfileInput struct {

	// name
	Name string `json:"name,omitempty"`

	// tenant id
	TenantID int64 `json:"tenant_id,omitempty"`

	// values
	Values interface{} `json:"values,omitempty"`
}

// Validate validates this profile input
func (m *ProfileInput) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}