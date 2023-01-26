# Bid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Bid identification | 
**ComplexBidId** | Pointer to **string** | Bid identification | [optional] 
**GroupId** | Pointer to **string** | Bid identification | [optional] 
**BidType** | [**BidType**](BidType.md) |  | [default to SIMPLE]
**Tender** | **string** | The tender identification is structured as follows:   &#x60;[productType]-[deliveryDate]-[market]-[sequenceNumber]&#x60;    * &#x60;productType&#x60; - one of &#x60;fcr&#x60;, &#x60;afrr&#x60; or &#x60;mfrr&#x60;   * &#x60;deliveryDate&#x60; - ISO 8601 date   * &#x60;market&#x60; - one of &#x60;cm&#x60;(capacity market) or &#x60;em&#x60;(energy market)   * &#x60;sequenceNumber&#x60; - An ascending number for the identification of different tender runs or time slices.   For more details on the new tender id format, see the [reference guide.](/docs/guide#tender-id).  | [readonly] 
**ProductType** | [**ProductType**](ProductType.md) |  | 
**DeliveryDate** | **string** | Delivery day of offered control reserve / energy (ISO 8601 format YYYY-MM-DD). | 
**Market** | [**ReserveMarket**](ReserveMarket.md) |  | 
**ProductName** | [**ProductName**](ProductName.md) |  | 
**Direction** | [**Direction**](Direction.md) |  | 
**TimeInterval** | [**TimeInterval**](TimeInterval.md) |  | 
**ConnectingZone** | [**ConnectingZone**](ConnectingZone.md) |  | 
**ResourceProvider** | **string** | Representation of EIC. The coding scheme is the Energy Identification Coding Scheme (EIC), maintained by ENTSO-e. | 
**Status** | **string** |  | 
**ReceiptTimestamp** | **time.Time** | Receipt time of the bid. | 
**RevisionTimestamp** | **time.Time** | Revision timestamp of the bid. | 
**Activation** | Pointer to [**ActivationType**](ActivationType.md) |  | [optional] [default to DIRECT]
**MinQuantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**OfferedQuantity** | [**Quantity**](Quantity.md) |  | 
**AcceptedQuantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**Prices** | [**[]Price**](Price.md) | Price information of the bid. | 
**TechnicalLinkage** | Pointer to [**TechnicalLinkage**](TechnicalLinkage.md) |  | [optional] 
**ConditionalLinkage** | Pointer to [**ConditionalLinkage**](ConditionalLinkage.md) |  | [optional] 
**BackupFor** | Pointer to [**BalancingServiceProvider**](BalancingServiceProvider.md) |  | [optional] 
**Tag** | Pointer to **string** | Custom tag for bids | [optional] 

## Methods

### NewBid

`func NewBid(id string, bidType BidType, tender string, productType ProductType, deliveryDate string, market ReserveMarket, productName ProductName, direction Direction, timeInterval TimeInterval, connectingZone ConnectingZone, resourceProvider string, status string, receiptTimestamp time.Time, revisionTimestamp time.Time, offeredQuantity Quantity, prices []Price, ) *Bid`

NewBid instantiates a new Bid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBidWithDefaults

`func NewBidWithDefaults() *Bid`

NewBidWithDefaults instantiates a new Bid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bid) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bid) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bid) SetId(v string)`

SetId sets Id field to given value.


### GetComplexBidId

`func (o *Bid) GetComplexBidId() string`

GetComplexBidId returns the ComplexBidId field if non-nil, zero value otherwise.

### GetComplexBidIdOk

`func (o *Bid) GetComplexBidIdOk() (*string, bool)`

GetComplexBidIdOk returns a tuple with the ComplexBidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexBidId

`func (o *Bid) SetComplexBidId(v string)`

SetComplexBidId sets ComplexBidId field to given value.

### HasComplexBidId

`func (o *Bid) HasComplexBidId() bool`

HasComplexBidId returns a boolean if a field has been set.

### GetGroupId

`func (o *Bid) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Bid) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Bid) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Bid) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetBidType

