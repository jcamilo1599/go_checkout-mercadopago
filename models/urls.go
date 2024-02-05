package models

type UrlsModel struct {
  Failure string `json:"failure"`
  Pending string `json:"pending"`
  Success string `json:"success"`
}
