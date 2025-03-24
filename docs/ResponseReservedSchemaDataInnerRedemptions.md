# ResponseReservedSchemaDataInnerRedemptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedeemSkuCodes** | Pointer to [**[]ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner**](ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner.md) | Contains redeemed SKU information of the voucher (for voucher type is conditional and support sku) | [optional] 
**RedemptionValue** | Pointer to **int32** | Actual redemption value of voucher type &#x3D; conditional | [optional] 

## Methods

### NewResponseReservedSchemaDataInnerRedemptions

`func NewResponseReservedSchemaDataInnerRedemptions() *ResponseReservedSchemaDataInnerRedemptions`

NewResponseReservedSchemaDataInnerRedemptions instantiates a new ResponseReservedSchemaDataInnerRedemptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseReservedSchemaDataInnerRedemptionsWithDefaults

`func NewResponseReservedSchemaDataInnerRedemptionsWithDefaults() *ResponseReservedSchemaDataInnerRedemptions`

NewResponseReservedSchemaDataInnerRedemptionsWithDefaults instantiates a new ResponseReservedSchemaDataInnerRedemptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedeemSkuCodes

`func (o *ResponseReservedSchemaDataInnerRedemptions) GetRedeemSkuCodes() []ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner`

GetRedeemSkuCodes returns the RedeemSkuCodes field if non-nil, zero value otherwise.

### GetRedeemSkuCodesOk

`func (o *ResponseReservedSchemaDataInnerRedemptions) GetRedeemSkuCodesOk() (*[]ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner, bool)`

GetRedeemSkuCodesOk returns a tuple with the RedeemSkuCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemSkuCodes

`func (o *ResponseReservedSchemaDataInnerRedemptions) SetRedeemSkuCodes(v []ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner)`

SetRedeemSkuCodes sets RedeemSkuCodes field to given value.

### HasRedeemSkuCodes

`func (o *ResponseReservedSchemaDataInnerRedemptions) HasRedeemSkuCodes() bool`

HasRedeemSkuCodes returns a boolean if a field has been set.

### GetRedemptionValue

`func (o *ResponseReservedSchemaDataInnerRedemptions) GetRedemptionValue() int32`

GetRedemptionValue returns the RedemptionValue field if non-nil, zero value otherwise.

### GetRedemptionValueOk

`func (o *ResponseReservedSchemaDataInnerRedemptions) GetRedemptionValueOk() (*int32, bool)`

GetRedemptionValueOk returns a tuple with the RedemptionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionValue

`func (o *ResponseReservedSchemaDataInnerRedemptions) SetRedemptionValue(v int32)`

SetRedemptionValue sets RedemptionValue field to given value.

### HasRedemptionValue

`func (o *ResponseReservedSchemaDataInnerRedemptions) HasRedemptionValue() bool`

HasRedemptionValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


