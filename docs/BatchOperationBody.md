# BatchOperationBody

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

### NewBatchOperationBody

`func NewBatchOperationBody(market ReserveMarket, productType ProductType, bidType BidType, direction Direction, connectingZone ConnectingZone, offeredQuantity Quantity, prices []Price, components []SimpleBid, ) *BatchOperationBody`

NewBatchOperationBody instantiates a new BatchOperationBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchOperationBodyWithDefaults

`func NewBatchOperationBodyWithDefaults() *BatchOperationBody`

NewBatchOperationBodyWithDefaults instantiates a new BatchOperationBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *BatchOperationBody) GetMarket() ReserveMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *BatchOperationBody) GetMarketOk() (*ReserveMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *BatchOperationBody) SetMarket(v ReserveMarket)`

SetMarket sets Market field to given value.


### GetProductType

`func (o *BatchOperationBody) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *BatchOperationBody) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *BatchOperationBody) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetGroupId

`func (o *BatchOperationBody) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BatchOperationBody) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BatchOperationBody) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BatchOperationBody) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetBidType

`func (o *BatchOperationBody) GetBidType() BidType`

GetBidType returns the BidType field if non-nil, zero value otherwise.

### GetBidTypeOk

`func (o *BatchOperationBody) GetBidTypeOk() (*BidType, bool)`

GetBidTypeOk returns a tuple with the BidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidType

`func (o *BatchOperationBody) SetBidType(v BidType)`

SetBidType sets BidType field to given value.


### GetDeliveryDate

`func (o *BatchOperationBody) GetDeliveryDate() string`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *BatchOperationBody) GetDeliveryDateOk() (*string, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *BatchOperationBody) SetDeliveryDate(v string)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *BatchOperationBody) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetProductName

`func (o *BatchOperationBody) GetProductName() ProductName`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *BatchOperationBody) GetProductNameOk() (*ProductName, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *BatchOperationBody) SetProductName(v ProductName)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *BatchOperationBody) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetTimeInterval

`func (o *BatchOperationBody) GetTimeInterval() TimeInterval`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *BatchOperationBody) GetTimeIntervalOk() (*TimeInterval, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *BatchOperationBody) SetTimeInterval(v TimeInterval)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *BatchOperationBody) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.

### GetDirection

`func (o *BatchOperationBody) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BatchOperationBody) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BatchOperationBody) SetDirection(v Direction)`

SetDirection sets Direction field to given value.


### GetConnectingZone

`func (o *BatchOperationBody) GetConnectingZone() ConnectingZone`

GetConnectingZone returns the ConnectingZone field if non-nil, zero value otherwise.

### GetConnectingZoneOk

`func (o *BatchOperationBody) GetConnectingZoneOk() (*ConnectingZone, bool)`

GetConnectingZoneOk returns a tuple with the ConnectingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectingZone

`func (o *BatchOperationBody) SetConnectingZone(v ConnectingZone)`

SetConnectingZone sets ConnectingZone field to given value.


### GetActivation

`func (o *BatchOperationBody) GetActivation() ActivationType`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *BatchOperationBody) GetActivationOk() (*ActivationType, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *BatchOperationBody) SetActivation(v ActivationType)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *BatchOperationBody) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetMinQuantity

`func (o *BatchOperationBody) GetMinQuantity() Quantity`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *BatchOperationBody) GetMinQuantityOk() (*Quantity, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *BatchOperationBody) SetMinQuantity(v Quantity)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *BatchOperationBody) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetOfferedQuantity

`func (o *BatchOperationBody) GetOfferedQuantity() Quantity`

GetOfferedQuantity returns the OfferedQuantity field if non-nil, zero value otherwise.

### GetOfferedQuantityOk

`func (o *BatchOperationBody) GetOfferedQuantityOk() (*Quantity, bool)`

GetOfferedQuantityOk returns a tuple with the OfferedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferedQuantity

`func (o *BatchOperationBody) SetOfferedQuantity(v Quantity)`

SetOfferedQuantity sets OfferedQuantity field to given value.


### GetPrices

`func (o *BatchOperationBody) GetPrices() []Price`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *BatchOperationBody) GetPricesOk() (*[]Price, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *BatchOperationBody) SetPrices(v []Price)`

SetPrices sets Prices field to given value.


### GetTechnicalLinkage

`func (o *BatchOperationBody) GetTechnicalLinkage() TechnicalLinkage`

GetTechnicalLinkage returns the TechnicalLinkage field if non-nil, zero value otherwise.

### GetTechnicalLinkageOk

`func (o *BatchOperationBody) GetTechnicalLinkageOk() (*TechnicalLinkage, bool)`

GetTechnicalLinkageOk returns a tuple with the TechnicalLinkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalLinkage

`func (o *BatchOperationBody) SetTechnicalLinkage(v TechnicalLinkage)`

SetTechnicalLinkage sets TechnicalLinkage field to given value.

### HasTechnicalLinkage

`func (o *BatchOperationBody) HasTechnicalLinkage() bool`

HasTechnicalLinkage returns a boolean if a field has been set.

### GetConditionalLinkage

`func (o *BatchOperationBody) GetConditionalLinkage() ConditionalLinkage`

GetConditionalLinkage returns the ConditionalLinkage field if non-nil, zero value otherwise.

### GetConditionalLinkageOk

`func (o *BatchOperationBody) GetConditionalLinkageOk() (*ConditionalLinkage, bool)`

GetConditionalLinkageOk returns a tuple with the ConditionalLinkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalLinkage

`func (o *BatchOperationBody) SetConditionalLinkage(v ConditionalLinkage)`

SetConditionalLinkage sets ConditionalLinkage field to given value.

### HasConditionalLinkage

`func (o *BatchOperationBody) HasConditionalLinkage() bool`

HasConditionalLinkage returns a boolean if a field has been set.

### GetBackupFor

`func (o *BatchOperationBody) GetBackupFor() BalancingServiceProvider`

GetBackupFor returns the BackupFor field if non-nil, zero value otherwise.

### GetBackupForOk

`func (o *BatchOperationBody) GetBackupForOk() (*BalancingServiceProvider, bool)`

GetBackupForOk returns a tuple with the BackupFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFor

`func (o *BatchOperationBody) SetBackupFor(v BalancingServiceProvider)`

SetBackupFor sets BackupFor field to given value.

### HasBackupFor

`func (o *BatchOperationBody) HasBackupFor() bool`

HasBackupFor returns a boolean if a field has been set.

### GetTag

`func (o *BatchOperationBody) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BatchOperationBody) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BatchOperationBody) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *BatchOperationBody) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetComponents

`func (o *BatchOperationBody) GetComponents() []SimpleBid`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BatchOperationBody) GetComponentsOk() (*[]SimpleBid, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BatchOperationBody) SetComponents(v []SimpleBid)`

SetComponents sets Components field to given value.


### GetComplexBidId

`func (o *BatchOperationBody) GetComplexBidId() string`

GetComplexBidId returns the ComplexBidId field if non-nil, zero value otherwise.

### GetComplexBidIdOk

`func (o *BatchOperationBody) GetComplexBidIdOk() (*string, bool)`

GetComplexBidIdOk returns a tuple with the ComplexBidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexBidId

`func (o *BatchOperationBody) SetComplexBidId(v string)`

SetComplexBidId sets ComplexBidId field to given value.

### HasComplexBidId

`func (o *BatchOperationBody) HasComplexBidId() bool`

HasComplexBidId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


