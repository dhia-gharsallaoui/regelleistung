# BatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Responses** | Pointer to [**[]BatchOperationResponse**](BatchOperationResponse.md) | Contains the list of individual responses in the order of the batch requests. | [optional] 
**SuccessfulOps** | Pointer to **int32** | Number of successful requests (2xx status code). | [optional] 
**FailedOps** | Pointer to **int32** | Number of failed requests (all non 2xx status codes). | [optional] 

## Methods

### NewBatchResponse

`func NewBatchResponse() *BatchResponse`

NewBatchResponse instantiates a new BatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseWithDefaults

`func NewBatchResponseWithDefaults() *BatchResponse`

NewBatchResponseWithDefaults instantiates a new BatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponses

`func (o *BatchResponse) GetResponses() []BatchOperationResponse`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *BatchResponse) GetResponsesOk() (*[]BatchOperationResponse, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *BatchResponse) SetResponses(v []BatchOperationResponse)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *BatchResponse) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### GetSuccessfulOps

`func (o *BatchResponse) GetSuccessfulOps() int32`

GetSuccessfulOps returns the SuccessfulOps field if non-nil, zero value otherwise.

### GetSuccessfulOpsOk

`func (o *BatchResponse) GetSuccessfulOpsOk() (*int32, bool)`

GetSuccessfulOpsOk returns a tuple with the SuccessfulOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulOps

`func (o *BatchResponse) SetSuccessfulOps(v int32)`

SetSuccessfulOps sets SuccessfulOps field to given value.

### HasSuccessfulOps

`func (o *BatchResponse) HasSuccessfulOps() bool`

HasSuccessfulOps returns a boolean if a field has been set.

### GetFailedOps

`func (o *BatchResponse) GetFailedOps() int32`

GetFailedOps returns the FailedOps field if non-nil, zero value otherwise.

### GetFailedOpsOk

`func (o *BatchResponse) GetFailedOpsOk() (*int32, bool)`

GetFailedOpsOk returns a tuple with the FailedOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedOps

`func (o *BatchResponse) SetFailedOps(v int32)`

SetFailedOps sets FailedOps field to given value.

### HasFailedOps

`func (o *BatchResponse) HasFailedOps() bool`

HasFailedOps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


