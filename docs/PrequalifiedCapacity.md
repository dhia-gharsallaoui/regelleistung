# PrequalifiedCapacity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductType** | Pointer to [**ProductType**](ProductType.md) |  | [optional] 
**ConnectingZone** | Pointer to [**ConnectingZone**](ConnectingZone.md) |  | [optional] 
**MaxNegativeCapacity** | Pointer to **float32** | Volume of maximum negative reserve in specific connecting zone. | [optional] 
**MaxPositiveCapacity** | Pointer to **float32** | Volume of maximum positive reserve in specific connecting zone. | [optional] 

## Methods

### NewPrequalifiedCapacity

`func NewPrequalifiedCapacity() *PrequalifiedCapacity`

NewPrequalifiedCapacity instantiates a new PrequalifiedCapacity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrequalifiedCapacityWithDefaults

`func NewPrequalifiedCapacityWithDefaults() *PrequalifiedCapacity`

NewPrequalifiedCapacityWithDefaults instantiates a new PrequalifiedCapacity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductType

`func (o *PrequalifiedCapacity) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PrequalifiedCapacity) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PrequalifiedCapacity) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *PrequalifiedCapacity) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetConnectingZone

`func (o *PrequalifiedCapacity) GetConnectingZone() ConnectingZone`

GetConnectingZone returns the ConnectingZone field if non-nil, zero value otherwise.

### GetConnectingZoneOk

`func (o *PrequalifiedCapacity) GetConnectingZoneOk() (*ConnectingZone, bool)`

GetConnectingZoneOk returns a tuple with the ConnectingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectingZone

`func (o *PrequalifiedCapacity) SetConnectingZone(v ConnectingZone)`

SetConnectingZone sets ConnectingZone field to given value.

### HasConnectingZone

`func (o *PrequalifiedCapacity) HasConnectingZone() bool`

HasConnectingZone returns a boolean if a field has been set.

### GetMaxNegativeCapacity

`func (o *PrequalifiedCapacity) GetMaxNegativeCapacity() float32`

GetMaxNegativeCapacity returns the MaxNegativeCapacity field if non-nil, zero value otherwise.

### GetMaxNegativeCapacityOk

`func (o *PrequalifiedCapacity) GetMaxNegativeCapacityOk() (*float32, bool)`

GetMaxNegativeCapacityOk returns a tuple with the MaxNegativeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNegativeCapacity

`func (o *PrequalifiedCapacity) SetMaxNegativeCapacity(v float32)`

SetMaxNegativeCapacity sets MaxNegativeCapacity field to given value.

### HasMaxNegativeCapacity

`func (o *PrequalifiedCapacity) HasMaxNegativeCapacity() bool`

HasMaxNegativeCapacity returns a boolean if a field has been set.

### GetMaxPositiveCapacity

`func (o *PrequalifiedCapacity) GetMaxPositiveCapacity() float32`

GetMaxPositiveCapacity returns the MaxPositiveCapacity field if non-nil, zero value otherwise.

### GetMaxPositiveCapacityOk

`func (o *PrequalifiedCapacity) GetMaxPositiveCapacityOk() (*float32, bool)`

GetMaxPositiveCapacityOk returns a tuple with the MaxPositiveCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPositiveCapacity

`func (o *PrequalifiedCapacity) SetMaxPositiveCapacity(v float32)`

SetMaxPositiveCapacity sets MaxPositiveCapacity field to given value.

### HasMaxPositiveCapacity

`func (o *PrequalifiedCapacity) HasMaxPositiveCapacity() bool`

HasMaxPositiveCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


