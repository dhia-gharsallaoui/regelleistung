# \DocumentsApi

All URIs are relative to *https://api.regelleistung.net/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDocument**](DocumentsApi.md#GetDocument) | **Get** /documents/{id} | Returns a document by its id.
[**GetDocumentContent**](DocumentsApi.md#GetDocumentContent) | **Get** /documents/{id}/content | Returns the content of requested document as \&quot;file download\&quot;
[**GetDocuments**](DocumentsApi.md#GetDocuments) | **Get** /documents | Result documents



## GetDocument

> ResultDocument GetDocument(ctx, id).Execute()

Returns a document by its id.

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
    id := "id_example" // string | Document identification

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.GetDocument(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocument`: ResultDocument
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Document identification | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultDocument**](ResultDocument.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentContent

> string GetDocumentContent(ctx, id).Execute()

Returns the content of requested document as \"file download\"

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
    id := "id_example" // string | Document identification

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.GetDocumentContent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetDocumentContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentContent`: string
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetDocumentContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Document identification | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocuments

> []ResultDocument GetDocuments(ctx).DeliveryDate(deliveryDate).Tender(tender).Market(market).ProductType(productType).Tag(tag).Offset(offset).Limit(limit).Execute()

Result documents



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
    deliveryDate := time.Now() // string | Only documents for the specified delivery date are listed (ISO 8601 format YYYY-MM-DD). If not specified and also no tender ID is specified, then only documents for today's delivery day will be listed. (optional)
    tender := "tender_example" // string | Tender Id filter. (optional)
    market := openapiclient.ReserveMarket("CAPACITY") // ReserveMarket | If not specified and also no tender ID is specified, documents for all markets are returned. (optional)
    productType := openapiclient.ProductType("aFRR") // ProductType | Product type filter. If this query parameter is not set and also no tender ID is specified, documents for all product types are returned. (optional)
    tag := "tag_example" // string | Tag filter. If this query parameter is not set, values for all tags are returned. (optional)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.GetDocuments(context.Background()).DeliveryDate(deliveryDate).Tender(tender).Market(market).ProductType(productType).Tag(tag).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocuments`: []ResultDocument
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deliveryDate** | **string** | Only documents for the specified delivery date are listed (ISO 8601 format YYYY-MM-DD). If not specified and also no tender ID is specified, then only documents for today&#39;s delivery day will be listed. | 
 **tender** | **string** | Tender Id filter. | 
 **market** | [**ReserveMarket**](ReserveMarket.md) | If not specified and also no tender ID is specified, documents for all markets are returned. | 
 **productType** | [**ProductType**](ProductType.md) | Product type filter. If this query parameter is not set and also no tender ID is specified, documents for all product types are returned. | 
 **tag** | **string** | Tag filter. If this query parameter is not set, values for all tags are returned. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 100]

### Return type

[**[]ResultDocument**](ResultDocument.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

