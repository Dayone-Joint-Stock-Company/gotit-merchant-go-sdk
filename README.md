# GotIt Merchant SDK

Technical document APIs for Merchant APIs

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
import gotit_merchant_apis "github.com/Dayone-Joint-Stock-Company/gotit-merchant-go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
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
 - [ResponseMarkUseMultipleSchemaDataInnerConditions](docs/ResponseMarkUseMultipleSchemaDataInnerConditions.md)
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
go test test/
```
