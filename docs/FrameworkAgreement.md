# FrameworkAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of a resource. | [optional] [readonly] 
**ProductType** | Pointer to [**ProductType**](ProductType.md) |  | [optional] 
**ConnectingZone** | Pointer to [**ConnectingZone**](ConnectingZone.md) |  | [optional] 
**ResourceProvider** | Pointer to **string** | Representation of EIC. The coding scheme is the Energy Identification Coding Scheme (EIC), maintained by ENTSO-e. | [optional] 
**ContractNumber** | Pointer to **string** |  | [optional] 
**ValidityPeriod** | Pointer to [**ContractValidityPeriod**](ContractValidityPeriod.md) |  | [optional] 

## Methods

### NewFrameworkAgreement

`func NewFrameworkAgreement() *FrameworkAgreement`

NewFrameworkAgreement instantiates a new FrameworkAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameworkAgreementWithDefaults

`func NewFrameworkAgreementWithDefaults() *FrameworkAgreement`

NewFrameworkAgreementWithDefaults instantiates a new FrameworkAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FrameworkAgreement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FrameworkAgreement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FrameworkAgreement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FrameworkAgreement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductType

`func (o *FrameworkAgreement) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *FrameworkAgreement) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *FrameworkAgreement) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *FrameworkAgreement) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetConnectingZone

`func (o *FrameworkAgreement) GetConnectingZone() ConnectingZone`

GetConnectingZone returns the ConnectingZone field if non-nil, zero value otherwise.

### GetConnectingZoneOk

`func (o *FrameworkAgreement) GetConnectingZoneOk() (*ConnectingZone, bool)`

GetConnectingZoneOk returns a tuple with the ConnectingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectingZone

`func (o *FrameworkAgreement) SetConnectingZone(v ConnectingZone)`

SetConnectingZone sets ConnectingZone field to given value.

### HasConnectingZone

`func (o *FrameworkAgreement) HasConnectingZone() bool`

HasConnectingZone returns a boolean if a field has been set.

### GetResourceProvider

`func (o *FrameworkAgreement) GetResourceProvider() string`

GetResourceProvider returns the ResourceProvider field if non-nil, zero value otherwise.

### GetResourceProviderOk

`func (o *FrameworkAgreement) GetResourceProviderOk() (*string, bool)`

GetResourceProviderOk returns a tuple with the ResourceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProvider

`func (o *FrameworkAgreement) SetResourceProvider(v string)`

SetResourceProvider sets ResourceProvider field to given value.

### HasResourceProvider

`func (o *FrameworkAgreement) HasResourceProvider() bool`

HasResourceProvider returns a boolean if a field has been set.

### GetContractNumber

`func (o *FrameworkAgreement) GetContractNumber() string`

GetContractNumber returns the ContractNumber field if non-nil, zero value otherwise.

### GetContractNumberOk

`func (o *FrameworkAgreement) GetContractNumberOk() (*string, bool)`

GetContractNumberOk returns a tuple with the ContractNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractNumber

`func (o *FrameworkAgreement) SetContractNumber(v string)`

SetContractNumber sets ContractNumber field to given value.

### HasContractNumber

`func (o *FrameworkAgreement) HasContractNumber() bool`

HasContractNumber returns a boolean if a field has been set.

### GetValidityPeriod

`func (o *FrameworkAgreement) GetValidityPeriod() ContractValidityPeriod`

GetValidityPeriod returns the ValidityPeriod field if non-nil, zero value otherwise.

### GetValidityPeriodOk

`func (o *FrameworkAgreement) GetValidityPeriodOk() (*ContractValidityPeriod, bool)`

GetValidityPeriodOk returns a tuple with the ValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriod

`func (o *FrameworkAgreement) SetValidityPeriod(v ContractValidityPeriod)`

SetValidityPeriod sets ValidityPeriod field to given value.

### HasValidityPeriod

`func (o *FrameworkAgreement) HasValidityPeriod() bool`

HasValidityPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


