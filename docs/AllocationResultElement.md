# AllocationResultElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductType** | Pointer to [**ProductType**](ProductType.md) |  | [optional] 
**DeliveryDate** | Pointer to **string** | Delivery day of offered control reserve / energy. (ISO 8601 format YYYY-MM-DD). | [optional] 
**ProductName** | Pointer to [**ProductName**](ProductName.md) |  | [optional] 
**Direction** | Pointer to [**Direction**](Direction.md) |  | [optional] 
**TimeInterval** | Pointer to [**TimeInterval**](TimeInterval.md) |  | [optional] 
**ConnectingZone** | Pointer to [**ConnectingZone**](ConnectingZone.md) |  | [optional] 
**Market** | Pointer to [**ReserveMarket**](ReserveMarket.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AllocationRanking** | Pointer to **int32** | Specifies the allocation ranking of bids as a natural number in ascending order starting at 1. | [optional] 
**ActivationRanking** | Pointer to **NullableInt32** | Specifies the activation ranking of accepted aFRR/mFRR bids as a natural number in ascending order starting at 1. | [optional] 
**Indivisible** | Pointer to **bool** |  | [optional] [default to false]
**OfferedQuantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**AcceptedQuantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**Prices** | Pointer to [**[]Price**](Price.md) | Price informations. | [optional] 

## Methods

### NewAllocationResultElement

`func NewAllocationResultElement() *AllocationResultElement`

NewAllocationResultElement instantiates a new AllocationResultElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationResultElementWithDefaults

`func NewAllocationResultElementWithDefaults() *AllocationResultElement`

NewAllocationResultElementWithDefaults instantiates a new AllocationResultElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductType

`func (o *AllocationResultElement) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AllocationResultElement) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AllocationResultElement) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AllocationResultElement) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *AllocationResultElement) GetDeliveryDate() string`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *AllocationResultElement) GetDeliveryDateOk() (*string, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *AllocationResultElement) SetDeliveryDate(v string)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *AllocationResultElement) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetProductName

`func (o *AllocationResultElement) GetProductName() ProductName`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *AllocationResultElement) GetProductNameOk() (*ProductName, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *AllocationResultElement) SetProductName(v ProductName)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *AllocationResultElement) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetDirection

`func (o *AllocationResultElement) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AllocationResultElement) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AllocationResultElement) SetDirection(v Direction)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *AllocationResultElement) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetTimeInterval

`func (o *AllocationResultElement) GetTimeInterval() TimeInterval`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *AllocationResultElement) GetTimeIntervalOk() (*TimeInterval, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *AllocationResultElement) SetTimeInterval(v TimeInterval)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *AllocationResultElement) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.

### GetConnectingZone

`func (o *AllocationResultElement) GetConnectingZone() ConnectingZone`

GetConnectingZone returns the ConnectingZone field if non-nil, zero value otherwise.

### GetConnectingZoneOk

`func (o *AllocationResultElement) GetConnectingZoneOk() (*ConnectingZone, bool)`

GetConnectingZoneOk returns a tuple with the ConnectingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectingZone

`func (o *AllocationResultElement) SetConnectingZone(v ConnectingZone)`

SetConnectingZone sets ConnectingZone field to given value.

### HasConnectingZone

`func (o *AllocationResultElement) HasConnectingZone() bool`

HasConnectingZone returns a boolean if a field has been set.

### GetMarket

`func (o *AllocationResultElement) GetMarket() ReserveMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *AllocationResultElement) GetMarketOk() (*ReserveMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *AllocationResultElement) SetMarket(v ReserveMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *AllocationResultElement) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetStatus

