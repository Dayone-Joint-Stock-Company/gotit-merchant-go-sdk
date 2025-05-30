package gotit_merchant_apis_test

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    merchantApiClient "github.com/Dayone-Joint-Stock-Company/gotit-merchant-go-sdk"
)

func TestGotItMerchantAPIService(t *testing.T) {
    configuration := merchantApiClient.NewConfiguration()
    configuration.Servers = merchantApiClient.ServerConfigurations{
        {
            URL: "https://openapi-stg.gotit.vn",
            Description: "Merchant APIs Staging",
        },
    }
    apiClient := merchantApiClient.NewAPIClient(configuration)

    t.Run("Test CheckMultiple", func(t *testing.T) {
        pin := "4205"
        billNumber := "BILL071717127083"
        sku1 := "3002275"
        quantity1 := int64(2)
        price1 := int64(100000)
        sku2 := "3002980"
        quantity2 := int64(3)
        price2 := int64(100000)

        body := merchantApiClient.RequestCheckMultipleBodySchema{
            Pin:        &pin,
            Codes:      []string{"071717127083"},
            BillNumber: &billNumber,
            SkusInfo: []merchantApiClient.RequestCheckMultipleBodySchemaSkusInfoInner{
                {Sku: &sku1, Quantity: &quantity1, Price: &price1},
                {Sku: &sku2, Quantity: &quantity2, Price: &price2},
            },
        }

        resp, httpRes, err := apiClient.GotItMerchantAPI.CheckMultiple(context.Background()).RequestCheckMultipleBodySchema(body).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        assert.False(t, *resp.Success)
    })

    t.Run("Test Reserved", func(t *testing.T) {
        pin := "4205"
        billNumber := "BILL071717127083"
        sku1 := "3002275"
        quantity1 := int64(2)
        price1 := int64(100000)
        sku2 := "3002980"
        quantity2 := int64(3)
        price2 := int64(100000)

        body := merchantApiClient.RequestReservedBodySchema{
            Pin:        &pin,
            Codes:      []string{"071717127083"},
            BillNumber: &billNumber,
            SkusInfo: []merchantApiClient.RequestCheckMultipleBodySchemaSkusInfoInner{
                {Sku: &sku1, Quantity: &quantity1, Price: &price1},
                {Sku: &sku2, Quantity: &quantity2, Price: &price2},
            },
        }

        resp, httpRes, err := apiClient.GotItMerchantAPI.Reserved(context.Background()).RequestReservedBodySchema(body).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        assert.False(t, *resp.Success)
    })

    t.Run("Test Unreserved", func(t *testing.T) {
        pin := "4205"
        billNumber := "BILL071717127083"

        body := merchantApiClient.RequestUnReservedBodySchema{
            Pin:        &pin,
            Codes:      []string{"071717127083"},
            BillNumber: &billNumber,
        }

        resp, httpRes, err := apiClient.GotItMerchantAPI.Unreserved(context.Background()).RequestUnReservedBodySchema(body).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        assert.False(t, *resp.Success)
    })

    t.Run("Test UseMultiple", func(t *testing.T) {
        pin := "4205"
        billNumber := "BILL071717127083"
        sku1 := "3002275"
        quantity1 := int64(2)
        price1 := int64(100000)
        sku2 := "3002980"
        quantity2 := int64(3)
        price2 := int64(100000)

        body := merchantApiClient.RequestMarkUseMultipleBodySchema{
            Pin:        &pin,
            Codes:      []string{"071717127083"},
            BillNumber: &billNumber,
            SkusInfo: []merchantApiClient.RequestCheckMultipleBodySchemaSkusInfoInner{
                {Sku: &sku1, Quantity: &quantity1, Price: &price1},
                {Sku: &sku2, Quantity: &quantity2, Price: &price2},
            },
        }

        resp, httpRes, err := apiClient.GotItMerchantAPI.UseMultiple(context.Background()).RequestMarkUseMultipleBodySchema(body).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        assert.False(t, *resp.Success)
    })
}
