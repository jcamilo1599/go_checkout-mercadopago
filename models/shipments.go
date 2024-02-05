package models

type ShipmentsModel struct {
  DefaultShippingMethod interface{}          `json:"default_shipping_method"`
  ReceiverAddress       ReceiverAddressModel `json:"receiver_address"`
}
