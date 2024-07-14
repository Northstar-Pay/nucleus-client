# CreateledgerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**MetaData** | [**MetaData**](MetaData.md) |  | 

## Methods

### NewCreateledgerRequest

`func NewCreateledgerRequest(name string, metaData MetaData, ) *CreateledgerRequest`

NewCreateledgerRequest instantiates a new CreateledgerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateledgerRequestWithDefaults

`func NewCreateledgerRequestWithDefaults() *CreateledgerRequest`

NewCreateledgerRequestWithDefaults instantiates a new CreateledgerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateledgerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateledgerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateledgerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMetaData

`func (o *CreateledgerRequest) GetMetaData() MetaData`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *CreateledgerRequest) GetMetaDataOk() (*MetaData, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *CreateledgerRequest) SetMetaData(v MetaData)`

SetMetaData sets MetaData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


