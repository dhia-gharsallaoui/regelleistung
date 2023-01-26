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

// BidType the model 'BidType'
type BidType string

// List of BidType
const (
	SIMPLE BidType = "SIMPLE"
	MULTIPART BidType = "MULTIPART"
	MULTIPART_COMPONENT BidType = "MULTIPART_COMPONENT"
	EXCLUSIVE BidType = "EXCLUSIVE"
	EXCLUSIVE_COMPONENT BidType = "EXCLUSIVE_COMPONENT"
	GROUP BidType = "GROUP"
)

// All allowed values of BidType enum
var AllowedBidTypeEnumValues = []BidType{
	"SIMPLE",
	"MULTIPART",
	"MULTIPART_COMPONENT",
	"EXCLUSIVE",
	"EXCLUSIVE_COMPONENT",
	"GROUP",
}

func (v *BidType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BidType(value)
	for _, existing := range AllowedBidTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BidType", value)
}

// NewBidTypeFromValue returns a pointer to a valid BidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBidTypeFromValue(v string) (*BidType, error) {
	ev := BidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BidType: valid values are %v", v, AllowedBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BidType) IsValid() bool {
	for _, existing := range AllowedBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BidType value
func (v BidType) Ptr() *BidType {
	return &v
}

type NullableBidType struct {
	value *BidType
	isSet bool
}

func (v NullableBidType) Get() *BidType {
	return v.value
}

func (v *NullableBidType) Set(val *BidType) {
	v.value = val
	v.isSet = true
}

func (v NullableBidType) IsSet() bool {
	return v.isSet
}

func (v *NullableBidType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidType(val *BidType) *NullableBidType {
	return &NullableBidType{value: val, isSet: true}
}

func (v NullableBidType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBidType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
