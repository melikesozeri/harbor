// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Tag tag
// swagger:model Tag
type Tag struct {

	// The ID of the artifact that the tag attached to
	ArtifactID int64 `json:"artifact_id,omitempty"`

	// The ID of the tag
	ID int64 `json:"id,omitempty"`

	// The immutable status of the tag
	Immutable bool `json:"immutable"`

	// The name of the tag
	Name string `json:"name,omitempty"`

	// The latest pull time of the tag
	// Format: date-time
	PullTime strfmt.DateTime `json:"pull_time,omitempty"`

	// The push time of the tag
	// Format: date-time
	PushTime strfmt.DateTime `json:"push_time,omitempty"`

	// The ID of the repository that the tag belongs to
	RepositoryID int64 `json:"repository_id,omitempty"`

	// The attribute indicates whether the tag is signed or not
	Signed bool `json:"signed"`
}

// Validate validates this tag
func (m *Tag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePullTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePushTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tag) validatePullTime(formats strfmt.Registry) error {

	if swag.IsZero(m.PullTime) { // not required
		return nil
	}

	if err := validate.FormatOf("pull_time", "body", "date-time", m.PullTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Tag) validatePushTime(formats strfmt.Registry) error {

	if swag.IsZero(m.PushTime) { // not required
		return nil
	}

	if err := validate.FormatOf("push_time", "body", "date-time", m.PushTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tag) UnmarshalBinary(b []byte) error {
	var res Tag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
