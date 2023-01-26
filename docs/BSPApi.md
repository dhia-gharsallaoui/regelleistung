# \BSPApi

All URIs are relative to *https://api.regelleistung.net/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBackupContracts**](BSPApi.md#GetBackupContracts) | **Get** /backup-contracts | Returns current backup contracts
[**GetFrameworkContracts**](BSPApi.md#GetFrameworkContracts) | **Get** /framework-agreements | Current framework agreements
[**GetPreQualifiedCapacities**](BSPApi.md#GetPreQualifiedCapacities) | **Get** /pre-qualified-capacities/{deliveryDate} | Prequalified capacities.



## GetBackupContracts

> []BackupContract GetBackupContracts(ctx).ValidDate(validDate).ProductType(productType).Execute()

Returns current backup contracts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    validDate := time.Now() // string | Contracts for a specified day can be queried. If not specified, contracts for today will be listed. (optional)
    productType := "productType_example" // string | Product type filter for contracts. If not specified, then contracts for all product types are returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BSPApi.GetBackupContracts(context.Background()).ValidDate(validDate).ProductType(productType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSPApi.GetBackupContracts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackupContracts`: []BackupContract
    fmt.Fprintf(os.Stdout, "Response from `BSPApi.GetBackupContracts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validDate** | **string** | Contracts for a specified day can be queried. If not specified, contracts for today will be listed. | 
 **productType** | **string** | Product type filter for contracts. If not specified, then contracts for all product types are returned. | 

### Return type

[**[]BackupContract**](BackupContract.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFrameworkContracts

> []FrameworkAgreement GetFrameworkContracts(ctx).ProductType(productType).Execute()

Current framework agreements



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    productType := "productType_example" // string | Product type filter for framework agreements. If this query parameter is not set, contracts for all product types are returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BSPApi.GetFrameworkContracts(context.Background()).ProductType(productType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSPApi.GetFrameworkContracts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFrameworkContracts`: []FrameworkAgreement
    fmt.Fprintf(os.Stdout, "Response from `BSPApi.GetFrameworkContracts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFrameworkContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productType** | **string** | Product type filter for framework agreements. If this query parameter is not set, contracts for all product types are returned. | 

### Return type

[**[]FrameworkAgreement**](FrameworkAgreement.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPreQualifiedCapacities

> []PrequalifiedCapacity GetPreQualifiedCapacities(ctx, deliveryDate).ProductType(productType).Execute()

Prequalified capacities.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    deliveryDate := time.Now() // string | Defines the delivery date for which the pre-qualified capacities are to be returned.
    productType := "productType_example" // string | Product type filter for contracts. If not specified, then contracts for all product types are returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BSPApi.GetPreQualifiedCapacities(context.Background(), deliveryDate).ProductType(productType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSPApi.GetPreQualifiedCapacities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPreQualifiedCapacities`: []PrequalifiedCapacity
    fmt.Fprintf(os.Stdout, "Response from `BSPApi.GetPreQualifiedCapacities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryDate** | **string** | Defines the delivery date for which the pre-qualified capacities are to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPreQualifiedCapacitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productType** | **string** | Product type filter for contracts. If not specified, then contracts for all product types are returned. | 

### Return type

[**[]PrequalifiedCapacity**](PrequalifiedCapacity.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

