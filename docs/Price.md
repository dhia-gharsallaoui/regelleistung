# Price

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Specify the amount. | 
**MeasureUnit** | **string** | The unit of measurement underlying the price (MW per unit (code MAW), MWh per unit (code MWH)) is used to differentiate between capacity fees (EUR/MW) and energy prices (EUR/MWh). The units of measurement are in accordance with UN/ECE Recommendation 20. | 
**Resolution** | Pointer to **string** | Capacity price resolution in ISO 8601 format. This field is read-only for control purposes. It is determined automatically depending on the modalities applicable to the tender. | [optional] [readonly] 
**Currency** | Pointer to **string** | The currency in which the monetary amount is expressed. The maximum length of this information is 3 alphanumeric characters respecting the standard ISO 4217. | [optional] [default to "EUR"]
**PaymentDirection** | Pointer to **string** | The direction of payment for activated control power relates to the energy price only (measureUnit &#x3D; MWH). | [optional] [default to "GRID_TO_PROVIDER"]
**Type** | Pointer to **string** | Type of price. When bidding, the type must have the value &#x60;OFFER&#x60;.  | [optional] 

## Methods

### NewPrice

`func NewPrice(amount float32, measureUnit string, ) *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Price) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Price) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Price) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetMeasureUnit

`func (o *Price) GetMeasureUnit() string`

GetMeasureUnit returns the MeasureUnit field if non-nil, zero value otherwise.

### GetMeasureUnitOk

`func (o *Price) GetMeasureUnitOk() (*string, bool)`

GetMeasureUnitOk returns a tuple with the MeasureUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUnit

`func (o *Price) SetMeasureUnit(v string)`

SetMeasureUnit sets MeasureUnit field to given value.


### GetResolution

`func (o *Price) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *Price) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *Price) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *Price) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetCurrency

`func (o *Price) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Price) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Price) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Price) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPaymentDirection

`func (o *Price) GetPaymentDirection() string`

GetPaymentDirection returns the PaymentDirection field if non-nil, zero value otherwise.

### GetPaymentDirectionOk

`func (o *Price) GetPaymentDirectionOk() (*string, bool)`

GetPaymentDirectionOk returns a tuple with the PaymentDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDirection

`func (o *Price) SetPaymentDirection(v string)`

SetPaymentDirection sets PaymentDirection field to given value.

### HasPaymentDirection

`func (o *Price) HasPaymentDirection() bool`

HasPaymentDirection returns a boolean if a field has been set.

### GetType

`func (o *Price) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Price) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Price) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Price) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


