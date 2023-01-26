# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of a resource. | [optional] [readonly] 
**Title** | Pointer to **string** | A short, human-readable summary. | [optional] 
**ReasonCode** | Pointer to **string** | The ReasonCode is used to identify the technical reason for the current notification. A list of reason codes can be found in the API documentation.  | [optional] 
**PublishedAt** | Pointer to **time.Time** | Publishing time of the notification. | [optional] 
**Reference** | Pointer to [**NotificationReference**](NotificationReference.md) |  | [optional] 

## Methods

### NewNotification

`func NewNotification() *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Notification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Notification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Notification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Notification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Notification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Notification) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Notification) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Notification) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetReasonCode

`func (o *Notification) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *Notification) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *Notification) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *Notification) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetPublishedAt

`func (o *Notification) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *Notification) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *Notification) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *Notification) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetReference

`func (o *Notification) GetReference() NotificationReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Notification) GetReferenceOk() (*NotificationReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Notification) SetReference(v NotificationReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Notification) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


