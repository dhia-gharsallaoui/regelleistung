# \HealthApi

All URIs are relative to *https://api.regelleistung.net/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiHealthInformations**](HealthApi.md#GetApiHealthInformations) | **Get** /health | API health endpoint.



## GetApiHealthInformations

> Health GetApiHealthInformations(ctx).Execute()

API health endpoint.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HealthApi.GetApiHealthInformations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.GetApiHealthInformations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiHealthInformations`: Health
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.GetApiHealthInformations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiHealthInformationsRequest struct via the builder pattern


### Return type

[**Health**](Health.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

