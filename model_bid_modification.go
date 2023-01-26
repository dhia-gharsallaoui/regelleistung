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

// checks if the BidModification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BidModification{}

// BidModification Representation of bid modification.
type BidModification struct {
	ConnectingZone *ConnectingZone `json:"connectingZone,omitempty"`
	MinQuantity *Quantity `json:"minQuantity,omitempty"`
	OfferedQuantity *Quantity `json:"offeredQuantity,omitempty"`
	// Price information of the bid.
	Prices []Price `json:"prices,omitempty"`
	Activation *ActivationType `json:"activation,omitempty"`
	TechnicalLinkage *TechnicalLinkage `json:"technicalLinkage,omitempty"`
	ConditionalLinkage *ConditionalLinkage `json:"conditionalLinkage,omitempty"`
	BackupFor *BalancingServiceProvider `json:"backupFor,omitempty"`
}

// NewBidModification instantiates a new BidModification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBidModification() *BidModification {
	this := BidModification{}
	var activation ActivationType = DIRECT
	this.Activation = &activation
	return &this
}

// NewBidModificationWithDefaults instantiates a new BidModification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBidModificationWithDefaults() *BidModification {
	this := BidModification{}
	var activation ActivationType = DIRECT
	this.Activation = &activation
	return &this
}

// GetConnectingZone returns the ConnectingZone field value if set, zero value otherwise.
func (o *BidModification) GetConnectingZone() ConnectingZone {
	if o == nil || isNil(o.ConnectingZone) {
		var ret ConnectingZone
		return ret
	}
	return *o.ConnectingZone
}

// GetConnectingZoneOk returns a tuple with the ConnectingZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidModification) GetConnectingZoneOk() (*ConnectingZone, bool) {
	if o == nil || isNil(o.ConnectingZone) {
		return nil, false
	}
	return o.ConnectingZone, true
}

// HasConnectingZone returns a boolean if a field has been set.
func (o *BidModification) HasConnectingZone() bool {
	if o != nil && !isNil(o.ConnectingZone) {
		return true
	}

	return false
}

// SetConnectingZone gets a reference to the given ConnectingZone and assigns it to the ConnectingZone field.
func (o *BidModification) SetConnectingZone(v ConnectingZone) {
	o.ConnectingZone = &v
}

// GetMinQuantity returns the MinQuantity field value if set, zero value otherwise.
func (o *BidModification) GetMinQuantity() Quantity {
	if o == nil || isNil(o.MinQuantity) {
		var ret Quantity
		return ret
	}
	return *o.MinQuantity
}

// GetMinQuantityOk returns a tuple with the MinQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidModification) GetMinQuantityOk() (*Quantity, bool) {
	if o == nil || isNil(o.MinQuantity) {
		return nil, false
	}
	return o.MinQuantity, true
}

// HasMinQuantity returns a boolean if a field has been set.
func (o *BidModification) HasMinQuantity() bool {
	if o != nil && !isNil(o.MinQuantity) {
		return true
	}

	return false
}

// SetMinQuantity gets a reference to the given Quantity and assigns it to the MinQuantity field.
func (o *BidModification) SetMinQuantity(v Quantity) {
	o.MinQuantity = &v
}

// GetOfferedQuantity returns the OfferedQuantity field value if set, zero value otherwise.
func (o *BidModification) GetOfferedQuantity() Quantity {
	if o == nil || isNil(o.OfferedQuantity) {
		var ret Quantity
		return ret
	}
	return *o.OfferedQuantity
}

// GetOfferedQuantityOk returns a tuple with the OfferedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidModification) GetOfferedQuantityOk() (*Quantity, bool) {
	if o == nil || isNil(o.OfferedQuantity) {
		return nil, false
	}
	return o.OfferedQuantity, true
}

// HasOfferedQuantity returns a boolean if a field has been set.
func (o *BidModification) HasOfferedQuantity() bool {
	if o != nil && !isNil(o.OfferedQuantity) {
		return true
	}

	return false
}

// SetOfferedQuantity gets a reference to the given Quantity and assigns it to the OfferedQuantity field.
func (o *BidModification) SetOfferedQuantity(v Quantity) {
	o.OfferedQuantity = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *BidModification) GetPrices() []Price {
	if o == nil || isNil(o.Prices) {
		var ret []Price
		return ret
	}
	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidModification) GetPricesOk() ([]Price, bool) {
	if o == nil || isNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *BidModification) HasPrices() bool {
	if o != nil && !isNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given []Price and assigns it to the Prices field.
func (o *BidModification) SetPrices(v []Price) {
	o.Prices = v
}

// GetActivation returns the Activation field value if set, zero value otherwise.
func (o *BidModification) GetActivation() ActivationType {
	if o == nil || isNil(o.Activation) {
		var ret ActivationType
		return ret
	}
	return *o.Activation
}

// GetActivationOk returns a tuple with the Activation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidModification) GetActivationOk() (*ActivationType, bool) {
	if o == nil || isNil(o.Activation) {
		return nil, false
	}
	return o.Activation, true
}

