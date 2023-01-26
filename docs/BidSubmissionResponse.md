# BidSubmissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Bid identification | [optional] 
**ComplexBidId** | Pointer to **string** | Bid identification | [optional] 
**GroupId** | Pointer to **string** | Bid identification | [optional] 
**BidType** | Pointer to [**BidType**](BidType.md) |  | [optional] [default to SIMPLE]
**Tender** | Pointer to **string** | The tender identification is structured as follows:   &#x60;[productType]-[deliveryDate]-[market]-[sequenceNumber]&#x60;    * &#x60;productType&#x60; - one of &#x60;fcr&#x60;, &#x60;afrr&#x60; or &#x60;mfrr&#x60;   * &#x60;deliveryDate&#x60; - ISO 8601 date   * &#x60;market&#x60; - one of &#x60;cm&#x60;(capacity market) or &#x60;em&#x60;(energy market)   * &#x60;sequenceNumber&#x60; - An ascending number for the identification of different tender runs or time slices.   For more details on the new tender id format, see the [reference guide.](/docs/guide#tender-id).  | [optional] [readonly] 
**ProductType** | Pointer to [**ProductType**](ProductType.md) |  | [optional] 
**DeliveryDate** | Pointer to **string** | Delivery day of offered control reserve / energy (ISO 8601 format YYYY-MM-DD). | [optional] 
**Market** | Pointer to [**ReserveMarket**](ReserveMarket.md) |  | [optional] 
**ProductName** | Pointer to [**ProductName**](ProductName.md) |  | [optional] 
**Direction** | Pointer to [**Direction**](Direction.md) |  | [optional] 
**TimeInterval** | Pointer to [**TimeInterval**](TimeInterval.md) |  | [optional] 
**ConnectingZone** | Pointer to [**ConnectingZone**](ConnectingZone.md) |  | [optional] 
**ResourceProvider** | Pointer to **string** | Representation of EIC. The coding scheme is the Energy Identification Coding Scheme (EIC), maintained by ENTSO-e. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ReceiptTimestamp** | Pointer to **time.Time** | Receipt time of the bid. | [optional] 
**RevisionTimestamp** | Pointer to **time.Time** | Revision timestamp of the bid. | [optional] 
**Activation** | Pointer to [**ActivationType**](ActivationType.md) |  | [optional] [default to DIRECT]
**MinQuantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**OfferedQuantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**AcceptedQuantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**Prices** | Pointer to [**[]Price**](Price.md) | Price information of the bid. | [optional] 
**TechnicalLinkage** | Pointer to [**TechnicalLinkage**](TechnicalLinkage.md) |  | [optional] 
**ConditionalLinkage** | Pointer to [**ConditionalLinkage**](ConditionalLinkage.md) |  | [optional] 
**BackupFor** | Pointer to [**BalancingServiceProvider**](BalancingServiceProvider.md) |  | [optional] 
**Tag** | Pointer to **string** | Custom tag for bids | [optional] 
**Components** | Pointer to [**[]Bid**](Bid.md) |  | [optional] 

## Methods

### NewBidSubmissionResponse

`func NewBidSubmissionResponse() *BidSubmissionResponse`

NewBidSubmissionResponse instantiates a new BidSubmissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBidSubmissionResponseWithDefaults

`func NewBidSubmissionResponseWithDefaults() *BidSubmissionResponse`

NewBidSubmissionResponseWithDefaults instantiates a new BidSubmissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BidSubmissionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BidSubmissionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BidSubmissionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BidSubmissionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComplexBidId

`func (o *BidSubmissionResponse) GetComplexBidId() string`

GetComplexBidId returns the ComplexBidId field if non-nil, zero value otherwise.

### GetComplexBidIdOk

`func (o *BidSubmissionResponse) GetComplexBidIdOk() (*string, bool)`

GetComplexBidIdOk returns a tuple with the ComplexBidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexBidId

`func (o *BidSubmissionResponse) SetComplexBidId(v string)`

SetComplexBidId sets ComplexBidId field to given value.

### HasComplexBidId

`func (o *BidSubmissionResponse) HasComplexBidId() bool`

HasComplexBidId returns a boolean if a field has been set.

### GetGroupId

