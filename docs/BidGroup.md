# BidGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | [**ReserveMarket**](ReserveMarket.md) |  | 
**ProductType** | [**ProductType**](ProductType.md) |  | 
**BidType** | [**BidType**](BidType.md) |  | [default to SIMPLE]
**TimeInterval** | Pointer to [**TimeInterval**](TimeInterval.md) |  | [optional] 
**Tag** | Pointer to **string** | Custom tag for bids | [optional] 
**Components** | [**[]SimpleBid**](SimpleBid.md) |  | 

## Methods

### NewBidGroup

`func NewBidGroup(market ReserveMarket, productType ProductType, bidType BidType, components []SimpleBid, ) *BidGroup`

NewBidGroup instantiates a new BidGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBidGroupWithDefaults

`func NewBidGroupWithDefaults() *BidGroup`

NewBidGroupWithDefaults instantiates a new BidGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *BidGroup) GetMarket() ReserveMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *BidGroup) GetMarketOk() (*ReserveMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *BidGroup) SetMarket(v ReserveMarket)`

SetMarket sets Market field to given value.


### GetProductType

`func (o *BidGroup) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *BidGroup) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *BidGroup) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetBidType

`func (o *BidGroup) GetBidType() BidType`

GetBidType returns the BidType field if non-nil, zero value otherwise.

### GetBidTypeOk

`func (o *BidGroup) GetBidTypeOk() (*BidType, bool)`

GetBidTypeOk returns a tuple with the BidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidType

`func (o *BidGroup) SetBidType(v BidType)`

SetBidType sets BidType field to given value.


### GetTimeInterval

`func (o *BidGroup) GetTimeInterval() TimeInterval`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *BidGroup) GetTimeIntervalOk() (*TimeInterval, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *BidGroup) SetTimeInterval(v TimeInterval)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *BidGroup) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.

### GetTag

`func (o *BidGroup) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BidGroup) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BidGroup) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *BidGroup) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetComponents

`func (o *BidGroup) GetComponents() []SimpleBid`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BidGroup) GetComponentsOk() (*[]SimpleBid, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BidGroup) SetComponents(v []SimpleBid)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


