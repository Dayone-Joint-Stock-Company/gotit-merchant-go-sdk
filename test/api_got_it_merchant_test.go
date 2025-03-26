package gotit_merchant_apis_test

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    openapiclient "gitlab.gotit.vn/manh.nguyen/gotit-merchant-go-sdk.git"
)

func TestGotItMerchantAPIService(t *testing.T) {
    configuration := openapiclient.NewConfiguration()
    configuration.Servers = openapiclient.ServerConfigurations{
        {
            URL: "https://merchant-api-stg.gotit.vn",
            Description: "Merchant APIs Staging",
        },
    }
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test CheckMultiple", func(t *testing.T) {
        pin := "4205"
        billNumber := "BILL071717127083"
        sku1 := "3002275"
        quantity1 := int32(2)
        price1 := int32(100000)
        sku2 := "3002980"
        quantity2 := int32(3)
        price2 := int32(100000)

        body := openapiclient.RequestCheckMultipleBodySchema{
            Pin:        &pin,
            Codes:      []string{"071717127083"},
            BillNumber: &billNumber,
            SkusInfo: []openapiclient.RequestCheckMultipleBodySchemaSkusInfoInner{
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
        quantity1 := int32(2)
        price1 := int32(100000)
        sku2 := "3002980"
        quantity2 := int32(3)
        price2 := int32(100000)

        body := openapiclient.RequestReservedBodySchema{
            Pin:        &pin,
            Codes:      []string{"071717127083"},
            BillNumber: &billNumber,
            SkusInfo: []openapiclient.RequestCheckMultipleBodySchemaSkusInfoInner{
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

        body := openapiclient.RequestUnReservedBodySchema{
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
        quantity1 := int32(2)
        price1 := int32(100000)
        sku2 := "3002980"
        quantity2 := int32(3)
        price2 := int32(100000)

        body := openapiclient.RequestMarkUseMultipleBodySchema{
            Pin:        &pin,
            Codes:      []string{"071717127083"},
            BillNumber: &billNumber,
            SkusInfo: []openapiclient.RequestCheckMultipleBodySchemaSkusInfoInner{
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
