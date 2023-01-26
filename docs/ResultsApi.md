# \ResultsApi

All URIs are relative to *https://api.regelleistung.net/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllocationResults**](ResultsApi.md#GetAllocationResults) | **Get** /results | Anonymous results



## GetAllocationResults

> []AllocationResultElement GetAllocationResults(ctx).DeliveryDate(deliveryDate).Market(market).Tender(tender).ProductType(productType).ProductName(productName).State(state).Zone(zone).Offset(offset).Limit(limit).Execute()

Anonymous results



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
    deliveryDate := time.Now() // string | Only results for the specified delivery date will be listed (ISO 8601 format YYYY-MM-DD). If not specified and also no tender ID is specified, then only results for today's delivery day will be listed. (optional)
    market := openapiclient.ReserveMarket("CAPACITY") // ReserveMarket | Market filter. If not specified and also no tender ID is specified, then only bids for capacity market will be listed. (optional)
    tender := "tender_example" // string | By specifying the tender ID, only results of this tender will be listed. (optional)
    productType := openapiclient.ProductType("aFRR") // ProductType | Product type filter. If this query parameter is not set and also no tender ID is specified, results for all product types are returned. (optional)
    productName := openapiclient.ProductName("NEG_00_04") // ProductName |  (optional)
    state := "state_example" // string |  (optional)
    zone := "DE" // string |  (optional)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResultsApi.GetAllocationResults(context.Background()).DeliveryDate(deliveryDate).Market(market).Tender(tender).ProductType(productType).ProductName(productName).State(state).Zone(zone).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResultsApi.GetAllocationResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllocationResults`: []AllocationResultElement
    fmt.Fprintf(os.Stdout, "Response from `ResultsApi.GetAllocationResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllocationResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deliveryDate** | **string** | Only results for the specified delivery date will be listed (ISO 8601 format YYYY-MM-DD). If not specified and also no tender ID is specified, then only results for today&#39;s delivery day will be listed. | 
 **market** | [**ReserveMarket**](ReserveMarket.md) | Market filter. If not specified and also no tender ID is specified, then only bids for capacity market will be listed. | 
 **tender** | **string** | By specifying the tender ID, only results of this tender will be listed. | 
 **productType** | [**ProductType**](ProductType.md) | Product type filter. If this query parameter is not set and also no tender ID is specified, results for all product types are returned. | 
 **productName** | [**ProductName**](ProductName.md) |  | 
 **state** | **string** |  | 
 **zone** | **string** |  | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 100]

### Return type

[**[]AllocationResultElement**](AllocationResultElement.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

