# SimpleBid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | [**ReserveMarket**](ReserveMarket.md) |  | 
**ProductType** | [**ProductType**](ProductType.md) |  | 
**GroupId** | Pointer to **string** | Bid identification | [optional] 
**BidType** | Pointer to [**BidType**](BidType.md) |  | [optional] [default to SIMPLE]
**DeliveryDate** | Pointer to **string** | Delivery day of offered control reserve / energy (ISO 8601 format YYYY-MM-DD). | [optional] 
**ProductName** | Pointer to [**ProductName**](ProductName.md) |  | [optional] 
**TimeInterval** | Pointer to [**TimeInterval**](TimeInterval.md) |  | [optional] 
**Direction** | Pointer to [**Direction**](Direction.md) |  | [optional] 
**ConnectingZone** | [**ConnectingZone**](ConnectingZone.md) |  | 
**Activation** | Pointer to [**ActivationType**](ActivationType.md) |  | [optional] [default to DIRECT]
**MinQuantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**OfferedQuantity** | [**Quantity**](Quantity.md) |  | 
**Prices** | [**[]Price**](Price.md) | Price information of the bid. | 
**TechnicalLinkage** | Pointer to [**TechnicalLinkage**](TechnicalLinkage.md) |  | [optional] 
**ConditionalLinkage** | Pointer to [**ConditionalLinkage**](ConditionalLinkage.md) |  | [optional] 
**BackupFor** | Pointer to [**BalancingServiceProvider**](BalancingServiceProvider.md) |  | [optional] 
**Tag** | Pointer to **string** | Custom tag for bids | [optional] 

## Methods

### NewSimpleBid

`func NewSimpleBid(market ReserveMarket, productType ProductType, connectingZone ConnectingZone, offeredQuantity Quantity, prices []Price, ) *SimpleBid`

NewSimpleBid instantiates a new SimpleBid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleBidWithDefaults

`func NewSimpleBidWithDefaults() *SimpleBid`

NewSimpleBidWithDefaults instantiates a new SimpleBid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *SimpleBid) GetMarket() ReserveMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *SimpleBid) GetMarketOk() (*ReserveMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *SimpleBid) SetMarket(v ReserveMarket)`

SetMarket sets Market field to given value.


### GetProductType

`func (o *SimpleBid) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *SimpleBid) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *SimpleBid) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetGroupId

`func (o *SimpleBid) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SimpleBid) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SimpleBid) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SimpleBid) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetBidType

`func (o *SimpleBid) GetBidType() BidType`

GetBidType returns the BidType field if non-nil, zero value otherwise.

### GetBidTypeOk

`func (o *SimpleBid) GetBidTypeOk() (*BidType, bool)`

GetBidTypeOk returns a tuple with the BidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidType

`func (o *SimpleBid) SetBidType(v BidType)`

SetBidType sets BidType field to given value.

### HasBidType

`func (o *SimpleBid) HasBidType() bool`

HasBidType returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *SimpleBid) GetDeliveryDate() string`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *SimpleBid) GetDeliveryDateOk() (*string, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *SimpleBid) SetDeliveryDate(v string)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *SimpleBid) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetProductName

`func (o *SimpleBid) GetProductName() ProductName`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *SimpleBid) GetProductNameOk() (*ProductName, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *SimpleBid) SetProductName(v ProductName)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *SimpleBid) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetTimeInterval

`func (o *SimpleBid) GetTimeInterval() TimeInterval`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *SimpleBid) GetTimeIntervalOk() (*TimeInterval, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *SimpleBid) SetTimeInterval(v TimeInterval)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *SimpleBid) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.

### GetDirection

