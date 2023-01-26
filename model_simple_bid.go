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

// checks if the SimpleBid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimpleBid{}

// SimpleBid Simple bid.
type SimpleBid struct {
	Market ReserveMarket `json:"market"`
	ProductType ProductType `json:"productType"`
	// Bid identification
	GroupId *string `json:"groupId,omitempty"`
	BidType *BidType `json:"bidType,omitempty"`
	// Delivery day of offered control reserve / energy (ISO 8601 format YYYY-MM-DD).
	DeliveryDate *string `json:"deliveryDate,omitempty"`
	ProductName *ProductName `json:"productName,omitempty"`
	TimeInterval *TimeInterval `json:"timeInterval,omitempty"`
	Direction *Direction `json:"direction,omitempty"`
	ConnectingZone ConnectingZone `json:"connectingZone"`
	Activation *ActivationType `json:"activation,omitempty"`
	MinQuantity *Quantity `json:"minQuantity,omitempty"`
	OfferedQuantity Quantity `json:"offeredQuantity"`
	// Price information of the bid.
	Prices []Price `json:"prices"`
	TechnicalLinkage *TechnicalLinkage `json:"technicalLinkage,omitempty"`
	ConditionalLinkage *ConditionalLinkage `json:"conditionalLinkage,omitempty"`
	BackupFor *BalancingServiceProvider `json:"backupFor,omitempty"`
	// Custom tag for bids
	Tag *string `json:"tag,omitempty"`
}

// NewSimpleBid instantiates a new SimpleBid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleBid(market ReserveMarket, productType ProductType, connectingZone ConnectingZone, offeredQuantity Quantity, prices []Price) *SimpleBid {
	this := SimpleBid{}
	this.Market = market
	this.ProductType = productType
	var bidType BidType = SIMPLE
	this.BidType = &bidType
	this.ConnectingZone = connectingZone
	var activation ActivationType = DIRECT
	this.Activation = &activation
	this.OfferedQuantity = offeredQuantity
	this.Prices = prices
	return &this
}

// NewSimpleBidWithDefaults instantiates a new SimpleBid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleBidWithDefaults() *SimpleBid {
	this := SimpleBid{}
	var bidType BidType = SIMPLE
	this.BidType = &bidType
	var activation ActivationType = DIRECT
	this.Activation = &activation
	return &this
}

