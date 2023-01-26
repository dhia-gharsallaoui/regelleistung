# BatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ops** | [**[]BatchOperation**](BatchOperation.md) | List of operations to be processed. | 
**Parallel** | Pointer to **bool** | If true, then the operations can be executed in arbitrary order and in parallel. If false, then serial execution in request order will be enforced. This property has no impact on the results order in the batch response object. | [optional] [default to false]

## Methods

### NewBatchRequest

`func NewBatchRequest(ops []BatchOperation, ) *BatchRequest`

NewBatchRequest instantiates a new BatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRequestWithDefaults

`func NewBatchRequestWithDefaults() *BatchRequest`

NewBatchRequestWithDefaults instantiates a new BatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOps

`func (o *BatchRequest) GetOps() []BatchOperation`

GetOps returns the Ops field if non-nil, zero value otherwise.

### GetOpsOk

`func (o *BatchRequest) GetOpsOk() (*[]BatchOperation, bool)`

GetOpsOk returns a tuple with the Ops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOps

`func (o *BatchRequest) SetOps(v []BatchOperation)`

SetOps sets Ops field to given value.


### GetParallel

`func (o *BatchRequest) GetParallel() bool`

GetParallel returns the Parallel field if non-nil, zero value otherwise.

### GetParallelOk

`func (o *BatchRequest) GetParallelOk() (*bool, bool)`

GetParallelOk returns a tuple with the Parallel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel

`func (o *BatchRequest) SetParallel(v bool)`

SetParallel sets Parallel field to given value.

### HasParallel

`func (o *BatchRequest) HasParallel() bool`

HasParallel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


