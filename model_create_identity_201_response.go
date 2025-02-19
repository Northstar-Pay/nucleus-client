/*
Blnk

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateIdentity201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateIdentity201Response{}

// CreateIdentity201Response struct for CreateIdentity201Response
type CreateIdentity201Response struct {
	IdentityId       string    `json:"identity_id"`
	IdentityType     string    `json:"identity_type"`
	OrganizationName string    `json:"organization_name"`
	Category         string    `json:"category"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	OtherNames       string    `json:"other_names"`
	Gender           string    `json:"gender"`
	Dob              string    `json:"dob"`
	EmailAddress     string    `json:"email_address"`
	PhoneNumber      string    `json:"phone_number"`
	Nationality      string    `json:"nationality"`
	Street           string    `json:"street"`
	Country          string    `json:"country"`
	State            string    `json:"state"`
	PostCode         string    `json:"post_code"`
	City             string    `json:"city"`
	CreatedAt        string    `json:"created_at"`
	MetaData         MetaData2 `json:"meta_data"`
}

type _CreateIdentity201Response CreateIdentity201Response

// NewCreateIdentity201Response instantiates a new CreateIdentity201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIdentity201Response(identityId string, identityType string, organizationName string, category string, firstName string, lastName string, otherNames string, gender string, dob string, emailAddress string, phoneNumber string, nationality string, street string, country string, state string, postCode string, city string, createdAt string, metaData MetaData2) *CreateIdentity201Response {
	this := CreateIdentity201Response{}
	this.IdentityId = identityId
	this.IdentityType = identityType
	this.OrganizationName = organizationName
	this.Category = category
	this.FirstName = firstName
	this.LastName = lastName
	this.OtherNames = otherNames
	this.Gender = gender
	this.Dob = dob
	this.EmailAddress = emailAddress
	this.PhoneNumber = phoneNumber
	this.Nationality = nationality
	this.Street = street
	this.Country = country
	this.State = state
	this.PostCode = postCode
	this.City = city
	this.CreatedAt = createdAt
	this.MetaData = metaData
	return &this
}

// NewCreateIdentity201ResponseWithDefaults instantiates a new CreateIdentity201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIdentity201ResponseWithDefaults() *CreateIdentity201Response {
	this := CreateIdentity201Response{}
	return &this
}

// GetIdentityId returns the IdentityId field value
func (o *CreateIdentity201Response) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *CreateIdentity201Response) SetIdentityId(v string) {
	o.IdentityId = v
}

// GetIdentityType returns the IdentityType field value
func (o *CreateIdentity201Response) GetIdentityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityType
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetIdentityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityType, true
}

// SetIdentityType sets field value
func (o *CreateIdentity201Response) SetIdentityType(v string) {
	o.IdentityType = v
}

// GetOrganizationName returns the OrganizationName field value
func (o *CreateIdentity201Response) GetOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationName, true
}

// SetOrganizationName sets field value
func (o *CreateIdentity201Response) SetOrganizationName(v string) {
	o.OrganizationName = v
}

// GetCategory returns the Category field value
func (o *CreateIdentity201Response) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *CreateIdentity201Response) SetCategory(v string) {
	o.Category = v
}

// GetFirstName returns the FirstName field value
func (o *CreateIdentity201Response) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *CreateIdentity201Response) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *CreateIdentity201Response) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *CreateIdentity201Response) SetLastName(v string) {
	o.LastName = v
}

// GetOtherNames returns the OtherNames field value
func (o *CreateIdentity201Response) GetOtherNames() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OtherNames
}

// GetOtherNamesOk returns a tuple with the OtherNames field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetOtherNamesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OtherNames, true
}

// SetOtherNames sets field value
func (o *CreateIdentity201Response) SetOtherNames(v string) {
	o.OtherNames = v
}

// GetGender returns the Gender field value
func (o *CreateIdentity201Response) GetGender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gender
}

// GetGenderOk returns a tuple with the Gender field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetGenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gender, true
}

// SetGender sets field value
func (o *CreateIdentity201Response) SetGender(v string) {
	o.Gender = v
}

// GetDob returns the Dob field value
func (o *CreateIdentity201Response) GetDob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dob
}

// GetDobOk returns a tuple with the Dob field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetDobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dob, true
}

// SetDob sets field value
func (o *CreateIdentity201Response) SetDob(v string) {
	o.Dob = v
}

// GetEmailAddress returns the EmailAddress field value
func (o *CreateIdentity201Response) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *CreateIdentity201Response) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *CreateIdentity201Response) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *CreateIdentity201Response) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetNationality returns the Nationality field value
func (o *CreateIdentity201Response) GetNationality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetNationalityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nationality, true
}

// SetNationality sets field value
func (o *CreateIdentity201Response) SetNationality(v string) {
	o.Nationality = v
}

// GetStreet returns the Street field value
func (o *CreateIdentity201Response) GetStreet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Street
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetStreetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Street, true
}

// SetStreet sets field value
func (o *CreateIdentity201Response) SetStreet(v string) {
	o.Street = v
}

// GetCountry returns the Country field value
func (o *CreateIdentity201Response) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *CreateIdentity201Response) SetCountry(v string) {
	o.Country = v
}

// GetState returns the State field value
func (o *CreateIdentity201Response) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateIdentity201Response) SetState(v string) {
	o.State = v
}

// GetPostCode returns the PostCode field value
func (o *CreateIdentity201Response) GetPostCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostCode
}

// GetPostCodeOk returns a tuple with the PostCode field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetPostCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostCode, true
}

// SetPostCode sets field value
func (o *CreateIdentity201Response) SetPostCode(v string) {
	o.PostCode = v
}

// GetCity returns the City field value
func (o *CreateIdentity201Response) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *CreateIdentity201Response) SetCity(v string) {
	o.City = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CreateIdentity201Response) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CreateIdentity201Response) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetMetaData returns the MetaData field value
func (o *CreateIdentity201Response) GetMetaData() MetaData2 {
	if o == nil {
		var ret MetaData2
		return ret
	}

	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity201Response) GetMetaDataOk() (*MetaData2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaData, true
}

// SetMetaData sets field value
func (o *CreateIdentity201Response) SetMetaData(v MetaData2) {
	o.MetaData = v
}

func (o CreateIdentity201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIdentity201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identity_id"] = o.IdentityId
	toSerialize["identity_type"] = o.IdentityType
	toSerialize["organization_name"] = o.OrganizationName
	toSerialize["category"] = o.Category
	toSerialize["first_name"] = o.FirstName
	toSerialize["last_name"] = o.LastName
	toSerialize["other_names"] = o.OtherNames
	toSerialize["gender"] = o.Gender
	toSerialize["dob"] = o.Dob
	toSerialize["email_address"] = o.EmailAddress
	toSerialize["phone_number"] = o.PhoneNumber
	toSerialize["nationality"] = o.Nationality
	toSerialize["street"] = o.Street
	toSerialize["country"] = o.Country
	toSerialize["state"] = o.State
	toSerialize["post_code"] = o.PostCode
	toSerialize["city"] = o.City
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["meta_data"] = o.MetaData
	return toSerialize, nil
}

func (o *CreateIdentity201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identity_id",
		"identity_type",
		"organization_name",
		"category",
		"first_name",
		"last_name",
		"other_names",
		"gender",
		"dob",
		"email_address",
		"phone_number",
		"nationality",
		"street",
		"country",
		"state",
		"post_code",
		"city",
		"created_at",
		"meta_data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateIdentity201Response := _CreateIdentity201Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateIdentity201Response)

	if err != nil {
		return err
	}

	*o = CreateIdentity201Response(varCreateIdentity201Response)

	return err
}

type NullableCreateIdentity201Response struct {
	value *CreateIdentity201Response
	isSet bool
}

func (v NullableCreateIdentity201Response) Get() *CreateIdentity201Response {
	return v.value
}

func (v *NullableCreateIdentity201Response) Set(val *CreateIdentity201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIdentity201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIdentity201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIdentity201Response(val *CreateIdentity201Response) *NullableCreateIdentity201Response {
	return &NullableCreateIdentity201Response{value: val, isSet: true}
}

func (v NullableCreateIdentity201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIdentity201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
