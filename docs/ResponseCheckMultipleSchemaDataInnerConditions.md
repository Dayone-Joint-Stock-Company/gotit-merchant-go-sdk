# ResponseCheckMultipleSchemaDataInnerConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **string** | Promo start date (YYYY-MM-DD) | [optional] 
**ExcludeSpecificDate** | Pointer to **[]string** | Promo non-effective dates (YYYY-MM-DD) | [optional] 
**ExcludeRecurringDay** | Pointer to **[]string** | Promo non-effective day of week | [optional] 
**RedeemableSkus** | Pointer to **[]string** | List of redeemable SKUs of the voucher code. For voucher type &#x3D; conditional, bill number must contain at least 1 redeemable SKU of the voucher. | [optional] 

## Methods

### NewResponseCheckMultipleSchemaDataInnerConditions

`func NewResponseCheckMultipleSchemaDataInnerConditions() *ResponseCheckMultipleSchemaDataInnerConditions`

NewResponseCheckMultipleSchemaDataInnerConditions instantiates a new ResponseCheckMultipleSchemaDataInnerConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCheckMultipleSchemaDataInnerConditionsWithDefaults

`func NewResponseCheckMultipleSchemaDataInnerConditionsWithDefaults() *ResponseCheckMultipleSchemaDataInnerConditions`

NewResponseCheckMultipleSchemaDataInnerConditionsWithDefaults instantiates a new ResponseCheckMultipleSchemaDataInnerConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExcludeSpecificDate

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetExcludeSpecificDate() []string`

GetExcludeSpecificDate returns the ExcludeSpecificDate field if non-nil, zero value otherwise.

### GetExcludeSpecificDateOk

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetExcludeSpecificDateOk() (*[]string, bool)`

GetExcludeSpecificDateOk returns a tuple with the ExcludeSpecificDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSpecificDate

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) SetExcludeSpecificDate(v []string)`

SetExcludeSpecificDate sets ExcludeSpecificDate field to given value.

### HasExcludeSpecificDate

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) HasExcludeSpecificDate() bool`

HasExcludeSpecificDate returns a boolean if a field has been set.

### GetExcludeRecurringDay

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetExcludeRecurringDay() []string`

GetExcludeRecurringDay returns the ExcludeRecurringDay field if non-nil, zero value otherwise.

### GetExcludeRecurringDayOk

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetExcludeRecurringDayOk() (*[]string, bool)`

GetExcludeRecurringDayOk returns a tuple with the ExcludeRecurringDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeRecurringDay

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) SetExcludeRecurringDay(v []string)`

SetExcludeRecurringDay sets ExcludeRecurringDay field to given value.

### HasExcludeRecurringDay

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) HasExcludeRecurringDay() bool`

HasExcludeRecurringDay returns a boolean if a field has been set.

### GetRedeemableSkus

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetRedeemableSkus() []string`

GetRedeemableSkus returns the RedeemableSkus field if non-nil, zero value otherwise.

### GetRedeemableSkusOk

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetRedeemableSkusOk() (*[]string, bool)`

GetRedeemableSkusOk returns a tuple with the RedeemableSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemableSkus

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) SetRedeemableSkus(v []string)`

SetRedeemableSkus sets RedeemableSkus field to given value.

### HasRedeemableSkus

`func (o *ResponseCheckMultipleSchemaDataInnerConditions) HasRedeemableSkus() bool`

HasRedeemableSkus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


