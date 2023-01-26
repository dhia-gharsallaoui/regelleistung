# BackupContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a resource. | [readonly] 
**ProductType** | Pointer to [**ProductType**](ProductType.md) |  | [optional] 
**BackupProvider** | Pointer to [**BalancingServiceProvider**](BalancingServiceProvider.md) |  | [optional] 
**BackedUpProvider** | Pointer to [**BalancingServiceProvider**](BalancingServiceProvider.md) |  | [optional] 
**ValidityPeriod** | Pointer to [**ContractValidityPeriod**](ContractValidityPeriod.md) |  | [optional] 

## Methods

### NewBackupContract

`func NewBackupContract(id string, ) *BackupContract`

NewBackupContract instantiates a new BackupContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupContractWithDefaults

`func NewBackupContractWithDefaults() *BackupContract`

NewBackupContractWithDefaults instantiates a new BackupContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupContract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupContract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupContract) SetId(v string)`

SetId sets Id field to given value.


### GetProductType

`func (o *BackupContract) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *BackupContract) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *BackupContract) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *BackupContract) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetBackupProvider

`func (o *BackupContract) GetBackupProvider() BalancingServiceProvider`

GetBackupProvider returns the BackupProvider field if non-nil, zero value otherwise.

### GetBackupProviderOk

`func (o *BackupContract) GetBackupProviderOk() (*BalancingServiceProvider, bool)`

GetBackupProviderOk returns a tuple with the BackupProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProvider

`func (o *BackupContract) SetBackupProvider(v BalancingServiceProvider)`

SetBackupProvider sets BackupProvider field to given value.

### HasBackupProvider

`func (o *BackupContract) HasBackupProvider() bool`

HasBackupProvider returns a boolean if a field has been set.

### GetBackedUpProvider

`func (o *BackupContract) GetBackedUpProvider() BalancingServiceProvider`

GetBackedUpProvider returns the BackedUpProvider field if non-nil, zero value otherwise.

### GetBackedUpProviderOk

`func (o *BackupContract) GetBackedUpProviderOk() (*BalancingServiceProvider, bool)`

GetBackedUpProviderOk returns a tuple with the BackedUpProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackedUpProvider

`func (o *BackupContract) SetBackedUpProvider(v BalancingServiceProvider)`

SetBackedUpProvider sets BackedUpProvider field to given value.

### HasBackedUpProvider

`func (o *BackupContract) HasBackedUpProvider() bool`

HasBackedUpProvider returns a boolean if a field has been set.

### GetValidityPeriod

`func (o *BackupContract) GetValidityPeriod() ContractValidityPeriod`

GetValidityPeriod returns the ValidityPeriod field if non-nil, zero value otherwise.

### GetValidityPeriodOk

`func (o *BackupContract) GetValidityPeriodOk() (*ContractValidityPeriod, bool)`

GetValidityPeriodOk returns a tuple with the ValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriod

`func (o *BackupContract) SetValidityPeriod(v ContractValidityPeriod)`

SetValidityPeriod sets ValidityPeriod field to given value.

### HasValidityPeriod

`func (o *BackupContract) HasValidityPeriod() bool`

HasValidityPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


