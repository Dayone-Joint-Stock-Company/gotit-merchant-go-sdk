# \GotItMerchantAPI

All URIs are relative to *https://openapi-stg.gotit.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckMultiple**](GotItMerchantAPI.md#CheckMultiple) | **Post** /merchant/v6.0/checkmultiple | Check multiple vouchers are valid or not
[**Reserved**](GotItMerchantAPI.md#Reserved) | **Post** /merchant/v6.0/reserved | Reserved multiple vouchers for a fixed bill number.
[**Unreserved**](GotItMerchantAPI.md#Unreserved) | **Post** /merchant/v6.0/unreserved | Reserved multiple vouchers for a fixed bill number.
[**UseMultiple**](GotItMerchantAPI.md#UseMultiple) | **Post** /merchant/v6.0/usemultiple | Reserved multiple vouchers for a fixed bill number.



## CheckMultiple

> ResponseCheckMultipleSchema CheckMultiple(ctx).RequestCheckMultipleBodySchema(requestCheckMultipleBodySchema).Execute()

Check multiple vouchers are valid or not



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Dayone-Joint-Stock-Company/gotit-merchant-go-sdk"
)

func main() {
	requestCheckMultipleBodySchema := *openapiclient.NewRequestCheckMultipleBodySchema() // RequestCheckMultipleBodySchema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GotItMerchantAPI.CheckMultiple(context.Background()).RequestCheckMultipleBodySchema(requestCheckMultipleBodySchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GotItMerchantAPI.CheckMultiple``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckMultiple`: ResponseCheckMultipleSchema
	fmt.Fprintf(os.Stdout, "Response from `GotItMerchantAPI.CheckMultiple`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckMultipleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCheckMultipleBodySchema** | [**RequestCheckMultipleBodySchema**](RequestCheckMultipleBodySchema.md) |  | 

### Return type

[**ResponseCheckMultipleSchema**](ResponseCheckMultipleSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Reserved

> ResponseReservedSchema Reserved(ctx).RequestReservedBodySchema(requestReservedBodySchema).Execute()

Reserved multiple vouchers for a fixed bill number.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Dayone-Joint-Stock-Company/gotit-merchant-go-sdk"
)

func main() {
	requestReservedBodySchema := *openapiclient.NewRequestReservedBodySchema() // RequestReservedBodySchema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GotItMerchantAPI.Reserved(context.Background()).RequestReservedBodySchema(requestReservedBodySchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GotItMerchantAPI.Reserved``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Reserved`: ResponseReservedSchema
	fmt.Fprintf(os.Stdout, "Response from `GotItMerchantAPI.Reserved`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReservedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestReservedBodySchema** | [**RequestReservedBodySchema**](RequestReservedBodySchema.md) |  | 

### Return type

[**ResponseReservedSchema**](ResponseReservedSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unreserved

> ResponseUnReservedSchema Unreserved(ctx).RequestUnReservedBodySchema(requestUnReservedBodySchema).Execute()

Reserved multiple vouchers for a fixed bill number.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Dayone-Joint-Stock-Company/gotit-merchant-go-sdk"
)

func main() {
	requestUnReservedBodySchema := *openapiclient.NewRequestUnReservedBodySchema() // RequestUnReservedBodySchema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GotItMerchantAPI.Unreserved(context.Background()).RequestUnReservedBodySchema(requestUnReservedBodySchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GotItMerchantAPI.Unreserved``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Unreserved`: ResponseUnReservedSchema
	fmt.Fprintf(os.Stdout, "Response from `GotItMerchantAPI.Unreserved`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnreservedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestUnReservedBodySchema** | [**RequestUnReservedBodySchema**](RequestUnReservedBodySchema.md) |  | 

### Return type

[**ResponseUnReservedSchema**](ResponseUnReservedSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UseMultiple

> ResponseMarkUseMultipleSchema UseMultiple(ctx).RequestMarkUseMultipleBodySchema(requestMarkUseMultipleBodySchema).Execute()

Reserved multiple vouchers for a fixed bill number.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Dayone-Joint-Stock-Company/gotit-merchant-go-sdk"
)

func main() {
	requestMarkUseMultipleBodySchema := *openapiclient.NewRequestMarkUseMultipleBodySchema() // RequestMarkUseMultipleBodySchema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GotItMerchantAPI.UseMultiple(context.Background()).RequestMarkUseMultipleBodySchema(requestMarkUseMultipleBodySchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GotItMerchantAPI.UseMultiple``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UseMultiple`: ResponseMarkUseMultipleSchema
	fmt.Fprintf(os.Stdout, "Response from `GotItMerchantAPI.UseMultiple`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUseMultipleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestMarkUseMultipleBodySchema** | [**RequestMarkUseMultipleBodySchema**](RequestMarkUseMultipleBodySchema.md) |  | 

### Return type

[**ResponseMarkUseMultipleSchema**](ResponseMarkUseMultipleSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

