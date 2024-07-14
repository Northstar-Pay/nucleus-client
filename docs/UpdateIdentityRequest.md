# UpdateIdentityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityType** | **string** |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**OtherNames** | **string** |  | 
**Gender** | **string** |  | 
**Dob** | **string** |  | 
**EmailAddress** | **string** |  | 
**PhoneNumber** | **string** |  | 
**Nationality** | **string** |  | 
**Street** | **string** |  | 
**Country** | **string** |  | 
**State** | **string** |  | 
**PostCode** | **string** |  | 
**City** | **string** |  | 
**CreatedAt** | **string** |  | 
**MetaData** | [**MetaData2**](MetaData2.md) |  | 

## Methods

### NewUpdateIdentityRequest

`func NewUpdateIdentityRequest(identityType string, firstName string, lastName string, otherNames string, gender string, dob string, emailAddress string, phoneNumber string, nationality string, street string, country string, state string, postCode string, city string, createdAt string, metaData MetaData2, ) *UpdateIdentityRequest`

NewUpdateIdentityRequest instantiates a new UpdateIdentityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIdentityRequestWithDefaults

`func NewUpdateIdentityRequestWithDefaults() *UpdateIdentityRequest`

NewUpdateIdentityRequestWithDefaults instantiates a new UpdateIdentityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityType

`func (o *UpdateIdentityRequest) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *UpdateIdentityRequest) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *UpdateIdentityRequest) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.


### GetFirstName

`func (o *UpdateIdentityRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateIdentityRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateIdentityRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UpdateIdentityRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateIdentityRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateIdentityRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetOtherNames

`func (o *UpdateIdentityRequest) GetOtherNames() string`

GetOtherNames returns the OtherNames field if non-nil, zero value otherwise.

### GetOtherNamesOk

`func (o *UpdateIdentityRequest) GetOtherNamesOk() (*string, bool)`

GetOtherNamesOk returns a tuple with the OtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherNames

`func (o *UpdateIdentityRequest) SetOtherNames(v string)`

SetOtherNames sets OtherNames field to given value.


### GetGender

`func (o *UpdateIdentityRequest) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UpdateIdentityRequest) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UpdateIdentityRequest) SetGender(v string)`

SetGender sets Gender field to given value.


### GetDob

`func (o *UpdateIdentityRequest) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *UpdateIdentityRequest) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *UpdateIdentityRequest) SetDob(v string)`

SetDob sets Dob field to given value.


### GetEmailAddress

`func (o *UpdateIdentityRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UpdateIdentityRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UpdateIdentityRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetPhoneNumber

`func (o *UpdateIdentityRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UpdateIdentityRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UpdateIdentityRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetNationality

`func (o *UpdateIdentityRequest) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *UpdateIdentityRequest) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *UpdateIdentityRequest) SetNationality(v string)`

SetNationality sets Nationality field to given value.


### GetStreet

`func (o *UpdateIdentityRequest) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *UpdateIdentityRequest) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *UpdateIdentityRequest) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetCountry

`func (o *UpdateIdentityRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateIdentityRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateIdentityRequest) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetState

`func (o *UpdateIdentityRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateIdentityRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateIdentityRequest) SetState(v string)`

SetState sets State field to given value.


### GetPostCode

`func (o *UpdateIdentityRequest) GetPostCode() string`

GetPostCode returns the PostCode field if non-nil, zero value otherwise.

### GetPostCodeOk

`func (o *UpdateIdentityRequest) GetPostCodeOk() (*string, bool)`

GetPostCodeOk returns a tuple with the PostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCode

`func (o *UpdateIdentityRequest) SetPostCode(v string)`

SetPostCode sets PostCode field to given value.


### GetCity

`func (o *UpdateIdentityRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UpdateIdentityRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UpdateIdentityRequest) SetCity(v string)`

SetCity sets City field to given value.


### GetCreatedAt

`func (o *UpdateIdentityRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateIdentityRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateIdentityRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetMetaData

`func (o *UpdateIdentityRequest) GetMetaData() MetaData2`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *UpdateIdentityRequest) GetMetaDataOk() (*MetaData2, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *UpdateIdentityRequest) SetMetaData(v MetaData2)`

SetMetaData sets MetaData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


