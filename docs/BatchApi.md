# \BatchApi

All URIs are relative to *https://api.regelleistung.net/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Batch**](BatchApi.md#Batch) | **Post** /batch | Submit batch requests



## Batch

> BatchResponse Batch(ctx).BatchRequest(batchRequest).Execute()

Submit batch requests



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
    batchRequest := *openapiclient.NewBatchRequest([]openapiclient.BatchOperation{*openapiclient.NewBatchOperation("/bids/Bmauz8GLs012JaXc3HsW", openapiclient.BatchSupportedHttpMethod("POST"))}) // BatchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.Batch(context.Background()).BatchRequest(batchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.Batch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Batch`: BatchResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.Batch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchRequest** | [**BatchRequest**](BatchRequest.md) |  | 

### Return type

[**BatchResponse**](BatchResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

