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

// checks if the ContractValidityPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractValidityPeriod{}

// ContractValidityPeriod Representation of a contract validity period.
type ContractValidityPeriod struct {
	// Must be before ´endDate´.
	StartDate *string `json:"startDate,omitempty"`
	// Must be after ´startDate´.
	EndDate *string `json:"endDate,omitempty"`
}

// NewContractValidityPeriod instantiates a new ContractValidityPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractValidityPeriod() *ContractValidityPeriod {
	this := ContractValidityPeriod{}
	return &this
}

// NewContractValidityPeriodWithDefaults instantiates a new ContractValidityPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractValidityPeriodWithDefaults() *ContractValidityPeriod {
	this := ContractValidityPeriod{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ContractValidityPeriod) GetStartDate() string {
	if o == nil || isNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractValidityPeriod) GetStartDateOk() (*string, bool) {
	if o == nil || isNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ContractValidityPeriod) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *ContractValidityPeriod) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ContractValidityPeriod) GetEndDate() string {
	if o == nil || isNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractValidityPeriod) GetEndDateOk() (*string, bool) {
	if o == nil || isNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ContractValidityPeriod) HasEndDate() bool {
	if o != nil && !isNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *ContractValidityPeriod) SetEndDate(v string) {
	o.EndDate = &v
}

func (o ContractValidityPeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractValidityPeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !isNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableContractValidityPeriod struct {
	value *ContractValidityPeriod
	isSet bool
}

func (v NullableContractValidityPeriod) Get() *ContractValidityPeriod {
	return v.value
}

func (v *NullableContractValidityPeriod) Set(val *ContractValidityPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableContractValidityPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableContractValidityPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractValidityPeriod(val *ContractValidityPeriod) *NullableContractValidityPeriod {
	return &NullableContractValidityPeriod{value: val, isSet: true}
}

func (v NullableContractValidityPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractValidityPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


