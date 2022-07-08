// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// NetworkProbeTaskID network probe task id
// Example: imsi1023001
//
// swagger:model network_probe_task_id
type NetworkProbeTaskID string

// Validate validates this network probe task id
func (m NetworkProbeTaskID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network probe task id based on context it is used
func (m NetworkProbeTaskID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}