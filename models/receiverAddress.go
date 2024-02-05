package models

type ReceiverAddressModel struct {
  ZipCode      string      `json:"zip_code"`
  StreetName   string      `json:"street_name"`
  StreetNumber interface{} `json:"street_number"`
  Floor        string      `json:"floor"`
  Apartment    string      `json:"apartment"`
  CityName     interface{} `json:"city_name"`
  StateName    interface{} `json:"state_name"`
  CountryName  interface{} `json:"country_name"`
}
