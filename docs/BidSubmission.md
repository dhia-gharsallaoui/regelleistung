# BidSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | [**ReserveMarket**](ReserveMarket.md) |  | 
**ProductType** | [**ProductType**](ProductType.md) |  | 
**GroupId** | Pointer to **string** | Bid identification | [optional] 
**BidType** | [**BidType**](BidType.md) |  | [default to SIMPLE]
**DeliveryDate** | Pointer to **string** | Delivery day of offered control reserve / energy (ISO 8601 format YYYY-MM-DD). | [optional] 
**ProductName** | Pointer to [**ProductName**](ProductName.md) |  | [optional] 
**TimeInterval** | Pointer to [**TimeInterval**](TimeInterval.md) |  | [optional] 
**Direction** | [**Direction**](Direction.md) |  | 
**ConnectingZone** | [**ConnectingZone**](ConnectingZone.md) |  | 
**Activation** | Pointer to [**ActivationType**](ActivationType.md) |  | [optional] [default to DIRECT]
**MinQuantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**OfferedQuantity** | [**Quantity**](Quantity.md) |  | 
**Prices** | [**[]Price**](Price.md) | Price information of the bid. | 
**TechnicalLinkage** | Pointer to [**TechnicalLinkage**](TechnicalLinkage.md) |  | [optional] 
**ConditionalLinkage** | Pointer to [**ConditionalLinkage**](ConditionalLinkage.md) |  | [optional] 
**BackupFor** | Pointer to [**BalancingServiceProvider**](BalancingServiceProvider.md) |  | [optional] 
**Tag** | Pointer to **string** | Custom tag for bids | [optional] 
**Components** | [**[]SimpleBid**](SimpleBid.md) |  | 
**ComplexBidId** | Pointer to **string** | Bid identification | [optional] 

## Methods

### NewBidSubmission

`func NewBidSubmission(market ReserveMarket, productType ProductType, bidType BidType, direction Direction, connectingZone ConnectingZone, offeredQuantity Quantity, prices []Price, components []SimpleBid, ) *BidSubmission`

NewBidSubmission instantiates a new BidSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBidSubmissionWithDefaults

`func NewBidSubmissionWithDefaults() *BidSubmission`

NewBidSubmissionWithDefaults instantiates a new BidSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *BidSubmission) GetMarket() ReserveMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *BidSubmission) GetMarketOk() (*ReserveMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *BidSubmission) SetMarket(v ReserveMarket)`

SetMarket sets Market field to given value.


### GetProductType

`func (o *BidSubmission) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *BidSubmission) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *BidSubmission) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetGroupId

`func (o *BidSubmission) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BidSubmission) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BidSubmission) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BidSubmission) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetBidType

`func (o *BidSubmission) GetBidType() BidType`

GetBidType returns the BidType field if non-nil, zero value otherwise.

### GetBidTypeOk

`func (o *BidSubmission) GetBidTypeOk() (*BidType, bool)`

GetBidTypeOk returns a tuple with the BidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidType

`func (o *BidSubmission) SetBidType(v BidType)`

SetBidType sets BidType field to given value.


### GetDeliveryDate

`func (o *BidSubmission) GetDeliveryDate() string`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *BidSubmission) GetDeliveryDateOk() (*string, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *BidSubmission) SetDeliveryDate(v string)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *BidSubmission) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetProductName

`func (o *BidSubmission) GetProductName() ProductName`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *BidSubmission) GetProductNameOk() (*ProductName, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *BidSubmission) SetProductName(v ProductName)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *BidSubmission) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetTimeInterval

`func (o *BidSubmission) GetTimeInterval() TimeInterval`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *BidSubmission) GetTimeIntervalOk() (*TimeInterval, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *BidSubmission) SetTimeInterval(v TimeInterval)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *BidSubmission) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.

### GetDirection

`func (o *BidSubmission) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BidSubmission) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BidSubmission) SetDirection(v Direction)`

SetDirection sets Direction field to given value.


### GetConnectingZone

`func (o *BidSubmission) GetConnectingZone() ConnectingZone`

GetConnectingZone returns the ConnectingZone field if non-nil, zero value otherwise.

### GetConnectingZoneOk

`func (o *BidSubmission) GetConnectingZoneOk() (*ConnectingZone, bool)`

GetConnectingZoneOk returns a tuple with the ConnectingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectingZone

`func (o *BidSubmission) SetConnectingZone(v ConnectingZone)`

SetConnectingZone sets ConnectingZone field to given value.


### GetActivation

`func (o *BidSubmission) GetActivation() ActivationType`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *BidSubmission) GetActivationOk() (*ActivationType, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *BidSubmission) SetActivation(v ActivationType)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *BidSubmission) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetMinQuantity

`func (o *BidSubmission) GetMinQuantity() Quantity`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *BidSubmission) GetMinQuantityOk() (*Quantity, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *BidSubmission) SetMinQuantity(v Quantity)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *BidSubmission) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetOfferedQuantity

`func (o *BidSubmission) GetOfferedQuantity() Quantity`

GetOfferedQuantity returns the OfferedQuantity field if non-nil, zero value otherwise.

### GetOfferedQuantityOk

`func (o *BidSubmission) GetOfferedQuantityOk() (*Quantity, bool)`

GetOfferedQuantityOk returns a tuple with the OfferedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferedQuantity

`func (o *BidSubmission) SetOfferedQuantity(v Quantity)`

SetOfferedQuantity sets OfferedQuantity field to given value.


### GetPrices

`func (o *BidSubmission) GetPrices() []Price`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *BidSubmission) GetPricesOk() (*[]Price, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *BidSubmission) SetPrices(v []Price)`

