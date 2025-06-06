/*
Merchant APIs

Technical document APIs for Merchant APIs

API version: 6.0
Contact: duong.vu@gotit.vn
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gotit_merchant_api

import (
	"encoding/json"
)

// checks if the ResponseCheckMultipleSchemaDataInnerConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCheckMultipleSchemaDataInnerConditions{}

// ResponseCheckMultipleSchemaDataInnerConditions Include information involve with voucher type is conditional
type ResponseCheckMultipleSchemaDataInnerConditions struct {
	// Promo start date (YYYY-MM-DD)
	StartDate *string `json:"start_date,omitempty"`
	// Promo non-effective dates (YYYY-MM-DD)
	ExcludeSpecificDate []string `json:"exclude_specific_date,omitempty"`
	// Promo non-effective day of week
	ExcludeRecurringDay []string `json:"exclude_recurring_day,omitempty"`
	// Order value of voucher type = conditional
	OrderValue NullableInt64 `json:"order_value,omitempty"`
	// List of redeemable SKUs of the voucher code. For voucher type = conditional, bill number must contain at least 1 redeemable SKU of the voucher.
	RedeemableSkus []string `json:"redeemable_skus,omitempty"`
}

// NewResponseCheckMultipleSchemaDataInnerConditions instantiates a new ResponseCheckMultipleSchemaDataInnerConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCheckMultipleSchemaDataInnerConditions() *ResponseCheckMultipleSchemaDataInnerConditions {
	this := ResponseCheckMultipleSchemaDataInnerConditions{}
	return &this
}

// NewResponseCheckMultipleSchemaDataInnerConditionsWithDefaults instantiates a new ResponseCheckMultipleSchemaDataInnerConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCheckMultipleSchemaDataInnerConditionsWithDefaults() *ResponseCheckMultipleSchemaDataInnerConditions {
	this := ResponseCheckMultipleSchemaDataInnerConditions{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) SetStartDate(v string) {
	o.StartDate = &v
}

// GetExcludeSpecificDate returns the ExcludeSpecificDate field value if set, zero value otherwise.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetExcludeSpecificDate() []string {
	if o == nil || IsNil(o.ExcludeSpecificDate) {
		var ret []string
		return ret
	}
	return o.ExcludeSpecificDate
}

// GetExcludeSpecificDateOk returns a tuple with the ExcludeSpecificDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetExcludeSpecificDateOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeSpecificDate) {
		return nil, false
	}
	return o.ExcludeSpecificDate, true
}

// HasExcludeSpecificDate returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) HasExcludeSpecificDate() bool {
	if o != nil && !IsNil(o.ExcludeSpecificDate) {
		return true
	}

	return false
}

// SetExcludeSpecificDate gets a reference to the given []string and assigns it to the ExcludeSpecificDate field.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) SetExcludeSpecificDate(v []string) {
	o.ExcludeSpecificDate = v
}

// GetExcludeRecurringDay returns the ExcludeRecurringDay field value if set, zero value otherwise.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetExcludeRecurringDay() []string {
	if o == nil || IsNil(o.ExcludeRecurringDay) {
		var ret []string
		return ret
	}
	return o.ExcludeRecurringDay
}

// GetExcludeRecurringDayOk returns a tuple with the ExcludeRecurringDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetExcludeRecurringDayOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeRecurringDay) {
		return nil, false
	}
	return o.ExcludeRecurringDay, true
}

// HasExcludeRecurringDay returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) HasExcludeRecurringDay() bool {
	if o != nil && !IsNil(o.ExcludeRecurringDay) {
		return true
	}

	return false
}

// SetExcludeRecurringDay gets a reference to the given []string and assigns it to the ExcludeRecurringDay field.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) SetExcludeRecurringDay(v []string) {
	o.ExcludeRecurringDay = v
}

// GetOrderValue returns the OrderValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetOrderValue() int64 {
	if o == nil || IsNil(o.OrderValue.Get()) {
		var ret int64
		return ret
	}
	return *o.OrderValue.Get()
}

// GetOrderValueOk returns a tuple with the OrderValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetOrderValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderValue.Get(), o.OrderValue.IsSet()
}

// HasOrderValue returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) HasOrderValue() bool {
	if o != nil && o.OrderValue.IsSet() {
		return true
	}

	return false
}

// SetOrderValue gets a reference to the given NullableInt64 and assigns it to the OrderValue field.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) SetOrderValue(v int64) {
	o.OrderValue.Set(&v)
}
// SetOrderValueNil sets the value for OrderValue to be an explicit nil
func (o *ResponseCheckMultipleSchemaDataInnerConditions) SetOrderValueNil() {
	o.OrderValue.Set(nil)
}

// UnsetOrderValue ensures that no value is present for OrderValue, not even an explicit nil
func (o *ResponseCheckMultipleSchemaDataInnerConditions) UnsetOrderValue() {
	o.OrderValue.Unset()
}

// GetRedeemableSkus returns the RedeemableSkus field value if set, zero value otherwise.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetRedeemableSkus() []string {
	if o == nil || IsNil(o.RedeemableSkus) {
		var ret []string
		return ret
	}
	return o.RedeemableSkus
}

// GetRedeemableSkusOk returns a tuple with the RedeemableSkus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) GetRedeemableSkusOk() ([]string, bool) {
	if o == nil || IsNil(o.RedeemableSkus) {
		return nil, false
	}
	return o.RedeemableSkus, true
}

// HasRedeemableSkus returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) HasRedeemableSkus() bool {
	if o != nil && !IsNil(o.RedeemableSkus) {
		return true
	}

	return false
}

// SetRedeemableSkus gets a reference to the given []string and assigns it to the RedeemableSkus field.
func (o *ResponseCheckMultipleSchemaDataInnerConditions) SetRedeemableSkus(v []string) {
	o.RedeemableSkus = v
}

func (o ResponseCheckMultipleSchemaDataInnerConditions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCheckMultipleSchemaDataInnerConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.ExcludeSpecificDate) {
		toSerialize["exclude_specific_date"] = o.ExcludeSpecificDate
	}
	if !IsNil(o.ExcludeRecurringDay) {
		toSerialize["exclude_recurring_day"] = o.ExcludeRecurringDay
	}
	if o.OrderValue.IsSet() {
		toSerialize["order_value"] = o.OrderValue.Get()
	}
	if !IsNil(o.RedeemableSkus) {
		toSerialize["redeemable_skus"] = o.RedeemableSkus
	}
	return toSerialize, nil
}

type NullableResponseCheckMultipleSchemaDataInnerConditions struct {
	value *ResponseCheckMultipleSchemaDataInnerConditions
	isSet bool
}

func (v NullableResponseCheckMultipleSchemaDataInnerConditions) Get() *ResponseCheckMultipleSchemaDataInnerConditions {
	return v.value
}

func (v *NullableResponseCheckMultipleSchemaDataInnerConditions) Set(val *ResponseCheckMultipleSchemaDataInnerConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCheckMultipleSchemaDataInnerConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCheckMultipleSchemaDataInnerConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCheckMultipleSchemaDataInnerConditions(val *ResponseCheckMultipleSchemaDataInnerConditions) *NullableResponseCheckMultipleSchemaDataInnerConditions {
	return &NullableResponseCheckMultipleSchemaDataInnerConditions{value: val, isSet: true}
}

func (v NullableResponseCheckMultipleSchemaDataInnerConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCheckMultipleSchemaDataInnerConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


