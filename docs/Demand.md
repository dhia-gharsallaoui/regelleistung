# Demand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectingZone** | Pointer to [**ConnectingZone**](ConnectingZone.md) |  | [optional] 
**ProductType** | Pointer to [**ProductType**](ProductType.md) |  | [optional] 
**ProductName** | Pointer to [**ProductName**](ProductName.md) |  | [optional] 
**Direction** | Pointer to [**Direction**](Direction.md) |  | [optional] 
**TimeInterval** | Pointer to [**TimeInterval**](TimeInterval.md) |  | [optional] 
**Data** | Pointer to [**DemandData**](DemandData.md) |  | [optional] 

## Methods

### NewDemand

`func NewDemand() *Demand`

NewDemand instantiates a new Demand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDemandWithDefaults

`func NewDemandWithDefaults() *Demand`

NewDemandWithDefaults instantiates a new Demand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectingZone

`func (o *Demand) GetConnectingZone() ConnectingZone`

GetConnectingZone returns the ConnectingZone field if non-nil, zero value otherwise.

### GetConnectingZoneOk

`func (o *Demand) GetConnectingZoneOk() (*ConnectingZone, bool)`

GetConnectingZoneOk returns a tuple with the ConnectingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectingZone

`func (o *Demand) SetConnectingZone(v ConnectingZone)`

SetConnectingZone sets ConnectingZone field to given value.

### HasConnectingZone

`func (o *Demand) HasConnectingZone() bool`

HasConnectingZone returns a boolean if a field has been set.

### GetProductType

`func (o *Demand) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *Demand) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *Demand) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *Demand) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetProductName

`func (o *Demand) GetProductName() ProductName`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *Demand) GetProductNameOk() (*ProductName, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *Demand) SetProductName(v ProductName)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *Demand) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetDirection

`func (o *Demand) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Demand) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Demand) SetDirection(v Direction)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Demand) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetTimeInterval

`func (o *Demand) GetTimeInterval() TimeInterval`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *Demand) GetTimeIntervalOk() (*TimeInterval, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *Demand) SetTimeInterval(v TimeInterval)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *Demand) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.

### GetData

`func (o *Demand) GetData() DemandData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Demand) GetDataOk() (*DemandData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Demand) SetData(v DemandData)`

SetData sets Data field to given value.

### HasData

`func (o *Demand) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


