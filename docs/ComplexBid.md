# ComplexBid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | [**ReserveMarket**](ReserveMarket.md) |  | 
**ProductType** | [**ProductType**](ProductType.md) |  | 
**BidType** | [**BidType**](BidType.md) |  | [default to SIMPLE]
**TimeInterval** | Pointer to [**TimeInterval**](TimeInterval.md) |  | [optional] 
**ConnectingZone** | Pointer to [**ConnectingZone**](ConnectingZone.md) |  | [optional] 
**Activation** | Pointer to [**ActivationType**](ActivationType.md) |  | [optional] [default to DIRECT]
**TechnicalLinkage** | Pointer to [**TechnicalLinkage**](TechnicalLinkage.md) |  | [optional] 
**BackupFor** | Pointer to [**BalancingServiceProvider**](BalancingServiceProvider.md) |  | [optional] 
**Tag** | Pointer to **string** | Custom tag for bids | [optional] 
**Components** | [**[]ComplexBidComponent**](ComplexBidComponent.md) |  | 

## Methods

### NewComplexBid

`func NewComplexBid(market ReserveMarket, productType ProductType, bidType BidType, components []ComplexBidComponent, ) *ComplexBid`

NewComplexBid instantiates a new ComplexBid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplexBidWithDefaults

`func NewComplexBidWithDefaults() *ComplexBid`

NewComplexBidWithDefaults instantiates a new ComplexBid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *ComplexBid) GetMarket() ReserveMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *ComplexBid) GetMarketOk() (*ReserveMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *ComplexBid) SetMarket(v ReserveMarket)`

SetMarket sets Market field to given value.


### GetProductType

`func (o *ComplexBid) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ComplexBid) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ComplexBid) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetBidType

`func (o *ComplexBid) GetBidType() BidType`

GetBidType returns the BidType field if non-nil, zero value otherwise.

### GetBidTypeOk

`func (o *ComplexBid) GetBidTypeOk() (*BidType, bool)`

GetBidTypeOk returns a tuple with the BidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidType

`func (o *ComplexBid) SetBidType(v BidType)`

SetBidType sets BidType field to given value.


### GetTimeInterval

`func (o *ComplexBid) GetTimeInterval() TimeInterval`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *ComplexBid) GetTimeIntervalOk() (*TimeInterval, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *ComplexBid) SetTimeInterval(v TimeInterval)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *ComplexBid) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.

### GetConnectingZone

`func (o *ComplexBid) GetConnectingZone() ConnectingZone`

GetConnectingZone returns the ConnectingZone field if non-nil, zero value otherwise.

### GetConnectingZoneOk

`func (o *ComplexBid) GetConnectingZoneOk() (*ConnectingZone, bool)`

GetConnectingZoneOk returns a tuple with the ConnectingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectingZone

`func (o *ComplexBid) SetConnectingZone(v ConnectingZone)`

SetConnectingZone sets ConnectingZone field to given value.

### HasConnectingZone

`func (o *ComplexBid) HasConnectingZone() bool`

HasConnectingZone returns a boolean if a field has been set.

### GetActivation

`func (o *ComplexBid) GetActivation() ActivationType`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *ComplexBid) GetActivationOk() (*ActivationType, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *ComplexBid) SetActivation(v ActivationType)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *ComplexBid) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetTechnicalLinkage

`func (o *ComplexBid) GetTechnicalLinkage() TechnicalLinkage`

GetTechnicalLinkage returns the TechnicalLinkage field if non-nil, zero value otherwise.

### GetTechnicalLinkageOk

`func (o *ComplexBid) GetTechnicalLinkageOk() (*TechnicalLinkage, bool)`

GetTechnicalLinkageOk returns a tuple with the TechnicalLinkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalLinkage

`func (o *ComplexBid) SetTechnicalLinkage(v TechnicalLinkage)`

SetTechnicalLinkage sets TechnicalLinkage field to given value.

### HasTechnicalLinkage

`func (o *ComplexBid) HasTechnicalLinkage() bool`

HasTechnicalLinkage returns a boolean if a field has been set.

### GetBackupFor

`func (o *ComplexBid) GetBackupFor() BalancingServiceProvider`

GetBackupFor returns the BackupFor field if non-nil, zero value otherwise.

### GetBackupForOk

`func (o *ComplexBid) GetBackupForOk() (*BalancingServiceProvider, bool)`

GetBackupForOk returns a tuple with the BackupFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFor

`func (o *ComplexBid) SetBackupFor(v BalancingServiceProvider)`

SetBackupFor sets BackupFor field to given value.

### HasBackupFor

`func (o *ComplexBid) HasBackupFor() bool`

HasBackupFor returns a boolean if a field has been set.

### GetTag

`func (o *ComplexBid) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ComplexBid) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ComplexBid) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ComplexBid) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetComponents

`func (o *ComplexBid) GetComponents() []ComplexBidComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ComplexBid) GetComponentsOk() (*[]ComplexBidComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ComplexBid) SetComponents(v []ComplexBidComponent)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


