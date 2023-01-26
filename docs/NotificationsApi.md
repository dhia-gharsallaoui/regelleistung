# \NotificationsApi

All URIs are relative to *https://api.regelleistung.net/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotifications**](NotificationsApi.md#GetNotifications) | **Get** /notifications | Notifications



## GetNotifications

> []Notification GetNotifications(ctx).ProductType(productType).After(after).Limit(limit).Execute()

Notifications



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
    productType := "productType_example" // string | Product type filter for notifications. If not specified, then notifications for all product types are returned. (optional)
    after := time.Now() // time.Time | Oldest publishing time of notfications to be returned. If not specified, then only notifications within the last 24 hours are returned. (optional)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetNotifications(context.Background()).ProductType(productType).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotifications`: []Notification
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productType** | **string** | Product type filter for notifications. If not specified, then notifications for all product types are returned. | 
 **after** | **time.Time** | Oldest publishing time of notfications to be returned. If not specified, then only notifications within the last 24 hours are returned. | 
 **limit** | **int32** | The numbers of items to return | [default to 100]

### Return type

[**[]Notification**](Notification.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

