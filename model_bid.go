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

// checks if the Bid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bid{}

// Bid Representation of an submitted bid.
type Bid struct {
	// Bid identification
	Id string `json:"id"`
	// Bid identification
	ComplexBidId *string `json:"complexBidId,omitempty"`
	// Bid identification
	GroupId *string `json:"groupId,omitempty"`
	BidType BidType `json:"bidType"`
	// The tender identification is structured as follows:   `[productType]-[deliveryDate]-[market]-[sequenceNumber]`    * `productType` - one of `fcr`, `afrr` or `mfrr`   * `deliveryDate` - ISO 8601 date   * `market` - one of `cm`(capacity market) or `em`(energy market)   * `sequenceNumber` - An ascending number for the identification of different tender runs or time slices.   For more details on the new tender id format, see the [reference guide.](/docs/guide#tender-id). 
	Tender string `json:"tender"`
	ProductType ProductType `json:"productType"`
	// Delivery day of offered control reserve / energy (ISO 8601 format YYYY-MM-DD).
	DeliveryDate string `json:"deliveryDate"`
	Market ReserveMarket `json:"market"`
	ProductName ProductName `json:"productName"`
	Direction Direction `json:"direction"`
	TimeInterval TimeInterval `json:"timeInterval"`
	ConnectingZone ConnectingZone `json:"connectingZone"`
	// Representation of EIC. The coding scheme is the Energy Identification Coding Scheme (EIC), maintained by ENTSO-e.
	ResourceProvider string `json:"resourceProvider"`
	Status string `json:"status"`
	// Receipt time of the bid.
	ReceiptTimestamp time.Time `json:"receiptTimestamp"`
	// Revision timestamp of the bid.
	RevisionTimestamp time.Time `json:"revisionTimestamp"`
	Activation *ActivationType `json:"activation,omitempty"`
	MinQuantity *Quantity `json:"minQuantity,omitempty"`
	OfferedQuantity Quantity `json:"offeredQuantity"`
	AcceptedQuantity *Quantity `json:"acceptedQuantity,omitempty"`
	// Price information of the bid.
	Prices []Price `json:"prices"`
	TechnicalLinkage *TechnicalLinkage `json:"technicalLinkage,omitempty"`
	ConditionalLinkage *ConditionalLinkage `json:"conditionalLinkage,omitempty"`
	BackupFor *BalancingServiceProvider `json:"backupFor,omitempty"`
	// Custom tag for bids
	Tag *string `json:"tag,omitempty"`
}

// NewBid instantiates a new Bid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBid(id string, bidType BidType, tender string, productType ProductType, deliveryDate string, market ReserveMarket, productName ProductName, direction Direction, timeInterval TimeInterval, connectingZone ConnectingZone, resourceProvider string, status string, receiptTimestamp time.Time, revisionTimestamp time.Time, offeredQuantity Quantity, prices []Price) *Bid {
	this := Bid{}
	this.Id = id
	this.BidType = bidType
	this.Tender = tender
	this.ProductType = productType
	this.DeliveryDate = deliveryDate
	this.Market = market
	this.ProductName = productName
	this.Direction = direction
	this.TimeInterval = timeInterval
	this.ConnectingZone = connectingZone
	this.ResourceProvider = resourceProvider
	this.Status = status
	this.ReceiptTimestamp = receiptTimestamp
	this.RevisionTimestamp = revisionTimestamp
	var activation ActivationType = DIRECT
	this.Activation = &activation
	this.OfferedQuantity = offeredQuantity
	this.Prices = prices
	return &this
}

// NewBidWithDefaults instantiates a new Bid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBidWithDefaults() *Bid {
	this := Bid{}
	var bidType BidType = SIMPLE
	this.BidType = bidType
	var activation ActivationType = DIRECT
	this.Activation = &activation
	return &this
}

// GetId returns the Id field value
func (o *Bid) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Bid) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Bid) SetId(v string) {
	o.Id = v
}

// GetComplexBidId returns the ComplexBidId field value if set, zero value otherwise.
func (o *Bid) GetComplexBidId() string {
	if o == nil || isNil(o.ComplexBidId) {
		var ret string
		return ret
	}
	return *o.ComplexBidId
}

// GetComplexBidIdOk returns a tuple with the ComplexBidId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bid) GetComplexBidIdOk() (*string, bool) {
	if o == nil || isNil(o.ComplexBidId) {
		return nil, false
	}
	return o.ComplexBidId, true
}

// HasComplexBidId returns a boolean if a field has been set.
func (o *Bid) HasComplexBidId() bool {
	if o != nil && !isNil(o.ComplexBidId) {
		return true
	}

	return false
}

