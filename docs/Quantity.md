# Quantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | **float32** | Specify the quantity. | 
**MeasureUnit** | **string** | The unit of measurement underlying the volume (MW per unit (code MAW), MWh per unit (code MWH)). The units of measurement are in accordance with UN/ECE Recommendation 20. | 

## Methods

### NewQuantity

`func NewQuantity(quantity float32, measureUnit string, ) *Quantity`

NewQuantity instantiates a new Quantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuantityWithDefaults

`func NewQuantityWithDefaults() *Quantity`

NewQuantityWithDefaults instantiates a new Quantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *Quantity) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Quantity) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Quantity) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetMeasureUnit

`func (o *Quantity) GetMeasureUnit() string`

GetMeasureUnit returns the MeasureUnit field if non-nil, zero value otherwise.

### GetMeasureUnitOk

`func (o *Quantity) GetMeasureUnitOk() (*string, bool)`

GetMeasureUnitOk returns a tuple with the MeasureUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUnit

`func (o *Quantity) SetMeasureUnit(v string)`

SetMeasureUnit sets MeasureUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


