// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableConsolePort writable console port
//
// swagger:model WritableConsolePort
type WritableConsolePort struct {

	// occupied
	// Read Only: true
	Occupied *bool `json:"_occupied,omitempty"`

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Cable end
	// Read Only: true
	// Min Length: 1
	CableEnd string `json:"cable_end,omitempty"`

	// Connected endpoints reachable
	// Read Only: true
	ConnectedEndpointsReachable *bool `json:"connected_endpoints_reachable,omitempty"`

	// Connected endpoints type
	// Read Only: true
	ConnectedEndpointsType string `json:"connected_endpoints_type,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device
	// Required: true
	Device *int64 `json:"device"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// Link peers type
	// Read Only: true
	LinkPeersType string `json:"link_peers_type,omitempty"`

	// Mark connected
	//
	// Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected"`

	// Module
	Module *int64 `json:"module"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Speed
	//
	// Port speed in bits per second
	// Enum: [1200 2400 4800 9600 19200 38400 57600 115200]
	Speed *int64 `json:"speed"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Type
	//
	// Physical port type
	// Enum: [de-9 db-25 rj-11 rj-12 rj-45 mini-din-8 usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b usb-micro-ab other]
	Type string `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable console port
func (m *WritableConsolePort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCableEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableConsolePort) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritableConsolePort) validateCableEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.CableEnd) { // not required
		return nil
	}

	if err := validate.MinLength("cable_end", "body", m.CableEnd, 1); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

var writableConsolePortTypeSpeedPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1200,2400,4800,9600,19200,38400,57600,115200]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableConsolePortTypeSpeedPropEnum = append(writableConsolePortTypeSpeedPropEnum, v)
	}
}

// prop value enum
func (m *WritableConsolePort) validateSpeedEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, writableConsolePortTypeSpeedPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableConsolePort) validateSpeed(formats strfmt.Registry) error {
	if swag.IsZero(m.Speed) { // not required
		return nil
	}

	// value enum
	if err := m.validateSpeedEnum("speed", "body", *m.Speed); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var writableConsolePortTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["de-9","db-25","rj-11","rj-12","rj-45","mini-din-8","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","usb-micro-ab","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableConsolePortTypeTypePropEnum = append(writableConsolePortTypeTypePropEnum, v)
	}
}

const (

	// WritableConsolePortTypeDeDash9 captures enum value "de-9"
	WritableConsolePortTypeDeDash9 string = "de-9"

	// WritableConsolePortTypeDbDash25 captures enum value "db-25"
	WritableConsolePortTypeDbDash25 string = "db-25"

	// WritableConsolePortTypeRjDash11 captures enum value "rj-11"
	WritableConsolePortTypeRjDash11 string = "rj-11"

	// WritableConsolePortTypeRjDash12 captures enum value "rj-12"
	WritableConsolePortTypeRjDash12 string = "rj-12"

	// WritableConsolePortTypeRjDash45 captures enum value "rj-45"
	WritableConsolePortTypeRjDash45 string = "rj-45"

	// WritableConsolePortTypeMiniDashDinDash8 captures enum value "mini-din-8"
	WritableConsolePortTypeMiniDashDinDash8 string = "mini-din-8"

	// WritableConsolePortTypeUsbDasha captures enum value "usb-a"
	WritableConsolePortTypeUsbDasha string = "usb-a"

	// WritableConsolePortTypeUsbDashb captures enum value "usb-b"
	WritableConsolePortTypeUsbDashb string = "usb-b"

	// WritableConsolePortTypeUsbDashc captures enum value "usb-c"
	WritableConsolePortTypeUsbDashc string = "usb-c"

	// WritableConsolePortTypeUsbDashMiniDasha captures enum value "usb-mini-a"
	WritableConsolePortTypeUsbDashMiniDasha string = "usb-mini-a"

	// WritableConsolePortTypeUsbDashMiniDashb captures enum value "usb-mini-b"
	WritableConsolePortTypeUsbDashMiniDashb string = "usb-mini-b"

	// WritableConsolePortTypeUsbDashMicroDasha captures enum value "usb-micro-a"
	WritableConsolePortTypeUsbDashMicroDasha string = "usb-micro-a"

	// WritableConsolePortTypeUsbDashMicroDashb captures enum value "usb-micro-b"
	WritableConsolePortTypeUsbDashMicroDashb string = "usb-micro-b"

	// WritableConsolePortTypeUsbDashMicroDashAb captures enum value "usb-micro-ab"
	WritableConsolePortTypeUsbDashMicroDashAb string = "usb-micro-ab"

	// WritableConsolePortTypeOther captures enum value "other"
	WritableConsolePortTypeOther string = "other"
)

// prop value enum
func (m *WritableConsolePort) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableConsolePortTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableConsolePort) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable console port based on the context it is used
func (m *WritableConsolePort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOccupied(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCableEnd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointsReachable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointsType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinkPeersType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableConsolePort) contextValidateOccupied(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_occupied", "body", m.Occupied); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {
		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritableConsolePort) contextValidateCableEnd(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_end", "body", string(m.CableEnd)); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) contextValidateConnectedEndpointsReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoints_reachable", "body", m.ConnectedEndpointsReachable); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) contextValidateConnectedEndpointsType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoints_type", "body", string(m.ConnectedEndpointsType)); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) contextValidateLinkPeersType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "link_peers_type", "body", string(m.LinkPeersType)); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePort) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableConsolePort) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableConsolePort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableConsolePort) UnmarshalBinary(b []byte) error {
	var res WritableConsolePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