// SetComplexBidId gets a reference to the given string and assigns it to the ComplexBidId field.
func (o *Bid) SetComplexBidId(v string) {
	o.ComplexBidId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *Bid) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bid) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *Bid) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *Bid) SetGroupId(v string) {
	o.GroupId = &v
}

// GetBidType returns the BidType field value
func (o *Bid) GetBidType() BidType {
	if o == nil {
		var ret BidType
		return ret
	}

	return o.BidType
}

// GetBidTypeOk returns a tuple with the BidType field value
// and a boolean to check if the value has been set.
func (o *Bid) GetBidTypeOk() (*BidType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BidType, true
}

// SetBidType sets field value
func (o *Bid) SetBidType(v BidType) {
	o.BidType = v
}

// GetTender returns the Tender field value
func (o *Bid) GetTender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tender
}

// GetTenderOk returns a tuple with the Tender field value
// and a boolean to check if the value has been set.
func (o *Bid) GetTenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tender, true
}

// SetTender sets field value
func (o *Bid) SetTender(v string) {
	o.Tender = v
}

// GetProductType returns the ProductType field value
func (o *Bid) GetProductType() ProductType {
	if o == nil {
		var ret ProductType
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *Bid) GetProductTypeOk() (*ProductType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *Bid) SetProductType(v ProductType) {
	o.ProductType = v
}

// GetDeliveryDate returns the DeliveryDate field value
func (o *Bid) GetDeliveryDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryDate
}

// GetDeliveryDateOk returns a tuple with the DeliveryDate field value
// and a boolean to check if the value has been set.
func (o *Bid) GetDeliveryDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryDate, true
}

// SetDeliveryDate sets field value
func (o *Bid) SetDeliveryDate(v string) {
	o.DeliveryDate = v
}

// GetMarket returns the Market field value
func (o *Bid) GetMarket() ReserveMarket {
	if o == nil {
		var ret ReserveMarket
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *Bid) GetMarketOk() (*ReserveMarket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value
func (o *Bid) SetMarket(v ReserveMarket) {
	o.Market = v
}

// GetProductName returns the ProductName field value
func (o *Bid) GetProductName() ProductName {
	if o == nil {
		var ret ProductName
		return ret
	}

	return o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value
// and a boolean to check if the value has been set.
func (o *Bid) GetProductNameOk() (*ProductName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductName, true
}

// SetProductName sets field value
func (o *Bid) SetProductName(v ProductName) {
	o.ProductName = v
}

// GetDirection returns the Direction field value
func (o *Bid) GetDirection() Direction {
	if o == nil {
		var ret Direction
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *Bid) GetDirectionOk() (*Direction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *Bid) SetDirection(v Direction) {
	o.Direction = v
}

// GetTimeInterval returns the TimeInterval field value
func (o *Bid) GetTimeInterval() TimeInterval {
	if o == nil {
		var ret TimeInterval
		return ret
	}

	return o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value
// and a boolean to check if the value has been set.
func (o *Bid) GetTimeIntervalOk() (*TimeInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeInterval, true
}

// SetTimeInterval sets field value
func (o *Bid) SetTimeInterval(v TimeInterval) {
	o.TimeInterval = v
}

// GetConnectingZone returns the ConnectingZone field value
func (o *Bid) GetConnectingZone() ConnectingZone {
	if o == nil {
		var ret ConnectingZone
		return ret
	}

	return o.ConnectingZone
}

// GetConnectingZoneOk returns a tuple with the ConnectingZone field value
// and a boolean to check if the value has been set.
func (o *Bid) GetConnectingZoneOk() (*ConnectingZone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectingZone, true
}

// SetConnectingZone sets field value
func (o *Bid) SetConnectingZone(v ConnectingZone) {
	o.ConnectingZone = v
}

// GetResourceProvider returns the ResourceProvider field value
func (o *Bid) GetResourceProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceProvider
}

// GetResourceProviderOk returns a tuple with the ResourceProvider field value
// and a boolean to check if the value has been set.
func (o *Bid) GetResourceProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceProvider, true
}

// SetResourceProvider sets field value
func (o *Bid) SetResourceProvider(v string) {
	o.ResourceProvider = v
}

// GetStatus returns the Status field value
func (o *Bid) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Bid) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Bid) SetStatus(v string) {
	o.Status = v
}

// GetReceiptTimestamp returns the ReceiptTimestamp field value
func (o *Bid) GetReceiptTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ReceiptTimestamp
}

// GetReceiptTimestampOk returns a tuple with the ReceiptTimestamp field value
// and a boolean to check if the value has been set.
func (o *Bid) GetReceiptTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiptTimestamp, true
}

