# ResponseMarkUseMultipleSchemaDataInnerRedemptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedeemSkuCodes** | Pointer to [**[]ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner**](ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner.md) | Contains redeemed SKU information of the voucher (for voucher type is conditional and support sku) | [optional] 
**RedemptionValue** | Pointer to **int64** | Actual redemption value of voucher type &#x3D; conditional | [optional] 
**UsedStore** | Pointer to [**ResponseCheckMultipleSchemaDataInnerRedemptionsUsedStore**](ResponseCheckMultipleSchemaDataInnerRedemptionsUsedStore.md) |  | [optional] 
**UsedDate** | Pointer to **string** | Date voucher marked as used in case the voucher has been redeemed. Format (YYYY-MM-DD HH:MM:SS) | [optional] 

## Methods

### NewResponseMarkUseMultipleSchemaDataInnerRedemptions

`func NewResponseMarkUseMultipleSchemaDataInnerRedemptions() *ResponseMarkUseMultipleSchemaDataInnerRedemptions`

NewResponseMarkUseMultipleSchemaDataInnerRedemptions instantiates a new ResponseMarkUseMultipleSchemaDataInnerRedemptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseMarkUseMultipleSchemaDataInnerRedemptionsWithDefaults

`func NewResponseMarkUseMultipleSchemaDataInnerRedemptionsWithDefaults() *ResponseMarkUseMultipleSchemaDataInnerRedemptions`

NewResponseMarkUseMultipleSchemaDataInnerRedemptionsWithDefaults instantiates a new ResponseMarkUseMultipleSchemaDataInnerRedemptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedeemSkuCodes

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) GetRedeemSkuCodes() []ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner`

GetRedeemSkuCodes returns the RedeemSkuCodes field if non-nil, zero value otherwise.

### GetRedeemSkuCodesOk

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) GetRedeemSkuCodesOk() (*[]ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner, bool)`

GetRedeemSkuCodesOk returns a tuple with the RedeemSkuCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemSkuCodes

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) SetRedeemSkuCodes(v []ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner)`

SetRedeemSkuCodes sets RedeemSkuCodes field to given value.

### HasRedeemSkuCodes

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) HasRedeemSkuCodes() bool`

HasRedeemSkuCodes returns a boolean if a field has been set.

### GetRedemptionValue

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) GetRedemptionValue() int64`

GetRedemptionValue returns the RedemptionValue field if non-nil, zero value otherwise.

### GetRedemptionValueOk

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) GetRedemptionValueOk() (*int64, bool)`

GetRedemptionValueOk returns a tuple with the RedemptionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionValue

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) SetRedemptionValue(v int64)`

SetRedemptionValue sets RedemptionValue field to given value.

### HasRedemptionValue

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) HasRedemptionValue() bool`

HasRedemptionValue returns a boolean if a field has been set.

### GetUsedStore

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) GetUsedStore() ResponseCheckMultipleSchemaDataInnerRedemptionsUsedStore`

GetUsedStore returns the UsedStore field if non-nil, zero value otherwise.

### GetUsedStoreOk

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) GetUsedStoreOk() (*ResponseCheckMultipleSchemaDataInnerRedemptionsUsedStore, bool)`

GetUsedStoreOk returns a tuple with the UsedStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStore

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) SetUsedStore(v ResponseCheckMultipleSchemaDataInnerRedemptionsUsedStore)`

SetUsedStore sets UsedStore field to given value.

### HasUsedStore

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) HasUsedStore() bool`

HasUsedStore returns a boolean if a field has been set.

### GetUsedDate

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) GetUsedDate() string`

GetUsedDate returns the UsedDate field if non-nil, zero value otherwise.

### GetUsedDateOk

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) GetUsedDateOk() (*string, bool)`

GetUsedDateOk returns a tuple with the UsedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedDate

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) SetUsedDate(v string)`

SetUsedDate sets UsedDate field to given value.

### HasUsedDate

`func (o *ResponseMarkUseMultipleSchemaDataInnerRedemptions) HasUsedDate() bool`

HasUsedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


