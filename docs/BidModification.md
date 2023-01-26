# BidModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectingZone** | Pointer to [**ConnectingZone**](ConnectingZone.md) |  | [optional] 
**MinQuantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**OfferedQuantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**Prices** | Pointer to [**[]Price**](Price.md) | Price information of the bid. | [optional] 
**Activation** | Pointer to [**ActivationType**](ActivationType.md) |  | [optional] [default to DIRECT]
**TechnicalLinkage** | Pointer to [**TechnicalLinkage**](TechnicalLinkage.md) |  | [optional] 
**ConditionalLinkage** | Pointer to [**ConditionalLinkage**](ConditionalLinkage.md) |  | [optional] 
**BackupFor** | Pointer to [**BalancingServiceProvider**](BalancingServiceProvider.md) |  | [optional] 

## Methods

### NewBidModification

`func NewBidModification() *BidModification`

NewBidModification instantiates a new BidModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBidModificationWithDefaults

`func NewBidModificationWithDefaults() *BidModification`

NewBidModificationWithDefaults instantiates a new BidModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectingZone

`func (o *BidModification) GetConnectingZone() ConnectingZone`

GetConnectingZone returns the ConnectingZone field if non-nil, zero value otherwise.

### GetConnectingZoneOk

`func (o *BidModification) GetConnectingZoneOk() (*ConnectingZone, bool)`

GetConnectingZoneOk returns a tuple with the ConnectingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectingZone

`func (o *BidModification) SetConnectingZone(v ConnectingZone)`

SetConnectingZone sets ConnectingZone field to given value.

### HasConnectingZone

`func (o *BidModification) HasConnectingZone() bool`

HasConnectingZone returns a boolean if a field has been set.

### GetMinQuantity

`func (o *BidModification) GetMinQuantity() Quantity`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *BidModification) GetMinQuantityOk() (*Quantity, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *BidModification) SetMinQuantity(v Quantity)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *BidModification) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetOfferedQuantity

`func (o *BidModification) GetOfferedQuantity() Quantity`

GetOfferedQuantity returns the OfferedQuantity field if non-nil, zero value otherwise.

### GetOfferedQuantityOk

`func (o *BidModification) GetOfferedQuantityOk() (*Quantity, bool)`

GetOfferedQuantityOk returns a tuple with the OfferedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferedQuantity

`func (o *BidModification) SetOfferedQuantity(v Quantity)`

SetOfferedQuantity sets OfferedQuantity field to given value.

### HasOfferedQuantity

`func (o *BidModification) HasOfferedQuantity() bool`

HasOfferedQuantity returns a boolean if a field has been set.

### GetPrices

`func (o *BidModification) GetPrices() []Price`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *BidModification) GetPricesOk() (*[]Price, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *BidModification) SetPrices(v []Price)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *BidModification) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetActivation

`func (o *BidModification) GetActivation() ActivationType`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *BidModification) GetActivationOk() (*ActivationType, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *BidModification) SetActivation(v ActivationType)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *BidModification) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetTechnicalLinkage

`func (o *BidModification) GetTechnicalLinkage() TechnicalLinkage`

GetTechnicalLinkage returns the TechnicalLinkage field if non-nil, zero value otherwise.

### GetTechnicalLinkageOk

`func (o *BidModification) GetTechnicalLinkageOk() (*TechnicalLinkage, bool)`

GetTechnicalLinkageOk returns a tuple with the TechnicalLinkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalLinkage

`func (o *BidModification) SetTechnicalLinkage(v TechnicalLinkage)`

SetTechnicalLinkage sets TechnicalLinkage field to given value.

### HasTechnicalLinkage

`func (o *BidModification) HasTechnicalLinkage() bool`

HasTechnicalLinkage returns a boolean if a field has been set.

### GetConditionalLinkage

`func (o *BidModification) GetConditionalLinkage() ConditionalLinkage`

GetConditionalLinkage returns the ConditionalLinkage field if non-nil, zero value otherwise.

### GetConditionalLinkageOk

`func (o *BidModification) GetConditionalLinkageOk() (*ConditionalLinkage, bool)`

GetConditionalLinkageOk returns a tuple with the ConditionalLinkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalLinkage

`func (o *BidModification) SetConditionalLinkage(v ConditionalLinkage)`

SetConditionalLinkage sets ConditionalLinkage field to given value.

### HasConditionalLinkage

`func (o *BidModification) HasConditionalLinkage() bool`

HasConditionalLinkage returns a boolean if a field has been set.

### GetBackupFor

`func (o *BidModification) GetBackupFor() BalancingServiceProvider`

GetBackupFor returns the BackupFor field if non-nil, zero value otherwise.

### GetBackupForOk

`func (o *BidModification) GetBackupForOk() (*BalancingServiceProvider, bool)`

GetBackupForOk returns a tuple with the BackupFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFor

`func (o *BidModification) SetBackupFor(v BalancingServiceProvider)`

SetBackupFor sets BackupFor field to given value.

### HasBackupFor

`func (o *BidModification) HasBackupFor() bool`

HasBackupFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