// SetReceiptTimestamp sets field value
func (o *Bid) SetReceiptTimestamp(v time.Time) {
	o.ReceiptTimestamp = v
}

// GetRevisionTimestamp returns the RevisionTimestamp field value
func (o *Bid) GetRevisionTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.RevisionTimestamp
}

// GetRevisionTimestampOk returns a tuple with the RevisionTimestamp field value
// and a boolean to check if the value has been set.
func (o *Bid) GetRevisionTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionTimestamp, true
}

// SetRevisionTimestamp sets field value
func (o *Bid) SetRevisionTimestamp(v time.Time) {
	o.RevisionTimestamp = v
}

// GetActivation returns the Activation field value if set, zero value otherwise.
func (o *Bid) GetActivation() ActivationType {
	if o == nil || isNil(o.Activation) {
		var ret ActivationType
		return ret
	}
	return *o.Activation
}

// GetActivationOk returns a tuple with the Activation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bid) GetActivationOk() (*ActivationType, bool) {
	if o == nil || isNil(o.Activation) {
		return nil, false
	}
	return o.Activation, true
}

// HasActivation returns a boolean if a field has been set.
func (o *Bid) HasActivation() bool {
	if o != nil && !isNil(o.Activation) {
		return true
	}

	return false
}

// SetActivation gets a reference to the given ActivationType and assigns it to the Activation field.
func (o *Bid) SetActivation(v ActivationType) {
	o.Activation = &v
}

// GetMinQuantity returns the MinQuantity field value if set, zero value otherwise.
func (o *Bid) GetMinQuantity() Quantity {
	if o == nil || isNil(o.MinQuantity) {
		var ret Quantity
		return ret
	}
	return *o.MinQuantity
}

// GetMinQuantityOk returns a tuple with the MinQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bid) GetMinQuantityOk() (*Quantity, bool) {
	if o == nil || isNil(o.MinQuantity) {
		return nil, false
	}
	return o.MinQuantity, true
}

// HasMinQuantity returns a boolean if a field has been set.
func (o *Bid) HasMinQuantity() bool {
	if o != nil && !isNil(o.MinQuantity) {
		return true
	}

	return false
}

// SetMinQuantity gets a reference to the given Quantity and assigns it to the MinQuantity field.
func (o *Bid) SetMinQuantity(v Quantity) {
	o.MinQuantity = &v
}

// GetOfferedQuantity returns the OfferedQuantity field value
func (o *Bid) GetOfferedQuantity() Quantity {
	if o == nil {
		var ret Quantity
		return ret
	}

	return o.OfferedQuantity
}

// GetOfferedQuantityOk returns a tuple with the OfferedQuantity field value
// and a boolean to check if the value has been set.
func (o *Bid) GetOfferedQuantityOk() (*Quantity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferedQuantity, true
}

// SetOfferedQuantity sets field value
func (o *Bid) SetOfferedQuantity(v Quantity) {
	o.OfferedQuantity = v
}

// GetAcceptedQuantity returns the AcceptedQuantity field value if set, zero value otherwise.
func (o *Bid) GetAcceptedQuantity() Quantity {
	if o == nil || isNil(o.AcceptedQuantity) {
		var ret Quantity
		return ret
	}
	return *o.AcceptedQuantity
}

// GetAcceptedQuantityOk returns a tuple with the AcceptedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bid) GetAcceptedQuantityOk() (*Quantity, bool) {
	if o == nil || isNil(o.AcceptedQuantity) {
		return nil, false
	}
	return o.AcceptedQuantity, true
}

// HasAcceptedQuantity returns a boolean if a field has been set.
func (o *Bid) HasAcceptedQuantity() bool {
	if o != nil && !isNil(o.AcceptedQuantity) {
		return true
	}

	return false
}

// SetAcceptedQuantity gets a reference to the given Quantity and assigns it to the AcceptedQuantity field.
func (o *Bid) SetAcceptedQuantity(v Quantity) {
	o.AcceptedQuantity = &v
}

// GetPrices returns the Prices field value
func (o *Bid) GetPrices() []Price {
	if o == nil {
		var ret []Price
		return ret
	}

	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value
// and a boolean to check if the value has been set.
func (o *Bid) GetPricesOk() ([]Price, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prices, true
}

// SetPrices sets field value
func (o *Bid) SetPrices(v []Price) {
	o.Prices = v
}

// GetTechnicalLinkage returns the TechnicalLinkage field value if set, zero value otherwise.
func (o *Bid) GetTechnicalLinkage() TechnicalLinkage {
	if o == nil || isNil(o.TechnicalLinkage) {
		var ret TechnicalLinkage
		return ret
	}
	return *o.TechnicalLinkage
}

// GetTechnicalLinkageOk returns a tuple with the TechnicalLinkage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bid) GetTechnicalLinkageOk() (*TechnicalLinkage, bool) {
	if o == nil || isNil(o.TechnicalLinkage) {
		return nil, false
	}
	return o.TechnicalLinkage, true
}

