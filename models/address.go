package models

type AddressModel struct {
  ZipCode      string      `json:"zip_code"`
  StreetName   string      `json:"street_name"`
  StreetNumber interface{} `json:"street_number"`
}
