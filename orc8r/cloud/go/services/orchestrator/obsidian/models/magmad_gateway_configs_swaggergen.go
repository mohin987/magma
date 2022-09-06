// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MagmadGatewayConfigs Configuration for the magmad gateway agent
//
// swagger:model magmad_gateway_configs
type MagmadGatewayConfigs struct {

	// autoupgrade enabled
	// Example: true
	// Required: true
	AutoupgradeEnabled *bool `json:"autoupgrade_enabled"`

	// autoupgrade poll interval
	// Example: 300
	// Required: true
	// Minimum: 30
	AutoupgradePollInterval int32 `json:"autoupgrade_poll_interval"`

	// checkin interval
	// Example: 60
	// Required: true
	// Minimum: 15
	CheckinInterval uint32 `json:"checkin_interval"`

	// checkin timeout
	// Example: 10
	// Required: true
	// Minimum: 5
	CheckinTimeout uint32 `json:"checkin_timeout"`

	// dynamic services
	// Example: []
	DynamicServices []string `json:"dynamic_services"`

	// feature flags
	// Example: {"newfeature1":true,"newfeature2":false}
	FeatureFlags map[string]bool `json:"feature_flags,omitempty"`

	// logging
	Logging *GatewayLoggingConfigs `json:"logging,omitempty"`

	// vpn
	Vpn *GatewayVpnConfigs `json:"vpn,omitempty"`
}

// Validate validates this magmad gateway configs
func (m *MagmadGatewayConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoupgradeEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoupgradePollInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCheckinInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCheckinTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDynamicServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogging(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MagmadGatewayConfigs) validateAutoupgradeEnabled(formats strfmt.Registry) error {

	if err := validate.Required("autoupgrade_enabled", "body", m.AutoupgradeEnabled); err != nil {
		return err
	}

	return nil
}

func (m *MagmadGatewayConfigs) validateAutoupgradePollInterval(formats strfmt.Registry) error {

	if err := validate.Required("autoupgrade_poll_interval", "body", int32(m.AutoupgradePollInterval)); err != nil {
		return err
	}

	if err := validate.MinimumInt("autoupgrade_poll_interval", "body", int64(m.AutoupgradePollInterval), 30, false); err != nil {
		return err
	}

	return nil
}

func (m *MagmadGatewayConfigs) validateCheckinInterval(formats strfmt.Registry) error {

	if err := validate.Required("checkin_interval", "body", uint32(m.CheckinInterval)); err != nil {
		return err
	}

	if err := validate.MinimumUint("checkin_interval", "body", uint64(m.CheckinInterval), 15, false); err != nil {
		return err
	}

	return nil
}

func (m *MagmadGatewayConfigs) validateCheckinTimeout(formats strfmt.Registry) error {

	if err := validate.Required("checkin_timeout", "body", uint32(m.CheckinTimeout)); err != nil {
		return err
	}

	if err := validate.MinimumUint("checkin_timeout", "body", uint64(m.CheckinTimeout), 5, false); err != nil {
		return err
	}

	return nil
}

func (m *MagmadGatewayConfigs) validateDynamicServices(formats strfmt.Registry) error {
	if swag.IsZero(m.DynamicServices) { // not required
		return nil
	}

	for i := 0; i < len(m.DynamicServices); i++ {

		if err := validate.MinLength("dynamic_services"+"."+strconv.Itoa(i), "body", m.DynamicServices[i], 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *MagmadGatewayConfigs) validateLogging(formats strfmt.Registry) error {
	if swag.IsZero(m.Logging) { // not required
		return nil
	}

	if m.Logging != nil {
		if err := m.Logging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logging")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logging")
			}
			return err
		}
	}

	return nil
}

func (m *MagmadGatewayConfigs) validateVpn(formats strfmt.Registry) error {
	if swag.IsZero(m.Vpn) { // not required
		return nil
	}

	if m.Vpn != nil {
		if err := m.Vpn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpn")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpn")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this magmad gateway configs based on the context it is used
func (m *MagmadGatewayConfigs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogging(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MagmadGatewayConfigs) contextValidateLogging(ctx context.Context, formats strfmt.Registry) error {

	if m.Logging != nil {
		if err := m.Logging.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logging")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logging")
			}
			return err
		}
	}

	return nil
}

func (m *MagmadGatewayConfigs) contextValidateVpn(ctx context.Context, formats strfmt.Registry) error {

	if m.Vpn != nil {
		if err := m.Vpn.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpn")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpn")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MagmadGatewayConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MagmadGatewayConfigs) UnmarshalBinary(b []byte) error {
	var res MagmadGatewayConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}