`func (o *SimpleBid) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SimpleBid) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SimpleBid) SetDirection(v Direction)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SimpleBid) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetConnectingZone

`func (o *SimpleBid) GetConnectingZone() ConnectingZone`

GetConnectingZone returns the ConnectingZone field if non-nil, zero value otherwise.

### GetConnectingZoneOk

`func (o *SimpleBid) GetConnectingZoneOk() (*ConnectingZone, bool)`

GetConnectingZoneOk returns a tuple with the ConnectingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectingZone

`func (o *SimpleBid) SetConnectingZone(v ConnectingZone)`

SetConnectingZone sets ConnectingZone field to given value.


### GetActivation

`func (o *SimpleBid) GetActivation() ActivationType`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *SimpleBid) GetActivationOk() (*ActivationType, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *SimpleBid) SetActivation(v ActivationType)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *SimpleBid) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetMinQuantity

`func (o *SimpleBid) GetMinQuantity() Quantity`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *SimpleBid) GetMinQuantityOk() (*Quantity, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *SimpleBid) SetMinQuantity(v Quantity)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *SimpleBid) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetOfferedQuantity

`func (o *SimpleBid) GetOfferedQuantity() Quantity`

GetOfferedQuantity returns the OfferedQuantity field if non-nil, zero value otherwise.

### GetOfferedQuantityOk

`func (o *SimpleBid) GetOfferedQuantityOk() (*Quantity, bool)`

GetOfferedQuantityOk returns a tuple with the OfferedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferedQuantity

`func (o *SimpleBid) SetOfferedQuantity(v Quantity)`

SetOfferedQuantity sets OfferedQuantity field to given value.


### GetPrices

`func (o *SimpleBid) GetPrices() []Price`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *SimpleBid) GetPricesOk() (*[]Price, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *SimpleBid) SetPrices(v []Price)`

SetPrices sets Prices field to given value.


### GetTechnicalLinkage

`func (o *SimpleBid) GetTechnicalLinkage() TechnicalLinkage`

GetTechnicalLinkage returns the TechnicalLinkage field if non-nil, zero value otherwise.

### GetTechnicalLinkageOk

`func (o *SimpleBid) GetTechnicalLinkageOk() (*TechnicalLinkage, bool)`

GetTechnicalLinkageOk returns a tuple with the TechnicalLinkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalLinkage

`func (o *SimpleBid) SetTechnicalLinkage(v TechnicalLinkage)`

SetTechnicalLinkage sets TechnicalLinkage field to given value.

### HasTechnicalLinkage

`func (o *SimpleBid) HasTechnicalLinkage() bool`

HasTechnicalLinkage returns a boolean if a field has been set.

### GetConditionalLinkage

`func (o *SimpleBid) GetConditionalLinkage() ConditionalLinkage`

GetConditionalLinkage returns the ConditionalLinkage field if non-nil, zero value otherwise.

### GetConditionalLinkageOk

`func (o *SimpleBid) GetConditionalLinkageOk() (*ConditionalLinkage, bool)`

GetConditionalLinkageOk returns a tuple with the ConditionalLinkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalLinkage

`func (o *SimpleBid) SetConditionalLinkage(v ConditionalLinkage)`

SetConditionalLinkage sets ConditionalLinkage field to given value.

### HasConditionalLinkage

`func (o *SimpleBid) HasConditionalLinkage() bool`

HasConditionalLinkage returns a boolean if a field has been set.

### GetBackupFor

`func (o *SimpleBid) GetBackupFor() BalancingServiceProvider`

GetBackupFor returns the BackupFor field if non-nil, zero value otherwise.

### GetBackupForOk

`func (o *SimpleBid) GetBackupForOk() (*BalancingServiceProvider, bool)`

GetBackupForOk returns a tuple with the BackupFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFor

`func (o *SimpleBid) SetBackupFor(v BalancingServiceProvider)`

SetBackupFor sets BackupFor field to given value.

### HasBackupFor

`func (o *SimpleBid) HasBackupFor() bool`

HasBackupFor returns a boolean if a field has been set.

### GetTag

`func (o *SimpleBid) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *SimpleBid) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *SimpleBid) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *SimpleBid) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


