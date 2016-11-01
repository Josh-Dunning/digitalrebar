package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// DeploymentsInput deployments input
// swagger:model deployments-input
type DeploymentsInput struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// nodes
	Nodes interface{} `json:"nodes,omitempty"`

	// parent id
	ParentID int64 `json:"parent_id,omitempty"`

	// system
	System bool `json:"system,omitempty"`

	// tenant id
	TenantID int64 `json:"tenant_id,omitempty"`
}

// Validate validates this deployments input
func (m *DeploymentsInput) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}