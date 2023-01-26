# BatchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | A permitted relative BSP API v2 resource path. | 
**Method** | [**BatchSupportedHttpMethod**](BatchSupportedHttpMethod.md) |  | 
**Body** | Pointer to [**BatchOperationBody**](BatchOperationBody.md) |  | [optional] 

## Methods

### NewBatchOperation

`func NewBatchOperation(path string, method BatchSupportedHttpMethod, ) *BatchOperation`

NewBatchOperation instantiates a new BatchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchOperationWithDefaults

`func NewBatchOperationWithDefaults() *BatchOperation`

NewBatchOperationWithDefaults instantiates a new BatchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *BatchOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BatchOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BatchOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetMethod

`func (o *BatchOperation) GetMethod() BatchSupportedHttpMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *BatchOperation) GetMethodOk() (*BatchSupportedHttpMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *BatchOperation) SetMethod(v BatchSupportedHttpMethod)`

SetMethod sets Method field to given value.


### GetBody

`func (o *BatchOperation) GetBody() BatchOperationBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BatchOperation) GetBodyOk() (*BatchOperationBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BatchOperation) SetBody(v BatchOperationBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *BatchOperation) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


