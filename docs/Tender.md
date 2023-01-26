# Tender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The tender identification is structured as follows:   &#x60;[productType]-[deliveryDate]-[market]-[sequenceNumber]&#x60;    * &#x60;productType&#x60; - one of &#x60;fcr&#x60;, &#x60;afrr&#x60; or &#x60;mfrr&#x60;   * &#x60;deliveryDate&#x60; - ISO 8601 date   * &#x60;market&#x60; - one of &#x60;cm&#x60;(capacity market) or &#x60;em&#x60;(energy market)   * &#x60;sequenceNumber&#x60; - An ascending number for the identification of different tender runs or time slices.   For more details on the new tender id format, see the [reference guide.](/docs/guide#tender-id).  | [optional] [readonly] 
**ProductType** | Pointer to [**ProductType**](ProductType.md) |  | [optional] 
**DeliveryDate** | Pointer to **string** | Delivery day of offered control reserve / energy. (ISO 8601 format YYYY-MM-DD). | [optional] 
**Market** | Pointer to [**ReserveMarket**](ReserveMarket.md) |  | [optional] 
**TimeInterval** | Pointer to [**TimeInterval**](TimeInterval.md) |  | [optional] 
**GateOpenTime** | Pointer to **time.Time** | Gate open time of tender. | [optional] 
**GateClosureTime** | Pointer to **time.Time** | Gate closure time of tender. | [optional] 
**State** | Pointer to **string** | Actual state of the tender. For detailed information on the tender state model, see the [reference guide.](/docs/guide#tender-state-model) | [optional] 
**Result** | Pointer to **string** | Result of the tender. For detailed information on the specific tender results, see the [reference guide.](/docs/guide#tender-state-model) | [optional] 
**RunNumber** | Pointer to **NullableInt32** |  | [optional] 
**Products** | Pointer to [**[]ProductName**](ProductName.md) | Available products for tender. | [optional] 
**BusinessRulesKey** | Pointer to **string** | Key to identify the current business rules for bidding. For detailed information, see the [reference guide.](/docs/guide#business-rules-key) | [optional] 

## Methods

### NewTender

`func NewTender() *Tender`

NewTender instantiates a new Tender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenderWithDefaults

`func NewTenderWithDefaults() *Tender`

NewTenderWithDefaults instantiates a new Tender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tender) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tender) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tender) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Tender) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductType

`func (o *Tender) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *Tender) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *Tender) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *Tender) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *Tender) GetDeliveryDate() string`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *Tender) GetDeliveryDateOk() (*string, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *Tender) SetDeliveryDate(v string)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *Tender) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetMarket

`func (o *Tender) GetMarket() ReserveMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *Tender) GetMarketOk() (*ReserveMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *Tender) SetMarket(v ReserveMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *Tender) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetTimeInterval

`func (o *Tender) GetTimeInterval() TimeInterval`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *Tender) GetTimeIntervalOk() (*TimeInterval, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *Tender) SetTimeInterval(v TimeInterval)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *Tender) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.

### GetGateOpenTime

`func (o *Tender) GetGateOpenTime() time.Time`

GetGateOpenTime returns the GateOpenTime field if non-nil, zero value otherwise.

### GetGateOpenTimeOk

`func (o *Tender) GetGateOpenTimeOk() (*time.Time, bool)`

GetGateOpenTimeOk returns a tuple with the GateOpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateOpenTime

`func (o *Tender) SetGateOpenTime(v time.Time)`

SetGateOpenTime sets GateOpenTime field to given value.

### HasGateOpenTime

`func (o *Tender) HasGateOpenTime() bool`

HasGateOpenTime returns a boolean if a field has been set.

### GetGateClosureTime

`func (o *Tender) GetGateClosureTime() time.Time`

GetGateClosureTime returns the GateClosureTime field if non-nil, zero value otherwise.

### GetGateClosureTimeOk

`func (o *Tender) GetGateClosureTimeOk() (*time.Time, bool)`

GetGateClosureTimeOk returns a tuple with the GateClosureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateClosureTime

`func (o *Tender) SetGateClosureTime(v time.Time)`

SetGateClosureTime sets GateClosureTime field to given value.

### HasGateClosureTime

`func (o *Tender) HasGateClosureTime() bool`

HasGateClosureTime returns a boolean if a field has been set.

### GetState

`func (o *Tender) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Tender) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Tender) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Tender) HasState() bool`

HasState returns a boolean if a field has been set.

### GetResult

`func (o *Tender) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Tender) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Tender) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *Tender) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRunNumber

`func (o *Tender) GetRunNumber() int32`

GetRunNumber returns the RunNumber field if non-nil, zero value otherwise.

### GetRunNumberOk

`func (o *Tender) GetRunNumberOk() (*int32, bool)`

GetRunNumberOk returns a tuple with the RunNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunNumber

`func (o *Tender) SetRunNumber(v int32)`

SetRunNumber sets RunNumber field to given value.

### HasRunNumber

`func (o *Tender) HasRunNumber() bool`

HasRunNumber returns a boolean if a field has been set.

### SetRunNumberNil

`func (o *Tender) SetRunNumberNil(b bool)`

 SetRunNumberNil sets the value for RunNumber to be an explicit nil

### UnsetRunNumber
`func (o *Tender) UnsetRunNumber()`

UnsetRunNumber ensures that no value is present for RunNumber, not even an explicit nil
### GetProducts

`func (o *Tender) GetProducts() []ProductName`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *Tender) GetProductsOk() (*[]ProductName, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *Tender) SetProducts(v []ProductName)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *Tender) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetBusinessRulesKey

`func (o *Tender) GetBusinessRulesKey() string`

GetBusinessRulesKey returns the BusinessRulesKey field if non-nil, zero value otherwise.

### GetBusinessRulesKeyOk

`func (o *Tender) GetBusinessRulesKeyOk() (*string, bool)`

GetBusinessRulesKeyOk returns a tuple with the BusinessRulesKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessRulesKey

`func (o *Tender) SetBusinessRulesKey(v string)`

SetBusinessRulesKey sets BusinessRulesKey field to given value.

### HasBusinessRulesKey

`func (o *Tender) HasBusinessRulesKey() bool`

HasBusinessRulesKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


