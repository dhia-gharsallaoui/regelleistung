# ComplexBidComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BidType** | Pointer to [**BidType**](BidType.md) |  | [optional] [default to SIMPLE]
**ComplexBidId** | Pointer to **string** | Bid identification | [optional] 
**Direction** | [**Direction**](Direction.md) |  | 
**MinQuantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**OfferedQuantity** | [**Quantity**](Quantity.md) |  | 
**Prices** | [**[]Price**](Price.md) | Price information of the bid. | 
**Tag** | Pointer to **string** | Custom tag for bids | [optional] 

## Methods

### NewComplexBidComponent

`func NewComplexBidComponent(direction Direction, offeredQuantity Quantity, prices []Price, ) *ComplexBidComponent`

NewComplexBidComponent instantiates a new ComplexBidComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplexBidComponentWithDefaults

`func NewComplexBidComponentWithDefaults() *ComplexBidComponent`

NewComplexBidComponentWithDefaults instantiates a new ComplexBidComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBidType

`func (o *ComplexBidComponent) GetBidType() BidType`

GetBidType returns the BidType field if non-nil, zero value otherwise.

### GetBidTypeOk

`func (o *ComplexBidComponent) GetBidTypeOk() (*BidType, bool)`

GetBidTypeOk returns a tuple with the BidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidType

`func (o *ComplexBidComponent) SetBidType(v BidType)`

SetBidType sets BidType field to given value.

### HasBidType

`func (o *ComplexBidComponent) HasBidType() bool`

HasBidType returns a boolean if a field has been set.

### GetComplexBidId

`func (o *ComplexBidComponent) GetComplexBidId() string`

GetComplexBidId returns the ComplexBidId field if non-nil, zero value otherwise.

### GetComplexBidIdOk

`func (o *ComplexBidComponent) GetComplexBidIdOk() (*string, bool)`

GetComplexBidIdOk returns a tuple with the ComplexBidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexBidId

`func (o *ComplexBidComponent) SetComplexBidId(v string)`

SetComplexBidId sets ComplexBidId field to given value.

### HasComplexBidId

`func (o *ComplexBidComponent) HasComplexBidId() bool`

HasComplexBidId returns a boolean if a field has been set.

### GetDirection

`func (o *ComplexBidComponent) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ComplexBidComponent) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ComplexBidComponent) SetDirection(v Direction)`

SetDirection sets Direction field to given value.


### GetMinQuantity

`func (o *ComplexBidComponent) GetMinQuantity() Quantity`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *ComplexBidComponent) GetMinQuantityOk() (*Quantity, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *ComplexBidComponent) SetMinQuantity(v Quantity)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *ComplexBidComponent) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetOfferedQuantity

`func (o *ComplexBidComponent) GetOfferedQuantity() Quantity`

GetOfferedQuantity returns the OfferedQuantity field if non-nil, zero value otherwise.

### GetOfferedQuantityOk

`func (o *ComplexBidComponent) GetOfferedQuantityOk() (*Quantity, bool)`

GetOfferedQuantityOk returns a tuple with the OfferedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferedQuantity

`func (o *ComplexBidComponent) SetOfferedQuantity(v Quantity)`

SetOfferedQuantity sets OfferedQuantity field to given value.


### GetPrices

`func (o *ComplexBidComponent) GetPrices() []Price`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ComplexBidComponent) GetPricesOk() (*[]Price, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ComplexBidComponent) SetPrices(v []Price)`

SetPrices sets Prices field to given value.


### GetTag

`func (o *ComplexBidComponent) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ComplexBidComponent) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ComplexBidComponent) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ComplexBidComponent) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