// HasActivation returns a boolean if a field has been set.
func (o *BidModification) HasActivation() bool {
	if o != nil && !isNil(o.Activation) {
		return true
	}

	return false
}

// SetActivation gets a reference to the given ActivationType and assigns it to the Activation field.
func (o *BidModification) SetActivation(v ActivationType) {
	o.Activation = &v
}

// GetTechnicalLinkage returns the TechnicalLinkage field value if set, zero value otherwise.
func (o *BidModification) GetTechnicalLinkage() TechnicalLinkage {
	if o == nil || isNil(o.TechnicalLinkage) {
		var ret TechnicalLinkage
		return ret
	}
	return *o.TechnicalLinkage
}

// GetTechnicalLinkageOk returns a tuple with the TechnicalLinkage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidModification) GetTechnicalLinkageOk() (*TechnicalLinkage, bool) {
	if o == nil || isNil(o.TechnicalLinkage) {
		return nil, false
	}
	return o.TechnicalLinkage, true
}

// HasTechnicalLinkage returns a boolean if a field has been set.
func (o *BidModification) HasTechnicalLinkage() bool {
	if o != nil && !isNil(o.TechnicalLinkage) {
		return true
	}

	return false
}

// SetTechnicalLinkage gets a reference to the given TechnicalLinkage and assigns it to the TechnicalLinkage field.
func (o *BidModification) SetTechnicalLinkage(v TechnicalLinkage) {
	o.TechnicalLinkage = &v
}

// GetConditionalLinkage returns the ConditionalLinkage field value if set, zero value otherwise.
func (o *BidModification) GetConditionalLinkage() ConditionalLinkage {
	if o == nil || isNil(o.ConditionalLinkage) {
		var ret ConditionalLinkage
		return ret
	}
	return *o.ConditionalLinkage
}

// GetConditionalLinkageOk returns a tuple with the ConditionalLinkage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidModification) GetConditionalLinkageOk() (*ConditionalLinkage, bool) {
	if o == nil || isNil(o.ConditionalLinkage) {
		return nil, false
	}
	return o.ConditionalLinkage, true
}

// HasConditionalLinkage returns a boolean if a field has been set.
func (o *BidModification) HasConditionalLinkage() bool {
	if o != nil && !isNil(o.ConditionalLinkage) {
		return true
	}

	return false
}

// SetConditionalLinkage gets a reference to the given ConditionalLinkage and assigns it to the ConditionalLinkage field.
func (o *BidModification) SetConditionalLinkage(v ConditionalLinkage) {
	o.ConditionalLinkage = &v
}

// GetBackupFor returns the BackupFor field value if set, zero value otherwise.
func (o *BidModification) GetBackupFor() BalancingServiceProvider {
	if o == nil || isNil(o.BackupFor) {
		var ret BalancingServiceProvider
		return ret
	}
	return *o.BackupFor
}

// GetBackupForOk returns a tuple with the BackupFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidModification) GetBackupForOk() (*BalancingServiceProvider, bool) {
	if o == nil || isNil(o.BackupFor) {
		return nil, false
	}
	return o.BackupFor, true
}

// HasBackupFor returns a boolean if a field has been set.
func (o *BidModification) HasBackupFor() bool {
	if o != nil && !isNil(o.BackupFor) {
		return true
	}

	return false
}

// SetBackupFor gets a reference to the given BalancingServiceProvider and assigns it to the BackupFor field.
func (o *BidModification) SetBackupFor(v BalancingServiceProvider) {
	o.BackupFor = &v
}

func (o BidModification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BidModification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ConnectingZone) {
		toSerialize["connectingZone"] = o.ConnectingZone
	}
	if !isNil(o.MinQuantity) {
		toSerialize["minQuantity"] = o.MinQuantity
	}
	if !isNil(o.OfferedQuantity) {
		toSerialize["offeredQuantity"] = o.OfferedQuantity
	}
	if !isNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	if !isNil(o.Activation) {
		toSerialize["activation"] = o.Activation
	}
	if !isNil(o.TechnicalLinkage) {
		toSerialize["technicalLinkage"] = o.TechnicalLinkage
	}
	if !isNil(o.ConditionalLinkage) {
		toSerialize["conditionalLinkage"] = o.ConditionalLinkage
	}
	if !isNil(o.BackupFor) {
		toSerialize["backupFor"] = o.BackupFor
	}
	return toSerialize, nil
}

type NullableBidModification struct {
	value *BidModification
	isSet bool
}

func (v NullableBidModification) Get() *BidModification {
	return v.value
}

func (v *NullableBidModification) Set(val *BidModification) {
	v.value = val
	v.isSet = true
}

func (v NullableBidModification) IsSet() bool {
	return v.isSet
}

func (v *NullableBidModification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidModification(val *BidModification) *NullableBidModification {
	return &NullableBidModification{value: val, isSet: true}
}

func (v NullableBidModification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBidModification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


