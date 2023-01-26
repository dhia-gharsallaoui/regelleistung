# ConditionalLinkageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkedTo** | **string** | Bid identification | 
**Condition** | Pointer to [**LinkageCondition**](LinkageCondition.md) |  | [optional] [default to _01]

## Methods

### NewConditionalLinkageItem

`func NewConditionalLinkageItem(linkedTo string, ) *ConditionalLinkageItem`

NewConditionalLinkageItem instantiates a new ConditionalLinkageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionalLinkageItemWithDefaults

`func NewConditionalLinkageItemWithDefaults() *ConditionalLinkageItem`

NewConditionalLinkageItemWithDefaults instantiates a new ConditionalLinkageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkedTo

`func (o *ConditionalLinkageItem) GetLinkedTo() string`

GetLinkedTo returns the LinkedTo field if non-nil, zero value otherwise.

### GetLinkedToOk

`func (o *ConditionalLinkageItem) GetLinkedToOk() (*string, bool)`

GetLinkedToOk returns a tuple with the LinkedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTo

`func (o *ConditionalLinkageItem) SetLinkedTo(v string)`

SetLinkedTo sets LinkedTo field to given value.


### GetCondition

`func (o *ConditionalLinkageItem) GetCondition() LinkageCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ConditionalLinkageItem) GetConditionOk() (*LinkageCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ConditionalLinkageItem) SetCondition(v LinkageCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ConditionalLinkageItem) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