`func (o *BidSubmissionResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BidSubmissionResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BidSubmissionResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BidSubmissionResponse) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetBidType

`func (o *BidSubmissionResponse) GetBidType() BidType`

GetBidType returns the BidType field if non-nil, zero value otherwise.

### GetBidTypeOk

`func (o *BidSubmissionResponse) GetBidTypeOk() (*BidType, bool)`

GetBidTypeOk returns a tuple with the BidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidType

`func (o *BidSubmissionResponse) SetBidType(v BidType)`

SetBidType sets BidType field to given value.

### HasBidType

`func (o *BidSubmissionResponse) HasBidType() bool`

HasBidType returns a boolean if a field has been set.

### GetTender

`func (o *BidSubmissionResponse) GetTender() string`

GetTender returns the Tender field if non-nil, zero value otherwise.

### GetTenderOk

`func (o *BidSubmissionResponse) GetTenderOk() (*string, bool)`

GetTenderOk returns a tuple with the Tender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTender

`func (o *BidSubmissionResponse) SetTender(v string)`

SetTender sets Tender field to given value.

### HasTender

`func (o *BidSubmissionResponse) HasTender() bool`

HasTender returns a boolean if a field has been set.

### GetProductType

`func (o *BidSubmissionResponse) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *BidSubmissionResponse) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *BidSubmissionResponse) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *BidSubmissionResponse) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *BidSubmissionResponse) GetDeliveryDate() string`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *BidSubmissionResponse) GetDeliveryDateOk() (*string, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *BidSubmissionResponse) SetDeliveryDate(v string)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *BidSubmissionResponse) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetMarket

`func (o *BidSubmissionResponse) GetMarket() ReserveMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *BidSubmissionResponse) GetMarketOk() (*ReserveMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *BidSubmissionResponse) SetMarket(v ReserveMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *BidSubmissionResponse) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetProductName

`func (o *BidSubmissionResponse) GetProductName() ProductName`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *BidSubmissionResponse) GetProductNameOk() (*ProductName, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *BidSubmissionResponse) SetProductName(v ProductName)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *BidSubmissionResponse) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetDirection

`func (o *BidSubmissionResponse) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BidSubmissionResponse) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BidSubmissionResponse) SetDirection(v Direction)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *BidSubmissionResponse) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetTimeInterval

`func (o *BidSubmissionResponse) GetTimeInterval() TimeInterval`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *BidSubmissionResponse) GetTimeIntervalOk() (*TimeInterval, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *BidSubmissionResponse) SetTimeInterval(v TimeInterval)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *BidSubmissionResponse) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.

### GetConnectingZone

`func (o *BidSubmissionResponse) GetConnectingZone() ConnectingZone`

GetConnectingZone returns the ConnectingZone field if non-nil, zero value otherwise.

### GetConnectingZoneOk

`func (o *BidSubmissionResponse) GetConnectingZoneOk() (*ConnectingZone, bool)`

GetConnectingZoneOk returns a tuple with the ConnectingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectingZone

`func (o *BidSubmissionResponse) SetConnectingZone(v ConnectingZone)`

SetConnectingZone sets ConnectingZone field to given value.

### HasConnectingZone

`func (o *BidSubmissionResponse) HasConnectingZone() bool`

HasConnectingZone returns a boolean if a field has been set.

### GetResourceProvider

`func (o *BidSubmissionResponse) GetResourceProvider() string`

GetResourceProvider returns the ResourceProvider field if non-nil, zero value otherwise.

### GetResourceProviderOk

`func (o *BidSubmissionResponse) GetResourceProviderOk() (*string, bool)`

GetResourceProviderOk returns a tuple with the ResourceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProvider

`func (o *BidSubmissionResponse) SetResourceProvider(v string)`

SetResourceProvider sets ResourceProvider field to given value.

### HasResourceProvider

`func (o *BidSubmissionResponse) HasResourceProvider() bool`

HasResourceProvider returns a boolean if a field has been set.

### GetStatus

`func (o *BidSubmissionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BidSubmissionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BidSubmissionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BidSubmissionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReceiptTimestamp

`func (o *BidSubmissionResponse) GetReceiptTimestamp() time.Time`

GetReceiptTimestamp returns the ReceiptTimestamp field if non-nil, zero value otherwise.

### GetReceiptTimestampOk

`func (o *BidSubmissionResponse) GetReceiptTimestampOk() (*time.Time, bool)`

GetReceiptTimestampOk returns a tuple with the ReceiptTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptTimestamp

`func (o *BidSubmissionResponse) SetReceiptTimestamp(v time.Time)`

SetReceiptTimestamp sets ReceiptTimestamp field to given value.

### HasReceiptTimestamp

`func (o *BidSubmissionResponse) HasReceiptTimestamp() bool`

HasReceiptTimestamp returns a boolean if a field has been set.

### GetRevisionTimestamp

`func (o *BidSubmissionResponse) GetRevisionTimestamp() time.Time`

GetRevisionTimestamp returns the RevisionTimestamp field if non-nil, zero value otherwise.

### GetRevisionTimestampOk

`func (o *BidSubmissionResponse) GetRevisionTimestampOk() (*time.Time, bool)`

GetRevisionTimestampOk returns a tuple with the RevisionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionTimestamp

`func (o *BidSubmissionResponse) SetRevisionTimestamp(v time.Time)`

SetRevisionTimestamp sets RevisionTimestamp field to given value.

### HasRevisionTimestamp

`func (o *BidSubmissionResponse) HasRevisionTimestamp() bool`

HasRevisionTimestamp returns a boolean if a field has been set.

### GetActivation

`func (o *BidSubmissionResponse) GetActivation() ActivationType`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *BidSubmissionResponse) GetActivationOk() (*ActivationType, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *BidSubmissionResponse) SetActivation(v ActivationType)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *BidSubmissionResponse) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetMinQuantity

`func (o *BidSubmissionResponse) GetMinQuantity() Quantity`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *BidSubmissionResponse) GetMinQuantityOk() (*Quantity, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *BidSubmissionResponse) SetMinQuantity(v Quantity)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *BidSubmissionResponse) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetOfferedQuantity

`func (o *BidSubmissionResponse) GetOfferedQuantity() Quantity`

GetOfferedQuantity returns the OfferedQuantity field if non-nil, zero value otherwise.

### GetOfferedQuantityOk

`func (o *BidSubmissionResponse) GetOfferedQuantityOk() (*Quantity, bool)`

GetOfferedQuantityOk returns a tuple with the OfferedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferedQuantity

`func (o *BidSubmissionResponse) SetOfferedQuantity(v Quantity)`

SetOfferedQuantity sets OfferedQuantity field to given value.

### HasOfferedQuantity

`func (o *BidSubmissionResponse) HasOfferedQuantity() bool`

HasOfferedQuantity returns a boolean if a field has been set.

### GetAcceptedQuantity

`func (o *BidSubmissionResponse) GetAcceptedQuantity() Quantity`

GetAcceptedQuantity returns the AcceptedQuantity field if non-nil, zero value otherwise.

### GetAcceptedQuantityOk

`func (o *BidSubmissionResponse) GetAcceptedQuantityOk() (*Quantity, bool)`

GetAcceptedQuantityOk returns a tuple with the AcceptedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedQuantity

`func (o *BidSubmissionResponse) SetAcceptedQuantity(v Quantity)`

SetAcceptedQuantity sets AcceptedQuantity field to given value.

### HasAcceptedQuantity

`func (o *BidSubmissionResponse) HasAcceptedQuantity() bool`

HasAcceptedQuantity returns a boolean if a field has been set.

### GetPrices

`func (o *BidSubmissionResponse) GetPrices() []Price`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *BidSubmissionResponse) GetPricesOk() (*[]Price, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *BidSubmissionResponse) SetPrices(v []Price)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *BidSubmissionResponse) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetTechnicalLinkage

`func (o *BidSubmissionResponse) GetTechnicalLinkage() TechnicalLinkage`

GetTechnicalLinkage returns the TechnicalLinkage field if non-nil, zero value otherwise.

### GetTechnicalLinkageOk

`func (o *BidSubmissionResponse) GetTechnicalLinkageOk() (*TechnicalLinkage, bool)`

GetTechnicalLinkageOk returns a tuple with the TechnicalLinkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalLinkage

`func (o *BidSubmissionResponse) SetTechnicalLinkage(v TechnicalLinkage)`

SetTechnicalLinkage sets TechnicalLinkage field to given value.

### HasTechnicalLinkage

`func (o *BidSubmissionResponse) HasTechnicalLinkage() bool`

HasTechnicalLinkage returns a boolean if a field has been set.

### GetConditionalLinkage

`func (o *BidSubmissionResponse) GetConditionalLinkage() ConditionalLinkage`

GetConditionalLinkage returns the ConditionalLinkage field if non-nil, zero value otherwise.

### GetConditionalLinkageOk

`func (o *BidSubmissionResponse) GetConditionalLinkageOk() (*ConditionalLinkage, bool)`

GetConditionalLinkageOk returns a tuple with the ConditionalLinkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalLinkage

`func (o *BidSubmissionResponse) SetConditionalLinkage(v ConditionalLinkage)`

SetConditionalLinkage sets ConditionalLinkage field to given value.

### HasConditionalLinkage

`func (o *BidSubmissionResponse) HasConditionalLinkage() bool`

HasConditionalLinkage returns a boolean if a field has been set.

### GetBackupFor

`func (o *BidSubmissionResponse) GetBackupFor() BalancingServiceProvider`

GetBackupFor returns the BackupFor field if non-nil, zero value otherwise.

### GetBackupForOk

`func (o *BidSubmissionResponse) GetBackupForOk() (*BalancingServiceProvider, bool)`

GetBackupForOk returns a tuple with the BackupFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFor

`func (o *BidSubmissionResponse) SetBackupFor(v BalancingServiceProvider)`

SetBackupFor sets BackupFor field to given value.

### HasBackupFor

`func (o *BidSubmissionResponse) HasBackupFor() bool`

HasBackupFor returns a boolean if a field has been set.

### GetTag

`func (o *BidSubmissionResponse) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BidSubmissionResponse) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BidSubmissionResponse) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *BidSubmissionResponse) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetComponents

`func (o *BidSubmissionResponse) GetComponents() []Bid`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BidSubmissionResponse) GetComponentsOk() (*[]Bid, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BidSubmissionResponse) SetComponents(v []Bid)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *BidSubmissionResponse) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


