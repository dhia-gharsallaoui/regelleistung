# NotificationReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductType** | Pointer to [**ProductType**](ProductType.md) |  | [optional] 
**DeliveryDate** | Pointer to **string** | Delivery day of offered control reserve / energy. (ISO 8601 format YYYY-MM-DD). | [optional] 
**Market** | Pointer to [**ReserveMarket**](ReserveMarket.md) |  | [optional] 
**Tender** | Pointer to **string** | The tender identification is structured as follows:   &#x60;[productType]-[deliveryDate]-[market]-[sequenceNumber]&#x60;    * &#x60;productType&#x60; - one of &#x60;fcr&#x60;, &#x60;afrr&#x60; or &#x60;mfrr&#x60;   * &#x60;deliveryDate&#x60; - ISO 8601 date   * &#x60;market&#x60; - one of &#x60;cm&#x60;(capacity market) or &#x60;em&#x60;(energy market)   * &#x60;sequenceNumber&#x60; - An ascending number for the identification of different tender runs or time slices.   For more details on the new tender id format, see the [reference guide.](/docs/guide#tender-id).  | [optional] [readonly] 
**Products** | Pointer to [**[]ProductName**](ProductName.md) |  | [optional] 

## Methods

### NewNotificationReference

`func NewNotificationReference() *NotificationReference`

NewNotificationReference instantiates a new NotificationReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationReferenceWithDefaults

`func NewNotificationReferenceWithDefaults() *NotificationReference`

NewNotificationReferenceWithDefaults instantiates a new NotificationReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductType

`func (o *NotificationReference) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *NotificationReference) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *NotificationReference) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *NotificationReference) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *NotificationReference) GetDeliveryDate() string`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *NotificationReference) GetDeliveryDateOk() (*string, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *NotificationReference) SetDeliveryDate(v string)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *NotificationReference) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetMarket

`func (o *NotificationReference) GetMarket() ReserveMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *NotificationReference) GetMarketOk() (*ReserveMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *NotificationReference) SetMarket(v ReserveMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *NotificationReference) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetTender

`func (o *NotificationReference) GetTender() string`

GetTender returns the Tender field if non-nil, zero value otherwise.

### GetTenderOk

`func (o *NotificationReference) GetTenderOk() (*string, bool)`

GetTenderOk returns a tuple with the Tender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTender

`func (o *NotificationReference) SetTender(v string)`

SetTender sets Tender field to given value.

### HasTender

`func (o *NotificationReference) HasTender() bool`

HasTender returns a boolean if a field has been set.

### GetProducts

`func (o *NotificationReference) GetProducts() []ProductName`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *NotificationReference) GetProductsOk() (*[]ProductName, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *NotificationReference) SetProducts(v []ProductName)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *NotificationReference) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


