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

// ReserveMarket Indication of the reserve market:   * `CAPACITY` - FCR/aFCR/mFCR capacity market (ger. \"Regelleistungsmarkt\"),   * `ENERGY`- aFRR/mFRR energy market (ger. \"Regelarbeitsmarkt\"). 
type ReserveMarket string

// List of ReserveMarket
const (
	CAPACITY ReserveMarket = "CAPACITY"
	ENERGY ReserveMarket = "ENERGY"
)

// All allowed values of ReserveMarket enum
var AllowedReserveMarketEnumValues = []ReserveMarket{
	"CAPACITY",
	"ENERGY",
}

func (v *ReserveMarket) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReserveMarket(value)
	for _, existing := range AllowedReserveMarketEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReserveMarket", value)
}

// NewReserveMarketFromValue returns a pointer to a valid ReserveMarket
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReserveMarketFromValue(v string) (*ReserveMarket, error) {
	ev := ReserveMarket(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReserveMarket: valid values are %v", v, AllowedReserveMarketEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReserveMarket) IsValid() bool {
	for _, existing := range AllowedReserveMarketEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReserveMarket value
func (v ReserveMarket) Ptr() *ReserveMarket {
	return &v
}

type NullableReserveMarket struct {
	value *ReserveMarket
	isSet bool
}

func (v NullableReserveMarket) Get() *ReserveMarket {
	return v.value
}

func (v *NullableReserveMarket) Set(val *ReserveMarket) {
	v.value = val
	v.isSet = true
}

func (v NullableReserveMarket) IsSet() bool {
	return v.isSet
}

func (v *NullableReserveMarket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReserveMarket(val *ReserveMarket) *NullableReserveMarket {
	return &NullableReserveMarket{value: val, isSet: true}
}

func (v NullableReserveMarket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReserveMarket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