`func (o *Bid) GetBidType() BidType`

GetBidType returns the BidType field if non-nil, zero value otherwise.

### GetBidTypeOk

`func (o *Bid) GetBidTypeOk() (*BidType, bool)`

GetBidTypeOk returns a tuple with the BidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidType

`func (o *Bid) SetBidType(v BidType)`

SetBidType sets BidType field to given value.


### GetTender

`func (o *Bid) GetTender() string`

GetTender returns the Tender field if non-nil, zero value otherwise.

### GetTenderOk

`func (o *Bid) GetTenderOk() (*string, bool)`

GetTenderOk returns a tuple with the Tender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTender

`func (o *Bid) SetTender(v string)`

SetTender sets Tender field to given value.


### GetProductType

`func (o *Bid) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *Bid) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *Bid) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetDeliveryDate

`func (o *Bid) GetDeliveryDate() string`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *Bid) GetDeliveryDateOk() (*string, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *Bid) SetDeliveryDate(v string)`

SetDeliveryDate sets DeliveryDate field to given value.


### GetMarket

`func (o *Bid) GetMarket() ReserveMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *Bid) GetMarketOk() (*ReserveMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *Bid) SetMarket(v ReserveMarket)`

SetMarket sets Market field to given value.


### GetProductName

`func (o *Bid) GetProductName() ProductName`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *Bid) GetProductNameOk() (*ProductName, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *Bid) SetProductName(v ProductName)`

SetProductName sets ProductName field to given value.


### GetDirection

`func (o *Bid) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Bid) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Bid) SetDirection(v Direction)`

SetDirection sets Direction field to given value.


### GetTimeInterval

`func (o *Bid) GetTimeInterval() TimeInterval`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *Bid) GetTimeIntervalOk() (*TimeInterval, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *Bid) SetTimeInterval(v TimeInterval)`

SetTimeInterval sets TimeInterval field to given value.


### GetConnectingZone

`func (o *Bid) GetConnectingZone() ConnectingZone`

GetConnectingZone returns the ConnectingZone field if non-nil, zero value otherwise.

### GetConnectingZoneOk

`func (o *Bid) GetConnectingZoneOk() (*ConnectingZone, bool)`

GetConnectingZoneOk returns a tuple with the ConnectingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectingZone

`func (o *Bid) SetConnectingZone(v ConnectingZone)`

SetConnectingZone sets ConnectingZone field to given value.


### GetResourceProvider

`func (o *Bid) GetResourceProvider() string`

GetResourceProvider returns the ResourceProvider field if non-nil, zero value otherwise.

### GetResourceProviderOk

`func (o *Bid) GetResourceProviderOk() (*string, bool)`

GetResourceProviderOk returns a tuple with the ResourceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProvider

`func (o *Bid) SetResourceProvider(v string)`

SetResourceProvider sets ResourceProvider field to given value.


### GetStatus