// HasTechnicalLinkage returns a boolean if a field has been set.
func (o *Bid) HasTechnicalLinkage() bool {
	if o != nil && !isNil(o.TechnicalLinkage) {
		return true
	}

	return false
}

// SetTechnicalLinkage gets a reference to the given TechnicalLinkage and assigns it to the TechnicalLinkage field.
func (o *Bid) SetTechnicalLinkage(v TechnicalLinkage) {
	o.TechnicalLinkage = &v
}

// GetConditionalLinkage returns the ConditionalLinkage field value if set, zero value otherwise.
func (o *Bid) GetConditionalLinkage() ConditionalLinkage {
	if o == nil || isNil(o.ConditionalLinkage) {
		var ret ConditionalLinkage
		return ret
	}
	return *o.ConditionalLinkage
}

// GetConditionalLinkageOk returns a tuple with the ConditionalLinkage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bid) GetConditionalLinkageOk() (*ConditionalLinkage, bool) {
	if o == nil || isNil(o.ConditionalLinkage) {
		return nil, false
	}
	return o.ConditionalLinkage, true
}

// HasConditionalLinkage returns a boolean if a field has been set.
func (o *Bid) HasConditionalLinkage() bool {
	if o != nil && !isNil(o.ConditionalLinkage) {
		return true
	}

	return false
}

// SetConditionalLinkage gets a reference to the given ConditionalLinkage and assigns it to the ConditionalLinkage field.
func (o *Bid) SetConditionalLinkage(v ConditionalLinkage) {
	o.ConditionalLinkage = &v
}

// GetBackupFor returns the BackupFor field value if set, zero value otherwise.
func (o *Bid) GetBackupFor() BalancingServiceProvider {
	if o == nil || isNil(o.BackupFor) {
		var ret BalancingServiceProvider
		return ret
	}
	return *o.BackupFor
}

// GetBackupForOk returns a tuple with the BackupFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bid) GetBackupForOk() (*BalancingServiceProvider, bool) {
	if o == nil || isNil(o.BackupFor) {
		return nil, false
	}
	return o.BackupFor, true
}

// HasBackupFor returns a boolean if a field has been set.
func (o *Bid) HasBackupFor() bool {
	if o != nil && !isNil(o.BackupFor) {
		return true
	}

	return false
}

// SetBackupFor gets a reference to the given BalancingServiceProvider and assigns it to the BackupFor field.
func (o *Bid) SetBackupFor(v BalancingServiceProvider) {
	o.BackupFor = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *Bid) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bid) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *Bid) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *Bid) SetTag(v string) {
	o.Tag = &v
}

func (o Bid) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !isNil(o.ComplexBidId) {
		toSerialize["complexBidId"] = o.ComplexBidId
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	toSerialize["bidType"] = o.BidType
	// skip: tender is readOnly
	toSerialize["productType"] = o.ProductType
	toSerialize["deliveryDate"] = o.DeliveryDate
	toSerialize["market"] = o.Market
	toSerialize["productName"] = o.ProductName
	toSerialize["direction"] = o.Direction
	toSerialize["timeInterval"] = o.TimeInterval
	toSerialize["connectingZone"] = o.ConnectingZone
	toSerialize["resourceProvider"] = o.ResourceProvider
	toSerialize["status"] = o.Status
	toSerialize["receiptTimestamp"] = o.ReceiptTimestamp
	toSerialize["revisionTimestamp"] = o.RevisionTimestamp
	if !isNil(o.Activation) {
		toSerialize["activation"] = o.Activation
	}
	if !isNil(o.MinQuantity) {
		toSerialize["minQuantity"] = o.MinQuantity
	}
	toSerialize["offeredQuantity"] = o.OfferedQuantity
	if !isNil(o.AcceptedQuantity) {
		toSerialize["acceptedQuantity"] = o.AcceptedQuantity
	}
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

type NullableBid struct {
	value *Bid
	isSet bool
}

func (v NullableBid) Get() *Bid {
	return v.value
}

func (v *NullableBid) Set(val *Bid) {
	v.value = val
	v.isSet = true
}

func (v NullableBid) IsSet() bool {
	return v.isSet
}

func (v *NullableBid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBid(val *Bid) *NullableBid {
	return &NullableBid{value: val, isSet: true}
}

func (v NullableBid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


