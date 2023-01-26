/*
IP RL BSP API

IP RL BSP API for participation in capacity/energy market tenders.

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ActivationType the model 'ActivationType'
type ActivationType string

// List of ActivationType
const (
	DIRECT ActivationType = "DIRECT"
	SCHEDULED ActivationType = "SCHEDULED"
)

// All allowed values of ActivationType enum
var AllowedActivationTypeEnumValues = []ActivationType{
	"DIRECT",
	"SCHEDULED",
}

func (v *ActivationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActivationType(value)
	for _, existing := range AllowedActivationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ActivationType", value)
}

// NewActivationTypeFromValue returns a pointer to a valid ActivationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActivationTypeFromValue(v string) (*ActivationType, error) {
	ev := ActivationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActivationType: valid values are %v", v, AllowedActivationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActivationType) IsValid() bool {
	for _, existing := range AllowedActivationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActivationType value
func (v ActivationType) Ptr() *ActivationType {
	return &v
}

type NullableActivationType struct {
	value *ActivationType
	isSet bool
}

func (v NullableActivationType) Get() *ActivationType {
	return v.value
}

func (v *NullableActivationType) Set(val *ActivationType) {
	v.value = val
	v.isSet = true
}

func (v NullableActivationType) IsSet() bool {
	return v.isSet
}

func (v *NullableActivationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivationType(val *ActivationType) *NullableActivationType {
	return &NullableActivationType{value: val, isSet: true}
}

func (v NullableActivationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