`func (o *AllocationResultElement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AllocationResultElement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AllocationResultElement) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AllocationResultElement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAllocationRanking

`func (o *AllocationResultElement) GetAllocationRanking() int32`

GetAllocationRanking returns the AllocationRanking field if non-nil, zero value otherwise.

### GetAllocationRankingOk

`func (o *AllocationResultElement) GetAllocationRankingOk() (*int32, bool)`

GetAllocationRankingOk returns a tuple with the AllocationRanking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationRanking

`func (o *AllocationResultElement) SetAllocationRanking(v int32)`

SetAllocationRanking sets AllocationRanking field to given value.

### HasAllocationRanking

`func (o *AllocationResultElement) HasAllocationRanking() bool`

HasAllocationRanking returns a boolean if a field has been set.

### GetActivationRanking

`func (o *AllocationResultElement) GetActivationRanking() int32`

GetActivationRanking returns the ActivationRanking field if non-nil, zero value otherwise.

### GetActivationRankingOk

`func (o *AllocationResultElement) GetActivationRankingOk() (*int32, bool)`

GetActivationRankingOk returns a tuple with the ActivationRanking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationRanking

`func (o *AllocationResultElement) SetActivationRanking(v int32)`

SetActivationRanking sets ActivationRanking field to given value.

### HasActivationRanking

`func (o *AllocationResultElement) HasActivationRanking() bool`

HasActivationRanking returns a boolean if a field has been set.

### SetActivationRankingNil

`func (o *AllocationResultElement) SetActivationRankingNil(b bool)`

 SetActivationRankingNil sets the value for ActivationRanking to be an explicit nil

### UnsetActivationRanking
`func (o *AllocationResultElement) UnsetActivationRanking()`

UnsetActivationRanking ensures that no value is present for ActivationRanking, not even an explicit nil
### GetIndivisible

`func (o *AllocationResultElement) GetIndivisible() bool`

GetIndivisible returns the Indivisible field if non-nil, zero value otherwise.

### GetIndivisibleOk

`func (o *AllocationResultElement) GetIndivisibleOk() (*bool, bool)`

GetIndivisibleOk returns a tuple with the Indivisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndivisible

`func (o *AllocationResultElement) SetIndivisible(v bool)`

SetIndivisible sets Indivisible field to given value.

### HasIndivisible

`func (o *AllocationResultElement) HasIndivisible() bool`

HasIndivisible returns a boolean if a field has been set.

### GetOfferedQuantity

`func (o *AllocationResultElement) GetOfferedQuantity() Quantity`

GetOfferedQuantity returns the OfferedQuantity field if non-nil, zero value otherwise.

### GetOfferedQuantityOk

`func (o *AllocationResultElement) GetOfferedQuantityOk() (*Quantity, bool)`

GetOfferedQuantityOk returns a tuple with the OfferedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferedQuantity

`func (o *AllocationResultElement) SetOfferedQuantity(v Quantity)`

SetOfferedQuantity sets OfferedQuantity field to given value.

### HasOfferedQuantity

`func (o *AllocationResultElement) HasOfferedQuantity() bool`

HasOfferedQuantity returns a boolean if a field has been set.

### GetAcceptedQuantity

`func (o *AllocationResultElement) GetAcceptedQuantity() Quantity`

GetAcceptedQuantity returns the AcceptedQuantity field if non-nil, zero value otherwise.

### GetAcceptedQuantityOk

`func (o *AllocationResultElement) GetAcceptedQuantityOk() (*Quantity, bool)`

GetAcceptedQuantityOk returns a tuple with the AcceptedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedQuantity

`func (o *AllocationResultElement) SetAcceptedQuantity(v Quantity)`

SetAcceptedQuantity sets AcceptedQuantity field to given value.

### HasAcceptedQuantity

`func (o *AllocationResultElement) HasAcceptedQuantity() bool`

HasAcceptedQuantity returns a boolean if a field has been set.

### GetPrices

`func (o *AllocationResultElement) GetPrices() []Price`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *AllocationResultElement) GetPricesOk() (*[]Price, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *AllocationResultElement) SetPrices(v []Price)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *AllocationResultElement) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


