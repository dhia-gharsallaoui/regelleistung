# \TendersApi

All URIs are relative to *https://api.regelleistung.net/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTender**](TendersApi.md#GetTender) | **Get** /tenders/{id} | Get tender
[**GetTenderDemand**](TendersApi.md#GetTenderDemand) | **Get** /tenders/{id}/demands | Demand of tender
[**GetTenderList**](TendersApi.md#GetTenderList) | **Get** /tenders | List tenders (optionally capacity/energy market)



## GetTender

> Tender GetTender(ctx, id).Execute()

Get tender

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
    id := "id_example" // string | Identificiation of requested tender. (For details on the new tender id format, see the [reference guide.](/docs/guide#tender-id))

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TendersApi.GetTender(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TendersApi.GetTender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTender`: Tender
    fmt.Fprintf(os.Stdout, "Response from `TendersApi.GetTender`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identificiation of requested tender. (For details on the new tender id format, see the [reference guide.](/docs/guide#tender-id)) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tender**](Tender.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenderDemand

> []Demand GetTenderDemand(ctx, id).Execute()

Demand of tender



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
    id := "id_example" // string | Identificiation of requested tender. (For details on the new tender id format, see the [reference guide.](/docs/guide#tender-id))

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TendersApi.GetTenderDemand(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TendersApi.GetTenderDemand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenderDemand`: []Demand
    fmt.Fprintf(os.Stdout, "Response from `TendersApi.GetTenderDemand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identificiation of requested tender. (For details on the new tender id format, see the [reference guide.](/docs/guide#tender-id)) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenderDemandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Demand**](Demand.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenderList

> []Tender GetTenderList(ctx).DeliveryDate(deliveryDate).Market(market).ProductType(productType).Execute()

List tenders (optionally capacity/energy market)



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
    deliveryDate := time.Now() // string | Only tenders for the specified delivery date are listed (ISO 8601 format YYYY-MM-DD). If not specified, then only tenders for today and tomorrow are listed. (optional)
    market := openapiclient.ReserveMarket("CAPACITY") // ReserveMarket | If not specified, then only tenders for capacity market will be listed. (optional)
    productType := openapiclient.ProductType("aFRR") // ProductType | Product type filter. If this query parameter is not set, tenders for all product types are returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TendersApi.GetTenderList(context.Background()).DeliveryDate(deliveryDate).Market(market).ProductType(productType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TendersApi.GetTenderList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenderList`: []Tender
    fmt.Fprintf(os.Stdout, "Response from `TendersApi.GetTenderList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenderListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deliveryDate** | **string** | Only tenders for the specified delivery date are listed (ISO 8601 format YYYY-MM-DD). If not specified, then only tenders for today and tomorrow are listed. | 
 **market** | [**ReserveMarket**](ReserveMarket.md) | If not specified, then only tenders for capacity market will be listed. | 
 **productType** | [**ProductType**](ProductType.md) | Product type filter. If this query parameter is not set, tenders for all product types are returned. | 

### Return type

[**[]Tender**](Tender.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

