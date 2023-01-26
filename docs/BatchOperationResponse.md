# BatchOperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | HTTP status code. | [optional] 
**Headers** | Pointer to **map[string]string** | HTTP response header. | [optional] 
**Body** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewBatchOperationResponse

`func NewBatchOperationResponse() *BatchOperationResponse`

NewBatchOperationResponse instantiates a new BatchOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchOperationResponseWithDefaults

`func NewBatchOperationResponseWithDefaults() *BatchOperationResponse`

NewBatchOperationResponseWithDefaults instantiates a new BatchOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BatchOperationResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchOperationResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchOperationResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BatchOperationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHeaders

`func (o *BatchOperationResponse) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BatchOperationResponse) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BatchOperationResponse) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BatchOperationResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *BatchOperationResponse) GetBody() interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BatchOperationResponse) GetBodyOk() (*interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BatchOperationResponse) SetBody(v interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *BatchOperationResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *BatchOperationResponse) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *BatchOperationResponse) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


