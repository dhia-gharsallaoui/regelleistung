# ResultDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of a resource. | [optional] [readonly] 
**FileName** | Pointer to **string** | Document name | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation time of the document. | [optional] 
**ProductType** | Pointer to [**ProductType**](ProductType.md) |  | [optional] 
**DeliveryDate** | Pointer to **string** | Delivery day | [optional] 
**Market** | Pointer to [**ReserveMarket**](ReserveMarket.md) |  | [optional] 
**Tender** | Pointer to **string** | The tender identification is structured as follows:   &#x60;[productType]-[deliveryDate]-[market]-[sequenceNumber]&#x60;    * &#x60;productType&#x60; - one of &#x60;fcr&#x60;, &#x60;afrr&#x60; or &#x60;mfrr&#x60;   * &#x60;deliveryDate&#x60; - ISO 8601 date   * &#x60;market&#x60; - one of &#x60;cm&#x60;(capacity market) or &#x60;em&#x60;(energy market)   * &#x60;sequenceNumber&#x60; - An ascending number for the identification of different tender runs or time slices.   For more details on the new tender id format, see the [reference guide.](/docs/guide#tender-id).  | [optional] [readonly] 
**Tags** | Pointer to **[]string** | Tags are used to group the documents. A list of current tags can be found in the API documentation.  | [optional] 

## Methods

### NewResultDocument

`func NewResultDocument() *ResultDocument`

NewResultDocument instantiates a new ResultDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultDocumentWithDefaults

`func NewResultDocumentWithDefaults() *ResultDocument`

NewResultDocumentWithDefaults instantiates a new ResultDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResultDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResultDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResultDocument) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResultDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFileName

`func (o *ResultDocument) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ResultDocument) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ResultDocument) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ResultDocument) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResultDocument) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResultDocument) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResultDocument) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResultDocument) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetProductType

`func (o *ResultDocument) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ResultDocument) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ResultDocument) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *ResultDocument) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *ResultDocument) GetDeliveryDate() string`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *ResultDocument) GetDeliveryDateOk() (*string, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *ResultDocument) SetDeliveryDate(v string)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *ResultDocument) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetMarket

`func (o *ResultDocument) GetMarket() ReserveMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *ResultDocument) GetMarketOk() (*ReserveMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *ResultDocument) SetMarket(v ReserveMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *ResultDocument) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetTender

`func (o *ResultDocument) GetTender() string`

GetTender returns the Tender field if non-nil, zero value otherwise.

### GetTenderOk

`func (o *ResultDocument) GetTenderOk() (*string, bool)`

GetTenderOk returns a tuple with the Tender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTender

`func (o *ResultDocument) SetTender(v string)`

SetTender sets Tender field to given value.

### HasTender

`func (o *ResultDocument) HasTender() bool`

HasTender returns a boolean if a field has been set.

### GetTags

`func (o *ResultDocument) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResultDocument) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResultDocument) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ResultDocument) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