// GetMarket returns the Market field value
func (o *SimpleBid) GetMarket() ReserveMarket {
	if o == nil {
		var ret ReserveMarket
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetMarketOk() (*ReserveMarket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value
func (o *SimpleBid) SetMarket(v ReserveMarket) {
	o.Market = v
}

// GetProductType returns the ProductType field value
func (o *SimpleBid) GetProductType() ProductType {
	if o == nil {
		var ret ProductType
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetProductTypeOk() (*ProductType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *SimpleBid) SetProductType(v ProductType) {
	o.ProductType = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *SimpleBid) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *SimpleBid) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *SimpleBid) SetGroupId(v string) {
	o.GroupId = &v
}

// GetBidType returns the BidType field value if set, zero value otherwise.
func (o *SimpleBid) GetBidType() BidType {
	if o == nil || isNil(o.BidType) {
		var ret BidType
		return ret
	}
	return *o.BidType
}

// GetBidTypeOk returns a tuple with the BidType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetBidTypeOk() (*BidType, bool) {
	if o == nil || isNil(o.BidType) {
		return nil, false
	}
	return o.BidType, true
}

// HasBidType returns a boolean if a field has been set.
func (o *SimpleBid) HasBidType() bool {
	if o != nil && !isNil(o.BidType) {
		return true
	}

	return false
}

// SetBidType gets a reference to the given BidType and assigns it to the BidType field.
func (o *SimpleBid) SetBidType(v BidType) {
	o.BidType = &v
}

// GetDeliveryDate returns the DeliveryDate field value if set, zero value otherwise.
func (o *SimpleBid) GetDeliveryDate() string {
	if o == nil || isNil(o.DeliveryDate) {
		var ret string
		return ret
	}
	return *o.DeliveryDate
}

// GetDeliveryDateOk returns a tuple with the DeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetDeliveryDateOk() (*string, bool) {
	if o == nil || isNil(o.DeliveryDate) {
		return nil, false
	}
	return o.DeliveryDate, true
}

// HasDeliveryDate returns a boolean if a field has been set.
func (o *SimpleBid) HasDeliveryDate() bool {
	if o != nil && !isNil(o.DeliveryDate) {
		return true
	}

	return false
}

// SetDeliveryDate gets a reference to the given string and assigns it to the DeliveryDate field.
func (o *SimpleBid) SetDeliveryDate(v string) {
	o.DeliveryDate = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *SimpleBid) GetProductName() ProductName {
	if o == nil || isNil(o.ProductName) {
		var ret ProductName
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetProductNameOk() (*ProductName, bool) {
	if o == nil || isNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *SimpleBid) HasProductName() bool {
	if o != nil && !isNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given ProductName and assigns it to the ProductName field.
func (o *SimpleBid) SetProductName(v ProductName) {
	o.ProductName = &v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *SimpleBid) GetTimeInterval() TimeInterval {
	if o == nil || isNil(o.TimeInterval) {
		var ret TimeInterval
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetTimeIntervalOk() (*TimeInterval, bool) {
	if o == nil || isNil(o.TimeInterval) {
		return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *SimpleBid) HasTimeInterval() bool {
	if o != nil && !isNil(o.TimeInterval) {
		return true
	}

	return false
}

// SetTimeInterval gets a reference to the given TimeInterval and assigns it to the TimeInterval field.
func (o *SimpleBid) SetTimeInterval(v TimeInterval) {
	o.TimeInterval = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *SimpleBid) GetDirection() Direction {
	if o == nil || isNil(o.Direction) {
		var ret Direction
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetDirectionOk() (*Direction, bool) {
	if o == nil || isNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *SimpleBid) HasDirection() bool {
	if o != nil && !isNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given Direction and assigns it to the Direction field.
func (o *SimpleBid) SetDirection(v Direction) {
	o.Direction = &v
}

// GetConnectingZone returns the ConnectingZone field value
func (o *SimpleBid) GetConnectingZone() ConnectingZone {
	if o == nil {
		var ret ConnectingZone
		return ret
	}

	return o.ConnectingZone
}

// GetConnectingZoneOk returns a tuple with the ConnectingZone field value
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetConnectingZoneOk() (*ConnectingZone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectingZone, true
}

// SetConnectingZone sets field value
func (o *SimpleBid) SetConnectingZone(v ConnectingZone) {
	o.ConnectingZone = v
}

// GetActivation returns the Activation field value if set, zero value otherwise.
func (o *SimpleBid) GetActivation() ActivationType {
	if o == nil || isNil(o.Activation) {
		var ret ActivationType
		return ret
	}
	return *o.Activation
}

// GetActivationOk returns a tuple with the Activation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetActivationOk() (*ActivationType, bool) {
	if o == nil || isNil(o.Activation) {
		return nil, false
	}
	return o.Activation, true
}

// HasActivation returns a boolean if a field has been set.
func (o *SimpleBid) HasActivation() bool {
	if o != nil && !isNil(o.Activation) {
		return true
	}

	return false
}

// SetActivation gets a reference to the given ActivationType and assigns it to the Activation field.
func (o *SimpleBid) SetActivation(v ActivationType) {
	o.Activation = &v
}

// GetMinQuantity returns the MinQuantity field value if set, zero value otherwise.
func (o *SimpleBid) GetMinQuantity() Quantity {
	if o == nil || isNil(o.MinQuantity) {
		var ret Quantity
		return ret
	}
	return *o.MinQuantity
}

// GetMinQuantityOk returns a tuple with the MinQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetMinQuantityOk() (*Quantity, bool) {
	if o == nil || isNil(o.MinQuantity) {
		return nil, false
	}
	return o.MinQuantity, true
}

// HasMinQuantity returns a boolean if a field has been set.
func (o *SimpleBid) HasMinQuantity() bool {
	if o != nil && !isNil(o.MinQuantity) {
		return true
	}

	return false
}

// SetMinQuantity gets a reference to the given Quantity and assigns it to the MinQuantity field.
func (o *SimpleBid) SetMinQuantity(v Quantity) {
	o.MinQuantity = &v
}

// GetOfferedQuantity returns the OfferedQuantity field value
func (o *SimpleBid) GetOfferedQuantity() Quantity {
	if o == nil {
		var ret Quantity
		return ret
	}

	return o.OfferedQuantity
}

// GetOfferedQuantityOk returns a tuple with the OfferedQuantity field value
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetOfferedQuantityOk() (*Quantity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferedQuantity, true
}

// SetOfferedQuantity sets field value
func (o *SimpleBid) SetOfferedQuantity(v Quantity) {
	o.OfferedQuantity = v
}

// GetPrices returns the Prices field value
func (o *SimpleBid) GetPrices() []Price {
	if o == nil {
		var ret []Price
		return ret
	}

	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetPricesOk() ([]Price, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prices, true
}

// SetPrices sets field value
func (o *SimpleBid) SetPrices(v []Price) {
	o.Prices = v
}

// GetTechnicalLinkage returns the TechnicalLinkage field value if set, zero value otherwise.
func (o *SimpleBid) GetTechnicalLinkage() TechnicalLinkage {
	if o == nil || isNil(o.TechnicalLinkage) {
		var ret TechnicalLinkage
		return ret
	}
	return *o.TechnicalLinkage
}

// GetTechnicalLinkageOk returns a tuple with the TechnicalLinkage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetTechnicalLinkageOk() (*TechnicalLinkage, bool) {
	if o == nil || isNil(o.TechnicalLinkage) {
		return nil, false
	}
	return o.TechnicalLinkage, true
}

// HasTechnicalLinkage returns a boolean if a field has been set.
func (o *SimpleBid) HasTechnicalLinkage() bool {
	if o != nil && !isNil(o.TechnicalLinkage) {
		return true
	}

	return false
}

// SetTechnicalLinkage gets a reference to the given TechnicalLinkage and assigns it to the TechnicalLinkage field.
func (o *SimpleBid) SetTechnicalLinkage(v TechnicalLinkage) {
	o.TechnicalLinkage = &v
}

// GetConditionalLinkage returns the ConditionalLinkage field value if set, zero value otherwise.
func (o *SimpleBid) GetConditionalLinkage() ConditionalLinkage {
	if o == nil || isNil(o.ConditionalLinkage) {
		var ret ConditionalLinkage
		return ret
	}
	return *o.ConditionalLinkage
}

// GetConditionalLinkageOk returns a tuple with the ConditionalLinkage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetConditionalLinkageOk() (*ConditionalLinkage, bool) {
	if o == nil || isNil(o.ConditionalLinkage) {
		return nil, false
	}
	return o.ConditionalLinkage, true
}

// HasConditionalLinkage returns a boolean if a field has been set.
func (o *SimpleBid) HasConditionalLinkage() bool {
	if o != nil && !isNil(o.ConditionalLinkage) {
		return true
	}

	return false
}

// SetConditionalLinkage gets a reference to the given ConditionalLinkage and assigns it to the ConditionalLinkage field.
func (o *SimpleBid) SetConditionalLinkage(v ConditionalLinkage) {
	o.ConditionalLinkage = &v
}

// GetBackupFor returns the BackupFor field value if set, zero value otherwise.
func (o *SimpleBid) GetBackupFor() BalancingServiceProvider {
	if o == nil || isNil(o.BackupFor) {
		var ret BalancingServiceProvider
		return ret
	}
	return *o.BackupFor
}

// GetBackupForOk returns a tuple with the BackupFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetBackupForOk() (*BalancingServiceProvider, bool) {
	if o == nil || isNil(o.BackupFor) {
		return nil, false
	}
	return o.BackupFor, true
}

// HasBackupFor returns a boolean if a field has been set.
func (o *SimpleBid) HasBackupFor() bool {
	if o != nil && !isNil(o.BackupFor) {
		return true
	}

	return false
}

// SetBackupFor gets a reference to the given BalancingServiceProvider and assigns it to the BackupFor field.
func (o *SimpleBid) SetBackupFor(v BalancingServiceProvider) {
	o.BackupFor = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *SimpleBid) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleBid) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *SimpleBid) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *SimpleBid) SetTag(v string) {
	o.Tag = &v
}

func (o SimpleBid) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimpleBid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["market"] = o.Market
	toSerialize["productType"] = o.ProductType
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.BidType) {
		toSerialize["bidType"] = o.BidType
	}
	if !isNil(o.DeliveryDate) {
		toSerialize["deliveryDate"] = o.DeliveryDate
	}
	if !isNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}
	if !isNil(o.TimeInterval) {
		toSerialize["timeInterval"] = o.TimeInterval
	}
	if !isNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	toSerialize["connectingZone"] = o.ConnectingZone
	if !isNil(o.Activation) {
		toSerialize["activation"] = o.Activation
	}
	if !isNil(o.MinQuantity) {
		toSerialize["minQuantity"] = o.MinQuantity
	}
	toSerialize["offeredQuantity"] = o.OfferedQuantity
	toSerialize["prices"] = o.Prices
	if !isNil(o.TechnicalLinkage) {
		toSerialize["technicalLinkage"] = o.TechnicalLinkage
	}
	if !isNil(o.ConditionalLinkage) {
		toSerialize["conditionalLinkage"] = o.ConditionalLinkage
	}
	if !isNil(o.BackupFor) {
		toSerialize["backupFor"] = o.BackupFor
	}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

type NullableSimpleBid struct {
	value *SimpleBid
	isSet bool
}

func (v NullableSimpleBid) Get() *SimpleBid {
	return v.value
}

func (v *NullableSimpleBid) Set(val *SimpleBid) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleBid) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleBid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleBid(val *SimpleBid) *NullableSimpleBid {
	return &NullableSimpleBid{value: val, isSet: true}
}

func (v NullableSimpleBid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleBid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


