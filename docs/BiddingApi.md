# \BiddingApi

All URIs are relative to *https://api.regelleistung.net/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBid**](BiddingApi.md#DeleteBid) | **Delete** /bids/{id} | Deletes a simple, complex bid or a bid component
[**GetBid**](BiddingApi.md#GetBid) | **Get** /bids/{id} | Get a simple, complex bid or a complex bid component
[**GetBids**](BiddingApi.md#GetBids) | **Get** /bids | List submitted bids (optionally capacity or energy market)
[**SubmitBid**](BiddingApi.md#SubmitBid) | **Post** /bids | Submit a simple, complex bid or a complex bid component
[**UpdateBid**](BiddingApi.md#UpdateBid) | **Patch** /bids/{id} | Updates a simple, complex bid or a bid component



## DeleteBid

> DeleteBid(ctx, id).Execute()

Deletes a simple, complex bid or a bid component



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
    id := "id_example" // string | Bid identification

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BiddingApi.DeleteBid(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BiddingApi.DeleteBid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Bid identification | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBid

> Bid GetBid(ctx, id).Execute()

Get a simple, complex bid or a complex bid component



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
    id := "id_example" // string | Unique bid identification or group bid id of complex bid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BiddingApi.GetBid(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BiddingApi.GetBid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBid`: Bid
    fmt.Fprintf(os.Stdout, "Response from `BiddingApi.GetBid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique bid identification or group bid id of complex bid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Bid**](Bid.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBids

> []Bid GetBids(ctx).DeliveryDate(deliveryDate).Market(market).Tender(tender).ProductType(productType).ProductName(productName).State(state).Zone(zone).From(from).To(to).Type_(type_).Cid(cid).Gid(gid).Link(link).LinkedTo(linkedTo).Tag(tag).BackupsOnly(backupsOnly).BackupFor(backupFor).Offset(offset).Limit(limit).Execute()

List submitted bids (optionally capacity or energy market)



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
    deliveryDate := time.Now() // string | Only bids for the specified delivery date will be listed (ISO 8601 format YYYY-MM-DD). If not specified and also no tender ID is specified, then only bids for today's delivery day will be listed. (optional)
    market := openapiclient.ReserveMarket("CAPACITY") // ReserveMarket | Market filter. If not specified and also no tender ID is specified, then only bids for capacity market will be listed. (optional)
    tender := "tender_example" // string | By specifying the tender ID, only bids of this tender will be listed. For details on the tender id format, see the [reference guide.](/docs/guide#tender-id) (optional)
    productType := openapiclient.ProductType("aFRR") // ProductType | Product type filter. If this query parameter is not set and also no tender ID is specified, bids for all product types are returned. (optional)
    productName := openapiclient.ProductName("NEG_00_04") // ProductName | Product name filter. If this query parameter is not set, bids for all product names are returned. (optional)
    state := "state_example" // string |  (optional)
    zone := "zone_example" // string |  (optional)
    from := time.Now() // time.Time |  (optional)
    to := time.Now() // time.Time |  (optional)
    type_ := openapiclient.BidType("SIMPLE") // BidType | By specifying the type, only bids of this type will be listed. (optional) (default to "SIMPLE")
    cid := "cid_example" // string | Filters bids that have the given complex bid ID. (optional)
    gid := "gid_example" // string | Filters bids that have the given bid group ID. (optional)
    link := "link_example" // string | Filters bids that have the given technical link ID. (optional)
    linkedTo := "linkedTo_example" // string | Filters bids that have a conditional link to the given bid ID. (optional)
    tag := "tag_example" // string | Filters bids that have the given tag. (optional)
    backupsOnly := true // bool |  (optional) (default to false)
    backupFor := "backupFor_example" // string | Filters bids which are a backup for the provider with the given EIC (optional)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BiddingApi.GetBids(context.Background()).DeliveryDate(deliveryDate).Market(market).Tender(tender).ProductType(productType).ProductName(productName).State(state).Zone(zone).From(from).To(to).Type_(type_).Cid(cid).Gid(gid).Link(link).LinkedTo(linkedTo).Tag(tag).BackupsOnly(backupsOnly).BackupFor(backupFor).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BiddingApi.GetBids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBids`: []Bid
    fmt.Fprintf(os.Stdout, "Response from `BiddingApi.GetBids`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deliveryDate** | **string** | Only bids for the specified delivery date will be listed (ISO 8601 format YYYY-MM-DD). If not specified and also no tender ID is specified, then only bids for today&#39;s delivery day will be listed. | 
 **market** | [**ReserveMarket**](ReserveMarket.md) | Market filter. If not specified and also no tender ID is specified, then only bids for capacity market will be listed. | 
 **tender** | **string** | By specifying the tender ID, only bids of this tender will be listed. For details on the tender id format, see the [reference guide.](/docs/guide#tender-id) | 
 **productType** | [**ProductType**](ProductType.md) | Product type filter. If this query parameter is not set and also no tender ID is specified, bids for all product types are returned. | 
 **productName** | [**ProductName**](ProductName.md) | Product name filter. If this query parameter is not set, bids for all product names are returned. | 
 **state** | **string** |  | 
 **zone** | **string** |  | 
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **type_** | [**BidType**](BidType.md) | By specifying the type, only bids of this type will be listed. | [default to &quot;SIMPLE&quot;]
 **cid** | **string** | Filters bids that have the given complex bid ID. | 
 **gid** | **string** | Filters bids that have the given bid group ID. | 
 **link** | **string** | Filters bids that have the given technical link ID. | 
 **linkedTo** | **string** | Filters bids that have a conditional link to the given bid ID. | 
 **tag** | **string** | Filters bids that have the given tag. | 
 **backupsOnly** | **bool** |  | [default to false]
 **backupFor** | **string** | Filters bids which are a backup for the provider with the given EIC | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 100]

### Return type

[**[]Bid**](Bid.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitBid

> BidSubmissionResponse SubmitBid(ctx).BidSubmission(bidSubmission).Execute()

Submit a simple, complex bid or a complex bid component



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
    bidSubmission := openapiclient.BidSubmission{BidGroup: openapiclient.NewBidGroup(openapiclient.ReserveMarket("CAPACITY"), openapiclient.ProductType("aFRR"), openapiclient.BidType("SIMPLE"), []openapiclient.SimpleBid{*openapiclient.NewSimpleBid(openapiclient.ReserveMarket("CAPACITY"), openapiclient.ProductType("aFRR"), *openapiclient.NewConnectingZone("LFC_AREA", "10YDE-VE-------2"), *openapiclient.NewQuantity(float32(5), "MeasureUnit_example"), []openapiclient.Price{*openapiclient.NewPrice(float32(7.98), "MeasureUnit_example")})})} // BidSubmission |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BiddingApi.SubmitBid(context.Background()).BidSubmission(bidSubmission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BiddingApi.SubmitBid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitBid`: BidSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `BiddingApi.SubmitBid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitBidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bidSubmission** | [**BidSubmission**](BidSubmission.md) |  | 

### Return type

[**BidSubmissionResponse**](BidSubmissionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBid

> Bid UpdateBid(ctx, id).BidModification(bidModification).Execute()

Updates a simple, complex bid or a bid component



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
    id := "id_example" // string | Bid identification
    bidModification := *openapiclient.NewBidModification() // BidModification |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BiddingApi.UpdateBid(context.Background(), id).BidModification(bidModification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BiddingApi.UpdateBid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBid`: Bid
    fmt.Fprintf(os.Stdout, "Response from `BiddingApi.UpdateBid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Bid identification | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bidModification** | [**BidModification**](BidModification.md) |  | 

### Return type

[**Bid**](Bid.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

