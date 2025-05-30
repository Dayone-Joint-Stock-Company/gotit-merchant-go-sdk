# GotIt Merchant SDK

SDK Technical document for Merchant APIs

## Requirements
Go 1.18 or later

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import gotit_merchant_apis "github.com/Dayone-Joint-Stock-Company/gotit-merchant-go-sdk/v1.0.0"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Basic Usage

```go
package main

import (
	"context"
	"fmt"
	"os"
	merchantApiClient "github.com/Dayone-Joint-Stock-Company/gotit-merchant-go-sdk"
)

func main() {
    configuration := merchantApiClient.NewConfiguration()
    configuration.Servers = merchantApiClient.ServerConfigurations{
        {
            URL: "https://openapi-stg.gotit.vn",
            Description: "Merchant APIs Staging",
        },
    }
    apiClient := merchantApiClient.NewAPIClient(configuration)

    pin := "4205"
    billNumber := "BILL071717127083"
    sku1 := "3002275"
    quantity1 := int32(2)
    price1 := int32(100000)
    sku2 := "3002980"
    quantity2 := int32(3)
    price2 := int32(100000)

    body := merchantApiClient.RequestCheckMultipleBodySchema{
        Pin:        &pin,
        Codes:      []string{"071717127083"},
        BillNumber: &billNumber,
        SkusInfo: []merchantApiClient.RequestCheckMultipleBodySchemaSkusInfoInner{
            {Sku: &sku1, Quantity: &quantity1, Price: &price1},
            {Sku: &sku2, Quantity: &quantity2, Price: &price2},
        },
    }

    resp, r, err := apiClient.GotItMerchantAPI.CheckMultiple(context.Background()).RequestCheckMultipleBodySchema(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GotItMerchantAPI.CheckMultiple``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckMultiple`: ResponseCheckMultipleSchema
    fmt.Fprintf(os.Stdout, "Response from `GotItMerchantAPI.CheckMultiple`: %v\n", resp)
}
```

## Documentation for API Endpoints

All URIs are relative to *https://openapi-stg.gotit.vn*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*GotItMerchantAPI* | [**CheckMultiple**](docs/GotItMerchantAPI.md#checkmultiple) | **Post** /merchant/v6.0/checkmultiple | Check multiple vouchers are valid or not
*GotItMerchantAPI* | [**Reserved**](docs/GotItMerchantAPI.md#reserved) | **Post** /merchant/v6.0/reserved | Reserved multiple vouchers for a fixed bill number.
*GotItMerchantAPI* | [**Unreserved**](docs/GotItMerchantAPI.md#unreserved) | **Post** /merchant/v6.0/unreserved | Reserved multiple vouchers for a fixed bill number.
*GotItMerchantAPI* | [**UseMultiple**](docs/GotItMerchantAPI.md#usemultiple) | **Post** /merchant/v6.0/usemultiple | Reserved multiple vouchers for a fixed bill number.


## Documentation For Models

 - [RequestCheckMultipleBodySchema](docs/RequestCheckMultipleBodySchema.md)
 - [RequestCheckMultipleBodySchemaSkusInfoInner](docs/RequestCheckMultipleBodySchemaSkusInfoInner.md)
 - [RequestMarkUseMultipleBodySchema](docs/RequestMarkUseMultipleBodySchema.md)
 - [RequestReservedBodySchema](docs/RequestReservedBodySchema.md)
 - [RequestUnReservedBodySchema](docs/RequestUnReservedBodySchema.md)
 - [ResponseCheckMultipleSchema](docs/ResponseCheckMultipleSchema.md)
 - [ResponseCheckMultipleSchemaDataInner](docs/ResponseCheckMultipleSchemaDataInner.md)
 - [ResponseCheckMultipleSchemaDataInnerConditions](docs/ResponseCheckMultipleSchemaDataInnerConditions.md)
 - [ResponseCheckMultipleSchemaDataInnerRedemptions](docs/ResponseCheckMultipleSchemaDataInnerRedemptions.md)
 - [ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner](docs/ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner.md)
 - [ResponseCheckMultipleSchemaDataInnerRedemptionsUsedStore](docs/ResponseCheckMultipleSchemaDataInnerRedemptionsUsedStore.md)
 - [ResponseMarkUseMultipleSchema](docs/ResponseMarkUseMultipleSchema.md)
 - [ResponseMarkUseMultipleSchemaDataInner](docs/ResponseMarkUseMultipleSchemaDataInner.md)
 - [ResponseMarkUseMultipleSchemaDataInnerRedemptions](docs/ResponseMarkUseMultipleSchemaDataInnerRedemptions.md)
 - [ResponseReservedSchema](docs/ResponseReservedSchema.md)
 - [ResponseReservedSchemaDataInner](docs/ResponseReservedSchemaDataInner.md)
 - [ResponseReservedSchemaDataInnerRedemptions](docs/ResponseReservedSchemaDataInnerRedemptions.md)
 - [ResponseReservedSchemaUsedStore](docs/ResponseReservedSchemaUsedStore.md)
 - [ResponseUnReservedSchema](docs/ResponseUnReservedSchema.md)
 - [ResponseUnReservedSchemaDataInner](docs/ResponseUnReservedSchemaDataInner.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Tests

To run the tests, use:

```bash
go test ./test/
```
