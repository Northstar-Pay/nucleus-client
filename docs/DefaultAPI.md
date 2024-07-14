# \DefaultAPI

All URIs are relative to *http://localhost:5001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackDB**](DefaultAPI.md#BackDB) | **Post** /backup | Back DB
[**BackDBToS3**](DefaultAPI.md#BackDBToS3) | **Post** /backup/s3 | Back DB To S3
[**CreateAccount**](DefaultAPI.md#CreateAccount) | **Post** /accounts | Create Account
[**CreateBalance**](DefaultAPI.md#CreateBalance) | **Post** /balances | Create Balance
[**CreateBalanceMonitor**](DefaultAPI.md#CreateBalanceMonitor) | **Post** /balance-monitors | Create Balance Monitor
[**CreateIdentity**](DefaultAPI.md#CreateIdentity) | **Post** /identities | Create Identity
[**Createledger**](DefaultAPI.md#Createledger) | **Post** /ledgers | Create ledger
[**GetAccount**](DefaultAPI.md#GetAccount) | **Get** /accounts/{id} | Get Account
[**GetBalance**](DefaultAPI.md#GetBalance) | **Get** /balances/{id} | Get Balance
[**GetBalanceMonitor**](DefaultAPI.md#GetBalanceMonitor) | **Get** /balance-monitors/{id} | Get Balance Monitor
[**GetIdentity**](DefaultAPI.md#GetIdentity) | **Get** /identities/{id} | Get Identity
[**GetLedger**](DefaultAPI.md#GetLedger) | **Get** /ledgers/{id} | Get Ledger
[**RecordTransaction**](DefaultAPI.md#RecordTransaction) | **Post** /transactions | Record Transaction
[**RefundTransaction**](DefaultAPI.md#RefundTransaction) | **Post** /refund-transaction/{id} | Refund Transaction
[**UpdateAccount**](DefaultAPI.md#UpdateAccount) | **Put** /accounts/{id} | Update Account
[**UpdateBalanceMonitor**](DefaultAPI.md#UpdateBalanceMonitor) | **Put** /balance-monitors/{id} | Update Balance Monitor
[**UpdateIdentity**](DefaultAPI.md#UpdateIdentity) | **Put** /identities/{id} | Update Identity
[**UpdateInflightTransaction**](DefaultAPI.md#UpdateInflightTransaction) | **Put** /transactions/inflight/{txID}/{status} | Update Inflight Transaction



## BackDB

> BackDB201Response BackDB(ctx).Execute()

Back DB

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.BackDB(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BackDB``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackDB`: BackDB201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BackDB`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBackDBRequest struct via the builder pattern


### Return type

[**BackDB201Response**](BackDB201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackDBToS3

> BackDB201Response BackDBToS3(ctx).Execute()

Back DB To S3

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.BackDBToS3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BackDBToS3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackDBToS3`: BackDB201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BackDBToS3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBackDBToS3Request struct via the builder pattern


### Return type

[**BackDB201Response**](BackDB201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccount

> CreateAccount201Response CreateAccount(ctx).CreateAccountRequest(createAccountRequest).Execute()

Create Account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createAccountRequest := *openapiclient.NewCreateAccountRequest("BankName_example", "Number_example", "BalanceId_example", "IdentityId_example") // CreateAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateAccount(context.Background()).CreateAccountRequest(createAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccount`: CreateAccount201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountRequest** | [**CreateAccountRequest**](CreateAccountRequest.md) |  | 

### Return type

[**CreateAccount201Response**](CreateAccount201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBalance

> CreateBalance201Response CreateBalance(ctx).CreateBalanceRequest(createBalanceRequest).Execute()

Create Balance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createBalanceRequest := *openapiclient.NewCreateBalanceRequest("LedgerId_example", "Currency_example") // CreateBalanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateBalance(context.Background()).CreateBalanceRequest(createBalanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBalance`: CreateBalance201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBalanceRequest** | [**CreateBalanceRequest**](CreateBalanceRequest.md) |  | 

### Return type

[**CreateBalance201Response**](CreateBalance201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBalanceMonitor

> CreateBalanceMonitor201Response CreateBalanceMonitor(ctx).CreateBalanceMonitorRequest(createBalanceMonitorRequest).Execute()

Create Balance Monitor

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createBalanceMonitorRequest := *openapiclient.NewCreateBalanceMonitorRequest("BalanceId_example", *openapiclient.NewCondition("Field_example", "Operator_example", int32(123)), "CallBackUrl_example") // CreateBalanceMonitorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateBalanceMonitor(context.Background()).CreateBalanceMonitorRequest(createBalanceMonitorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateBalanceMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBalanceMonitor`: CreateBalanceMonitor201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateBalanceMonitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBalanceMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBalanceMonitorRequest** | [**CreateBalanceMonitorRequest**](CreateBalanceMonitorRequest.md) |  | 

### Return type

[**CreateBalanceMonitor201Response**](CreateBalanceMonitor201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdentity

> CreateIdentity201Response CreateIdentity(ctx).CreateIdentityRequest(createIdentityRequest).Execute()

Create Identity

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createIdentityRequest := *openapiclient.NewCreateIdentityRequest("IdentityType_example", "Name_example", "Category_example", "Street_example", "Country_example", "State_example", "PostCode_example", "City_example", "CreatedAt_example", *openapiclient.NewMetaData2(false, "Reference_example")) // CreateIdentityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateIdentity(context.Background()).CreateIdentityRequest(createIdentityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdentity`: CreateIdentity201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateIdentity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIdentityRequest** | [**CreateIdentityRequest**](CreateIdentityRequest.md) |  | 

### Return type

[**CreateIdentity201Response**](CreateIdentity201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Createledger

> Createledger201Response Createledger(ctx).CreateledgerRequest(createledgerRequest).Execute()

Create ledger

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createledgerRequest := *openapiclient.NewCreateledgerRequest("Name_example", *openapiclient.NewMetaData("ProjectOwner_example")) // CreateledgerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Createledger(context.Background()).CreateledgerRequest(createledgerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Createledger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Createledger`: Createledger201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Createledger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateledgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createledgerRequest** | [**CreateledgerRequest**](CreateledgerRequest.md) |  | 

### Return type

[**Createledger201Response**](Createledger201Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> GetAccount200Response GetAccount(ctx, id).Include(include).Execute()

Get Account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	include := "identity" // string | 
	id := "acc_463b25b8-7b24-4e8b-aa49-7844d9074019" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetAccount(context.Background(), id).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: GetAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** |  | 


### Return type

[**GetAccount200Response**](GetAccount200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalance

> GetBalance200Response GetBalance(ctx, id).Include(include).Execute()

Get Balance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	include := "ledger" // string | 
	id := "bln_0be360ca-86fe-457d-be43-daa3f966d8f0" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetBalance(context.Background(), id).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalance`: GetBalance200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** |  | 


### Return type

[**GetBalance200Response**](GetBalance200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalanceMonitor

> GetBalanceMonitor200Response GetBalanceMonitor(ctx, id).Execute()

Get Balance Monitor

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "mon_e0e77b0c-4985-472a-9bf5-76a48b0259b0" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetBalanceMonitor(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetBalanceMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalanceMonitor`: GetBalanceMonitor200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetBalanceMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBalanceMonitor200Response**](GetBalanceMonitor200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentity

> GetIdentity200Response GetIdentity(ctx, id).Execute()

Get Identity

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "idt_0b5f333d-e0b8-4422-8e0e-5d41a767f1db" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetIdentity(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentity`: GetIdentity200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIdentity200Response**](GetIdentity200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLedger

> GetLedger200Response GetLedger(ctx, id).Execute()

Get Ledger

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "ldg_073f7ffe-9dfd-42ce-aa50-d1dca1788adc" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetLedger(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetLedger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedger`: GetLedger200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetLedger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLedger200Response**](GetLedger200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordTransaction

> RecordTransaction201Response RecordTransaction(ctx).RecordTransactionRequest(recordTransactionRequest).Execute()

Record Transaction

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	recordTransactionRequest := *openapiclient.NewRecordTransactionRequest(int32(123), "Reference_example", []openapiclient.Source{*openapiclient.NewSource("Identifier_example", "Distribution_example", "Narration_example")}, []openapiclient.Destination{*openapiclient.NewDestination("Identifier_example", "Distribution_example", "Narration_example")}, "Currency_example", "Source_example", "Destination_example", false, false) // RecordTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RecordTransaction(context.Background()).RecordTransactionRequest(recordTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RecordTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordTransaction`: RecordTransaction201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RecordTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecordTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordTransactionRequest** | [**RecordTransactionRequest**](RecordTransactionRequest.md) |  | 

### Return type

[**RecordTransaction201Response**](RecordTransaction201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundTransaction

> RefundTransaction201Response RefundTransaction(ctx, id).Execute()

Refund Transaction

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "txn_5bbbe4d3-2d82-4da1-8191-aaa491d025de" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RefundTransaction(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RefundTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundTransaction`: RefundTransaction201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RefundTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefundTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RefundTransaction201Response**](RefundTransaction201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> CreateAccount201Response UpdateAccount(ctx, id).UpdateAccountRequest(updateAccountRequest).Execute()

Update Account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	updateAccountRequest := *openapiclient.NewUpdateAccountRequest("BankName_example", "Number_example", "BalanceId_example", "IdentityId_example") // UpdateAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateAccount(context.Background(), id).UpdateAccountRequest(updateAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccount`: CreateAccount201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAccountRequest** | [**UpdateAccountRequest**](UpdateAccountRequest.md) |  | 

### Return type

[**CreateAccount201Response**](CreateAccount201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBalanceMonitor

> UpdateBalanceMonitor(ctx, id).UpdateBalanceMonitorRequest(updateBalanceMonitorRequest).Execute()

Update Balance Monitor

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "mon_da573822-902c-4349-86c9-2c02ab697c6e" // string | 
	updateBalanceMonitorRequest := *openapiclient.NewUpdateBalanceMonitorRequest("BalanceId_example", *openapiclient.NewCondition("Field_example", "Operator_example", int32(123))) // UpdateBalanceMonitorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateBalanceMonitor(context.Background(), id).UpdateBalanceMonitorRequest(updateBalanceMonitorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateBalanceMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBalanceMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBalanceMonitorRequest** | [**UpdateBalanceMonitorRequest**](UpdateBalanceMonitorRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentity

> UpdateIdentity200Response UpdateIdentity(ctx, id).UpdateIdentityRequest(updateIdentityRequest).Execute()

Update Identity

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "idt_0b5f333d-e0b8-4422-8e0e-5d41a767f1db" // string | 
	updateIdentityRequest := *openapiclient.NewUpdateIdentityRequest("IdentityType_example", "FirstName_example", "LastName_example", "OtherNames_example", "Gender_example", "Dob_example", "EmailAddress_example", "PhoneNumber_example", "Nationality_example", "Street_example", "Country_example", "State_example", "PostCode_example", "City_example", "CreatedAt_example", *openapiclient.NewMetaData2(false, "Reference_example")) // UpdateIdentityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateIdentity(context.Background(), id).UpdateIdentityRequest(updateIdentityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIdentity`: UpdateIdentity200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIdentityRequest** | [**UpdateIdentityRequest**](UpdateIdentityRequest.md) |  | 

### Return type

[**UpdateIdentity200Response**](UpdateIdentity200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInflightTransaction

> UpdateInflightTransaction200Response UpdateInflightTransaction(ctx, txID, status).Execute()

Update Inflight Transaction

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	txID := "txn_d0440bb5-e2c2-4937-bb4c-e91b50b7df22" // string | 
	status := "commit" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateInflightTransaction(context.Background(), txID, status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateInflightTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInflightTransaction`: UpdateInflightTransaction200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateInflightTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txID** | **string** |  | 
**status** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInflightTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UpdateInflightTransaction200Response**](UpdateInflightTransaction200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

