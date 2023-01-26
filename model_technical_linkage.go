/*
IP RL BSP API

IP RL BSP API for participation in capacity/energy market tenders.

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TechnicalLinkage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TechnicalLinkage{}

// TechnicalLinkage struct for TechnicalLinkage
type TechnicalLinkage struct {
	// Custom technical link ID
	LinkId NullableString `json:"linkId"`
}

// NewTechnicalLinkage instantiates a new TechnicalLinkage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechnicalLinkage(linkId NullableString) *TechnicalLinkage {
	this := TechnicalLinkage{}
	this.LinkId = linkId
	return &this
}

// NewTechnicalLinkageWithDefaults instantiates a new TechnicalLinkage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechnicalLinkageWithDefaults() *TechnicalLinkage {
	this := TechnicalLinkage{}
	return &this
}

// GetLinkId returns the LinkId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TechnicalLinkage) GetLinkId() string {
	if o == nil || o.LinkId.Get() == nil {
		var ret string
		return ret
	}

	return *o.LinkId.Get()
}

// GetLinkIdOk returns a tuple with the LinkId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TechnicalLinkage) GetLinkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkId.Get(), o.LinkId.IsSet()
}

// SetLinkId sets field value
func (o *TechnicalLinkage) SetLinkId(v string) {
	o.LinkId.Set(&v)
}

func (o TechnicalLinkage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TechnicalLinkage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["linkId"] = o.LinkId.Get()
	return toSerialize, nil
}

type NullableTechnicalLinkage struct {
	value *TechnicalLinkage
	isSet bool
}

func (v NullableTechnicalLinkage) Get() *TechnicalLinkage {
	return v.value
}

func (v *NullableTechnicalLinkage) Set(val *TechnicalLinkage) {
	v.value = val
	v.isSet = true
}

func (v NullableTechnicalLinkage) IsSet() bool {
	return v.isSet
}

func (v *NullableTechnicalLinkage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechnicalLinkage(val *TechnicalLinkage) *NullableTechnicalLinkage {
	return &NullableTechnicalLinkage{value: val, isSet: true}
}

func (v NullableTechnicalLinkage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechnicalLinkage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


