/*
IP RL BSP API

IP RL BSP API for participation in capacity/energy market tenders.

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the TimeInterval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeInterval{}

// TimeInterval Representation of a time interval with start and end times.
type TimeInterval struct {
	// Must be before `endTime` in quarter-hour accuracy (UTC).
	StartTime *time.Time `json:"startTime,omitempty"`
	// Must be after `startTime` in quarter-hour accuracy (UTC).
	EndTime *time.Time `json:"endTime,omitempty"`
}

// NewTimeInterval instantiates a new TimeInterval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeInterval() *TimeInterval {
	this := TimeInterval{}
	return &this
}

// NewTimeIntervalWithDefaults instantiates a new TimeInterval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeIntervalWithDefaults() *TimeInterval {
	this := TimeInterval{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *TimeInterval) GetStartTime() time.Time {
	if o == nil || isNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeInterval) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *TimeInterval) HasStartTime() bool {
	if o != nil && !isNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *TimeInterval) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TimeInterval) GetEndTime() time.Time {
	if o == nil || isNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeInterval) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TimeInterval) HasEndTime() bool {
	if o != nil && !isNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *TimeInterval) SetEndTime(v time.Time) {
	o.EndTime = &v
}

func (o TimeInterval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeInterval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !isNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	return toSerialize, nil
}

type NullableTimeInterval struct {
	value *TimeInterval
	isSet bool
}

func (v NullableTimeInterval) Get() *TimeInterval {
	return v.value
}

func (v *NullableTimeInterval) Set(val *TimeInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeInterval(val *TimeInterval) *NullableTimeInterval {
	return &NullableTimeInterval{value: val, isSet: true}
}

func (v NullableTimeInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


