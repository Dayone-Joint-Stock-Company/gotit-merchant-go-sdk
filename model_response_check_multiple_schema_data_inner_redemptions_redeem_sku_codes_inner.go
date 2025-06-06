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

// checks if the ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner{}

// ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner struct for ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner
type ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner struct {
	// SKU code is redeemed for voucher
	Sku *string `json:"sku,omitempty"`
	// SKU quantity is redeemed for voucher
	Quantity *int64 `json:"quantity,omitempty"`
	// Selling price of SKU in bill.
	Price *int64 `json:"price,omitempty"`
}

// NewResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner instantiates a new ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner() *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner {
	this := ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner{}
	return &this
}

// NewResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInnerWithDefaults instantiates a new ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInnerWithDefaults() *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner {
	this := ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) SetSku(v string) {
	o.Sku = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) GetPrice() int64 {
	if o == nil || IsNil(o.Price) {
		var ret int64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) GetPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given int64 and assigns it to the Price field.
func (o *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) SetPrice(v int64) {
	o.Price = &v
}

func (o ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner struct {
	value *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner
	isSet bool
}

func (v NullableResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) Get() *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner {
	return v.value
}

func (v *NullableResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) Set(val *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner(val *ResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) *NullableResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner {
	return &NullableResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner{value: val, isSet: true}
}

func (v NullableResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCheckMultipleSchemaDataInnerRedemptionsRedeemSkuCodesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


