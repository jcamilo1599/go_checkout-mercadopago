package models

type CheckoutPreference struct {
  Items             []ItemModel `json:"items"`
  BackUrls          UrlsModel   `json:"back_urls"`
  ExternalReference string      `json:"external_reference"`
  NotificationURL   string      `json:"notification_url"`
}
