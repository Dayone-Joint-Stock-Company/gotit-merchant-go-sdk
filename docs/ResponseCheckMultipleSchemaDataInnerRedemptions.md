# ResponseCheckMultipleSchemaDataInnerRedemptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedeemSkuCodes** | Pointer to [**[]ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner**](ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner.md) | Contains redeemed SKU information of the voucher (for voucher type is conditional and support sku) | [optional] 
**RedemptionValue** | Pointer to **int64** | Actual redemption value of voucher type &#x3D; conditional | [optional] 
**UsedStore** | Pointer to [**ResponseCheckMultipleSchemaDataInnerRedemptionsUsedStore**](ResponseCheckMultipleSchemaDataInnerRedemptionsUsedStore.md) |  | [optional] 
**UsedDate** | Pointer to **string** | Date voucher marked as used in case the voucher has been redeemed. Format (YYYY-MM-DD HH:MM:SS) | [optional] 
**BillNumber** | Pointer to **string** | Bill number for which voucher used/reserved | [optional] 

## Methods

### NewResponseCheckMultipleSchemaDataInnerRedemptions

`func NewResponseCheckMultipleSchemaDataInnerRedemptions() *ResponseCheckMultipleSchemaDataInnerRedemptions`

NewResponseCheckMultipleSchemaDataInnerRedemptions instantiates a new ResponseCheckMultipleSchemaDataInnerRedemptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCheckMultipleSchemaDataInnerRedemptionsWithDefaults

`func NewResponseCheckMultipleSchemaDataInnerRedemptionsWithDefaults() *ResponseCheckMultipleSchemaDataInnerRedemptions`

NewResponseCheckMultipleSchemaDataInnerRedemptionsWithDefaults instantiates a new ResponseCheckMultipleSchemaDataInnerRedemptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedeemSkuCodes

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) GetRedeemSkuCodes() []ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner`

GetRedeemSkuCodes returns the RedeemSkuCodes field if non-nil, zero value otherwise.

### GetRedeemSkuCodesOk

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) GetRedeemSkuCodesOk() (*[]ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner, bool)`

GetRedeemSkuCodesOk returns a tuple with the RedeemSkuCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemSkuCodes

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) SetRedeemSkuCodes(v []ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner)`

SetRedeemSkuCodes sets RedeemSkuCodes field to given value.

### HasRedeemSkuCodes

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) HasRedeemSkuCodes() bool`

HasRedeemSkuCodes returns a boolean if a field has been set.

### GetRedemptionValue

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) GetRedemptionValue() int64`

GetRedemptionValue returns the RedemptionValue field if non-nil, zero value otherwise.

### GetRedemptionValueOk

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) GetRedemptionValueOk() (*int64, bool)`

GetRedemptionValueOk returns a tuple with the RedemptionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionValue

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) SetRedemptionValue(v int64)`

SetRedemptionValue sets RedemptionValue field to given value.

### HasRedemptionValue

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) HasRedemptionValue() bool`

HasRedemptionValue returns a boolean if a field has been set.

### GetUsedStore

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) GetUsedStore() ResponseCheckMultipleSchemaDataInnerRedemptionsUsedStore`

GetUsedStore returns the UsedStore field if non-nil, zero value otherwise.

### GetUsedStoreOk

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) GetUsedStoreOk() (*ResponseCheckMultipleSchemaDataInnerRedemptionsUsedStore, bool)`

GetUsedStoreOk returns a tuple with the UsedStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStore

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) SetUsedStore(v ResponseCheckMultipleSchemaDataInnerRedemptionsUsedStore)`

SetUsedStore sets UsedStore field to given value.

### HasUsedStore

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) HasUsedStore() bool`

HasUsedStore returns a boolean if a field has been set.

### GetUsedDate

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) GetUsedDate() string`

GetUsedDate returns the UsedDate field if non-nil, zero value otherwise.

### GetUsedDateOk

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) GetUsedDateOk() (*string, bool)`

GetUsedDateOk returns a tuple with the UsedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedDate

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) SetUsedDate(v string)`

SetUsedDate sets UsedDate field to given value.

### HasUsedDate

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) HasUsedDate() bool`

HasUsedDate returns a boolean if a field has been set.

### GetBillNumber

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) GetBillNumber() string`

GetBillNumber returns the BillNumber field if non-nil, zero value otherwise.

### GetBillNumberOk

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) GetBillNumberOk() (*string, bool)`

GetBillNumberOk returns a tuple with the BillNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillNumber

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) SetBillNumber(v string)`

SetBillNumber sets BillNumber field to given value.

### HasBillNumber

`func (o *ResponseCheckMultipleSchemaDataInnerRedemptions) HasBillNumber() bool`

HasBillNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


