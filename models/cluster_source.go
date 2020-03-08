// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ClusterSource cluster source
// swagger:model ClusterSource
type ClusterSource string

const (

	// ClusterSourceUI captures enum value "UI"
	ClusterSourceUI ClusterSource = "UI"

	// ClusterSourceJOB captures enum value "JOB"
	ClusterSourceJOB ClusterSource = "JOB"

	// ClusterSourceAPI captures enum value "API"
	ClusterSourceAPI ClusterSource = "API"
)

// for schema
var clusterSourceEnum []interface{}

func init() {
	var res []ClusterSource
	if err := json.Unmarshal([]byte(`["UI","JOB","API"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterSourceEnum = append(clusterSourceEnum, v)
	}
}

func (m ClusterSource) validateClusterSourceEnum(path, location string, value ClusterSource) error {
	if err := validate.Enum(path, location, value, clusterSourceEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this cluster source
func (m ClusterSource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClusterSourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
