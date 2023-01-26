/*
IP RL BSP API

IP RL BSP API for participation in capacity/energy market tenders.

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BatchOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchOperationResponse{}

// BatchOperationResponse Response of a single batch operation.
type BatchOperationResponse struct {
	// HTTP status code.
	Status *int32 `json:"status,omitempty"`
	// HTTP response header.
	Headers *map[string]string `json:"headers,omitempty"`
	Body interface{} `json:"body,omitempty"`
}

// NewBatchOperationResponse instantiates a new BatchOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchOperationResponse() *BatchOperationResponse {
	this := BatchOperationResponse{}
	return &this
}

// NewBatchOperationResponseWithDefaults instantiates a new BatchOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchOperationResponseWithDefaults() *BatchOperationResponse {
	this := BatchOperationResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BatchOperationResponse) GetStatus() int32 {
	if o == nil || isNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchOperationResponse) GetStatusOk() (*int32, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BatchOperationResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *BatchOperationResponse) SetStatus(v int32) {
	o.Status = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *BatchOperationResponse) GetHeaders() map[string]string {
	if o == nil || isNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchOperationResponse) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *BatchOperationResponse) HasHeaders() bool {
	if o != nil && !isNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *BatchOperationResponse) SetHeaders(v map[string]string) {
	o.Headers = &v
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BatchOperationResponse) GetBody() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchOperationResponse) GetBodyOk() (*interface{}, bool) {
	if o == nil || isNil(o.Body) {
		return nil, false
	}
	return &o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BatchOperationResponse) HasBody() bool {
	if o != nil && isNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given interface{} and assigns it to the Body field.
func (o *BatchOperationResponse) SetBody(v interface{}) {
	o.Body = v
}

func (o BatchOperationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

type NullableBatchOperationResponse struct {
	value *BatchOperationResponse
	isSet bool
}

func (v NullableBatchOperationResponse) Get() *BatchOperationResponse {
	return v.value
}

func (v *NullableBatchOperationResponse) Set(val *BatchOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchOperationResponse(val *BatchOperationResponse) *NullableBatchOperationResponse {
	return &NullableBatchOperationResponse{value: val, isSet: true}
}

func (v NullableBatchOperationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

