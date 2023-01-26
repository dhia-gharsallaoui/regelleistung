# BalancingServiceProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortName** | Pointer to **string** | A short, human-readable name of the provider. | [optional] 
**Name** | Pointer to **string** | Full, human-readable name of the provider. | [optional] 
**Eic** | **string** | Representation of EIC. The coding scheme is the Energy Identification Coding Scheme (EIC), maintained by ENTSO-e. | 

## Methods

### NewBalancingServiceProvider

`func NewBalancingServiceProvider(eic string, ) *BalancingServiceProvider`

NewBalancingServiceProvider instantiates a new BalancingServiceProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalancingServiceProviderWithDefaults

`func NewBalancingServiceProviderWithDefaults() *BalancingServiceProvider`

NewBalancingServiceProviderWithDefaults instantiates a new BalancingServiceProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortName

`func (o *BalancingServiceProvider) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *BalancingServiceProvider) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *BalancingServiceProvider) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *BalancingServiceProvider) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetName

`func (o *BalancingServiceProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BalancingServiceProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BalancingServiceProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BalancingServiceProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEic

`func (o *BalancingServiceProvider) GetEic() string`

GetEic returns the Eic field if non-nil, zero value otherwise.

### GetEicOk

`func (o *BalancingServiceProvider) GetEicOk() (*string, bool)`

GetEicOk returns a tuple with the Eic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEic

`func (o *BalancingServiceProvider) SetEic(v string)`

SetEic sets Eic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


