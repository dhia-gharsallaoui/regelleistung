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

// checks if the Demand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Demand{}

// Demand Representation of demand object.
type Demand struct {
	ConnectingZone *ConnectingZone `json:"connectingZone,omitempty"`
	ProductType *ProductType `json:"productType,omitempty"`
	ProductName *ProductName `json:"productName,omitempty"`
	Direction *Direction `json:"direction,omitempty"`
	TimeInterval *TimeInterval `json:"timeInterval,omitempty"`
	Data *DemandData `json:"data,omitempty"`
}

// NewDemand instantiates a new Demand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDemand() *Demand {
	this := Demand{}
	return &this
}

// NewDemandWithDefaults instantiates a new Demand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDemandWithDefaults() *Demand {
	this := Demand{}
	return &this
}

// GetConnectingZone returns the ConnectingZone field value if set, zero value otherwise.
func (o *Demand) GetConnectingZone() ConnectingZone {
	if o == nil || isNil(o.ConnectingZone) {
		var ret ConnectingZone
		return ret
	}
	return *o.ConnectingZone
}

// GetConnectingZoneOk returns a tuple with the ConnectingZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Demand) GetConnectingZoneOk() (*ConnectingZone, bool) {
	if o == nil || isNil(o.ConnectingZone) {
		return nil, false
	}
	return o.ConnectingZone, true
}

// HasConnectingZone returns a boolean if a field has been set.
func (o *Demand) HasConnectingZone() bool {
	if o != nil && !isNil(o.ConnectingZone) {
		return true
	}

	return false
}

// SetConnectingZone gets a reference to the given ConnectingZone and assigns it to the ConnectingZone field.
func (o *Demand) SetConnectingZone(v ConnectingZone) {
	o.ConnectingZone = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *Demand) GetProductType() ProductType {
	if o == nil || isNil(o.ProductType) {
		var ret ProductType
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Demand) GetProductTypeOk() (*ProductType, bool) {
	if o == nil || isNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *Demand) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given ProductType and assigns it to the ProductType field.
func (o *Demand) SetProductType(v ProductType) {
	o.ProductType = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *Demand) GetProductName() ProductName {
	if o == nil || isNil(o.ProductName) {
		var ret ProductName
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Demand) GetProductNameOk() (*ProductName, bool) {
	if o == nil || isNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *Demand) HasProductName() bool {
	if o != nil && !isNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given ProductName and assigns it to the ProductName field.
func (o *Demand) SetProductName(v ProductName) {
	o.ProductName = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *Demand) GetDirection() Direction {
	if o == nil || isNil(o.Direction) {
		var ret Direction
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Demand) GetDirectionOk() (*Direction, bool) {
	if o == nil || isNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *Demand) HasDirection() bool {
	if o != nil && !isNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given Direction and assigns it to the Direction field.
func (o *Demand) SetDirection(v Direction) {
	o.Direction = &v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *Demand) GetTimeInterval() TimeInterval {
	if o == nil || isNil(o.TimeInterval) {
		var ret TimeInterval
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Demand) GetTimeIntervalOk() (*TimeInterval, bool) {
	if o == nil || isNil(o.TimeInterval) {
		return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *Demand) HasTimeInterval() bool {
	if o != nil && !isNil(o.TimeInterval) {
		return true
	}

	return false
}

// SetTimeInterval gets a reference to the given TimeInterval and assigns it to the TimeInterval field.
func (o *Demand) SetTimeInterval(v TimeInterval) {
	o.TimeInterval = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Demand) GetData() DemandData {
	if o == nil || isNil(o.Data) {
		var ret DemandData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Demand) GetDataOk() (*DemandData, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Demand) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DemandData and assigns it to the Data field.
func (o *Demand) SetData(v DemandData) {
	o.Data = &v
}

func (o Demand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Demand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ConnectingZone) {
		toSerialize["connectingZone"] = o.ConnectingZone
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !isNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}
	if !isNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !isNil(o.TimeInterval) {
		toSerialize["timeInterval"] = o.TimeInterval
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDemand struct {
	value *Demand
	isSet bool
}

func (v NullableDemand) Get() *Demand {
	return v.value
}

func (v *NullableDemand) Set(val *Demand) {
	v.value = val
	v.isSet = true
}

func (v NullableDemand) IsSet() bool {
	return v.isSet
}

func (v *NullableDemand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDemand(val *Demand) *NullableDemand {
	return &NullableDemand{value: val, isSet: true}
}

func (v NullableDemand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDemand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

