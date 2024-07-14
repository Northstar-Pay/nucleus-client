# CreateAccount201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Name** | **string** |  | 
**Number** | **string** |  | 
**BankName** | **string** |  | 
**Currency** | **string** |  | 
**CreatedAt** | **string** |  | 
**BalanceId** | **string** |  | 
**IdentityId** | **string** |  | 
**LedgerId** | **string** |  | 
**MetaData** | **NullableString** |  | 
**Ledger** | **NullableString** |  | 
**Balance** | **NullableString** |  | 
**Identity** | **NullableString** |  | 

## Methods

### NewCreateAccount201Response

`func NewCreateAccount201Response(accountId string, name string, number string, bankName string, currency string, createdAt string, balanceId string, identityId string, ledgerId string, metaData NullableString, ledger NullableString, balance NullableString, identity NullableString, ) *CreateAccount201Response`

NewCreateAccount201Response instantiates a new CreateAccount201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccount201ResponseWithDefaults

`func NewCreateAccount201ResponseWithDefaults() *CreateAccount201Response`

NewCreateAccount201ResponseWithDefaults instantiates a new CreateAccount201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateAccount201Response) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateAccount201Response) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateAccount201Response) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *CreateAccount201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccount201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccount201Response) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *CreateAccount201Response) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CreateAccount201Response) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CreateAccount201Response) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetBankName

`func (o *CreateAccount201Response) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *CreateAccount201Response) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *CreateAccount201Response) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetCurrency

`func (o *CreateAccount201Response) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateAccount201Response) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateAccount201Response) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCreatedAt

`func (o *CreateAccount201Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateAccount201Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateAccount201Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetBalanceId

`func (o *CreateAccount201Response) GetBalanceId() string`

GetBalanceId returns the BalanceId field if non-nil, zero value otherwise.

### GetBalanceIdOk

`func (o *CreateAccount201Response) GetBalanceIdOk() (*string, bool)`

GetBalanceIdOk returns a tuple with the BalanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceId

`func (o *CreateAccount201Response) SetBalanceId(v string)`

SetBalanceId sets BalanceId field to given value.


### GetIdentityId

`func (o *CreateAccount201Response) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *CreateAccount201Response) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *CreateAccount201Response) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetLedgerId

`func (o *CreateAccount201Response) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *CreateAccount201Response) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *CreateAccount201Response) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.


### GetMetaData

`func (o *CreateAccount201Response) GetMetaData() string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *CreateAccount201Response) GetMetaDataOk() (*string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *CreateAccount201Response) SetMetaData(v string)`

SetMetaData sets MetaData field to given value.


### SetMetaDataNil

`func (o *CreateAccount201Response) SetMetaDataNil(b bool)`

 SetMetaDataNil sets the value for MetaData to be an explicit nil

### UnsetMetaData
`func (o *CreateAccount201Response) UnsetMetaData()`

UnsetMetaData ensures that no value is present for MetaData, not even an explicit nil
### GetLedger

`func (o *CreateAccount201Response) GetLedger() string`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *CreateAccount201Response) GetLedgerOk() (*string, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *CreateAccount201Response) SetLedger(v string)`

SetLedger sets Ledger field to given value.


### SetLedgerNil

`func (o *CreateAccount201Response) SetLedgerNil(b bool)`

 SetLedgerNil sets the value for Ledger to be an explicit nil

### UnsetLedger
`func (o *CreateAccount201Response) UnsetLedger()`

UnsetLedger ensures that no value is present for Ledger, not even an explicit nil
### GetBalance

`func (o *CreateAccount201Response) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CreateAccount201Response) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CreateAccount201Response) SetBalance(v string)`

SetBalance sets Balance field to given value.


### SetBalanceNil

`func (o *CreateAccount201Response) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *CreateAccount201Response) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetIdentity

`func (o *CreateAccount201Response) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CreateAccount201Response) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CreateAccount201Response) SetIdentity(v string)`

SetIdentity sets Identity field to given value.


### SetIdentityNil

`func (o *CreateAccount201Response) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *CreateAccount201Response) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