`func (o *Bid) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Bid) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Bid) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReceiptTimestamp

`func (o *Bid) GetReceiptTimestamp() time.Time`

GetReceiptTimestamp returns the ReceiptTimestamp field if non-nil, zero value otherwise.

### GetReceiptTimestampOk

`func (o *Bid) GetReceiptTimestampOk() (*time.Time, bool)`

GetReceiptTimestampOk returns a tuple with the ReceiptTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptTimestamp

`func (o *Bid) SetReceiptTimestamp(v time.Time)`

SetReceiptTimestamp sets ReceiptTimestamp field to given value.


### GetRevisionTimestamp

`func (o *Bid) GetRevisionTimestamp() time.Time`

GetRevisionTimestamp returns the RevisionTimestamp field if non-nil, zero value otherwise.

### GetRevisionTimestampOk

`func (o *Bid) GetRevisionTimestampOk() (*time.Time, bool)`

GetRevisionTimestampOk returns a tuple with the RevisionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionTimestamp

`func (o *Bid) SetRevisionTimestamp(v time.Time)`

SetRevisionTimestamp sets RevisionTimestamp field to given value.


### GetActivation

`func (o *Bid) GetActivation() ActivationType`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *Bid) GetActivationOk() (*ActivationType, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *Bid) SetActivation(v ActivationType)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *Bid) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetMinQuantity

`func (o *Bid) GetMinQuantity() Quantity`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *Bid) GetMinQuantityOk() (*Quantity, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *Bid) SetMinQuantity(v Quantity)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *Bid) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetOfferedQuantity

`func (o *Bid) GetOfferedQuantity() Quantity`

GetOfferedQuantity returns the OfferedQuantity field if non-nil, zero value otherwise.

### GetOfferedQuantityOk

`func (o *Bid) GetOfferedQuantityOk() (*Quantity, bool)`

GetOfferedQuantityOk returns a tuple with the OfferedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferedQuantity

`func (o *Bid) SetOfferedQuantity(v Quantity)`

SetOfferedQuantity sets OfferedQuantity field to given value.


### GetAcceptedQuantity

`func (o *Bid) GetAcceptedQuantity() Quantity`

GetAcceptedQuantity returns the AcceptedQuantity field if non-nil, zero value otherwise.

### GetAcceptedQuantityOk

`func (o *Bid) GetAcceptedQuantityOk() (*Quantity, bool)`

GetAcceptedQuantityOk returns a tuple with the AcceptedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedQuantity

`func (o *Bid) SetAcceptedQuantity(v Quantity)`

SetAcceptedQuantity sets AcceptedQuantity field to given value.

### HasAcceptedQuantity

`func (o *Bid) HasAcceptedQuantity() bool`

HasAcceptedQuantity returns a boolean if a field has been set.

### GetPrices

`func (o *Bid) GetPrices() []Price`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *Bid) GetPricesOk() (*[]Price, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *Bid) SetPrices(v []Price)`

SetPrices sets Prices field to given value.


### GetTechnicalLinkage

`func (o *Bid) GetTechnicalLinkage() TechnicalLinkage`

GetTechnicalLinkage returns the TechnicalLinkage field if non-nil, zero value otherwise.

### GetTechnicalLinkageOk

`func (o *Bid) GetTechnicalLinkageOk() (*TechnicalLinkage, bool)`

GetTechnicalLinkageOk returns a tuple with the TechnicalLinkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalLinkage

`func (o *Bid) SetTechnicalLinkage(v TechnicalLinkage)`

SetTechnicalLinkage sets TechnicalLinkage field to given value.

### HasTechnicalLinkage

`func (o *Bid) HasTechnicalLinkage() bool`

HasTechnicalLinkage returns a boolean if a field has been set.

### GetConditionalLinkage

`func (o *Bid) GetConditionalLinkage() ConditionalLinkage`

GetConditionalLinkage returns the ConditionalLinkage field if non-nil, zero value otherwise.

### GetConditionalLinkageOk

`func (o *Bid) GetConditionalLinkageOk() (*ConditionalLinkage, bool)`

GetConditionalLinkageOk returns a tuple with the ConditionalLinkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalLinkage

`func (o *Bid) SetConditionalLinkage(v ConditionalLinkage)`

SetConditionalLinkage sets ConditionalLinkage field to given value.

### HasConditionalLinkage

`func (o *Bid) HasConditionalLinkage() bool`

HasConditionalLinkage returns a boolean if a field has been set.

### GetBackupFor

`func (o *Bid) GetBackupFor() BalancingServiceProvider`

GetBackupFor returns the BackupFor field if non-nil, zero value otherwise.

### GetBackupForOk

`func (o *Bid) GetBackupForOk() (*BalancingServiceProvider, bool)`

GetBackupForOk returns a tuple with the BackupFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFor

`func (o *Bid) SetBackupFor(v BalancingServiceProvider)`

SetBackupFor sets BackupFor field to given value.

### HasBackupFor

`func (o *Bid) HasBackupFor() bool`

HasBackupFor returns a boolean if a field has been set.

### GetTag

`func (o *Bid) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Bid) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Bid) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Bid) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


