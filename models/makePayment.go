package models

import "encoding/json"

func UnmarshalMakePayment(data []byte) (MakePayment, error) {
  var r MakePayment
  err := json.Unmarshal(data, &r)
  return r, err
}

func (r *MakePayment) Marshal() ([]byte, error) {
  return json.Marshal(r)
}

type MakePayment struct {
  AdditionalInfo     string              `json:"additional_info"`
  AutoReturn         string              `json:"auto_return"`
  BackUrls           UrlsModel           `json:"back_urls"`
  BinaryMode         bool                `json:"binary_mode"`
  ClientID           string              `json:"client_id"`
  CollectorID        int64               `json:"collector_id"`
  CouponCode         interface{}         `json:"coupon_code"`
  CouponLabels       interface{}         `json:"coupon_labels"`
  DateCreated        string              `json:"date_created"`
  DateOfExpiration   interface{}         `json:"date_of_expiration"`
  ExpirationDateFrom interface{}         `json:"expiration_date_from"`
  ExpirationDateTo   interface{}         `json:"expiration_date_to"`
  Expires            bool                `json:"expires"`
  ExternalReference  string              `json:"external_reference"`
  ID                 string              `json:"id"`
  InitPoint          string              `json:"init_point"`
  InternalMetadata   interface{}         `json:"internal_metadata"`
  Items              []ItemModel         `json:"items"`
  Marketplace        string              `json:"marketplace"`
  MarketplaceFee     int64               `json:"marketplace_fee"`
  Metadata           interface{}         `json:"metadata"`
  NotificationURL    string              `json:"notification_url"`
  OperationType      string              `json:"operation_type"`
  Payer              PayerModel          `json:"payer"`
  PaymentMethods     PaymentMethodsModel `json:"payment_methods"`
  ProcessingModes    interface{}         `json:"processing_modes"`
  ProductID          interface{}         `json:"product_id"`
  RedirectUrls       UrlsModel           `json:"redirect_urls"`
  SandboxInitPoint   string              `json:"sandbox_init_point"`
  SiteID             string              `json:"site_id"`
  Shipments          ShipmentsModel      `json:"shipments"`
  TotalAmount        interface{}         `json:"total_amount"`
  LastUpdated        interface{}         `json:"last_updated"`
}