SetPrices sets Prices field to given value.


### GetTechnicalLinkage

`func (o *BidSubmission) GetTechnicalLinkage() TechnicalLinkage`

GetTechnicalLinkage returns the TechnicalLinkage field if non-nil, zero value otherwise.

### GetTechnicalLinkageOk

`func (o *BidSubmission) GetTechnicalLinkageOk() (*TechnicalLinkage, bool)`

GetTechnicalLinkageOk returns a tuple with the TechnicalLinkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalLinkage

`func (o *BidSubmission) SetTechnicalLinkage(v TechnicalLinkage)`

SetTechnicalLinkage sets TechnicalLinkage field to given value.

### HasTechnicalLinkage

`func (o *BidSubmission) HasTechnicalLinkage() bool`

HasTechnicalLinkage returns a boolean if a field has been set.

### GetConditionalLinkage

`func (o *BidSubmission) GetConditionalLinkage() ConditionalLinkage`

GetConditionalLinkage returns the ConditionalLinkage field if non-nil, zero value otherwise.

### GetConditionalLinkageOk

`func (o *BidSubmission) GetConditionalLinkageOk() (*ConditionalLinkage, bool)`

GetConditionalLinkageOk returns a tuple with the ConditionalLinkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalLinkage

`func (o *BidSubmission) SetConditionalLinkage(v ConditionalLinkage)`

SetConditionalLinkage sets ConditionalLinkage field to given value.

### HasConditionalLinkage

`func (o *BidSubmission) HasConditionalLinkage() bool`

HasConditionalLinkage returns a boolean if a field has been set.

### GetBackupFor

`func (o *BidSubmission) GetBackupFor() BalancingServiceProvider`

GetBackupFor returns the BackupFor field if non-nil, zero value otherwise.

### GetBackupForOk

`func (o *BidSubmission) GetBackupForOk() (*BalancingServiceProvider, bool)`

GetBackupForOk returns a tuple with the BackupFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFor

`func (o *BidSubmission) SetBackupFor(v BalancingServiceProvider)`

SetBackupFor sets BackupFor field to given value.

### HasBackupFor

`func (o *BidSubmission) HasBackupFor() bool`

HasBackupFor returns a boolean if a field has been set.

### GetTag

`func (o *BidSubmission) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BidSubmission) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BidSubmission) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *BidSubmission) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetComponents

`func (o *BidSubmission) GetComponents() []SimpleBid`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BidSubmission) GetComponentsOk() (*[]SimpleBid, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BidSubmission) SetComponents(v []SimpleBid)`

SetComponents sets Components field to given value.


### GetComplexBidId

`func (o *BidSubmission) GetComplexBidId() string`

GetComplexBidId returns the ComplexBidId field if non-nil, zero value otherwise.

### GetComplexBidIdOk

`func (o *BidSubmission) GetComplexBidIdOk() (*string, bool)`

GetComplexBidIdOk returns a tuple with the ComplexBidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexBidId

`func (o *BidSubmission) SetComplexBidId(v string)`

SetComplexBidId sets ComplexBidId field to given value.

### HasComplexBidId

`func (o *BidSubmission) HasComplexBidId() bool`

HasComplexBidId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


