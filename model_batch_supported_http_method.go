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

// BatchSupportedHttpMethod What type of request to make.
type BatchSupportedHttpMethod string

// List of BatchSupportedHttpMethod
const (
	POST BatchSupportedHttpMethod = "POST"
	PATCH BatchSupportedHttpMethod = "PATCH"
	DELETE BatchSupportedHttpMethod = "DELETE"
)

// All allowed values of BatchSupportedHttpMethod enum
var AllowedBatchSupportedHttpMethodEnumValues = []BatchSupportedHttpMethod{
	"POST",
	"PATCH",
	"DELETE",
}

func (v *BatchSupportedHttpMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BatchSupportedHttpMethod(value)
	for _, existing := range AllowedBatchSupportedHttpMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BatchSupportedHttpMethod", value)
}

// NewBatchSupportedHttpMethodFromValue returns a pointer to a valid BatchSupportedHttpMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBatchSupportedHttpMethodFromValue(v string) (*BatchSupportedHttpMethod, error) {
	ev := BatchSupportedHttpMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BatchSupportedHttpMethod: valid values are %v", v, AllowedBatchSupportedHttpMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BatchSupportedHttpMethod) IsValid() bool {
	for _, existing := range AllowedBatchSupportedHttpMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BatchSupportedHttpMethod value
func (v BatchSupportedHttpMethod) Ptr() *BatchSupportedHttpMethod {
	return &v
}

type NullableBatchSupportedHttpMethod struct {
	value *BatchSupportedHttpMethod
	isSet bool
}

func (v NullableBatchSupportedHttpMethod) Get() *BatchSupportedHttpMethod {
	return v.value
}

func (v *NullableBatchSupportedHttpMethod) Set(val *BatchSupportedHttpMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchSupportedHttpMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchSupportedHttpMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchSupportedHttpMethod(val *BatchSupportedHttpMethod) *NullableBatchSupportedHttpMethod {
	return &NullableBatchSupportedHttpMethod{value: val, isSet: true}
}

func (v NullableBatchSupportedHttpMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchSupportedHttpMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

