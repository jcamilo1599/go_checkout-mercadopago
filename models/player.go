package models

type PayerModel struct {
  Phone          PhoneModel     `json:"phone"`
  Address        AddressModel   `json:"address"`
  Email          string         `json:"email"`
  Identification Identification `json:"identification"`
  Name           string         `json:"name"`
  Surname        string         `json:"surname"`
  DateCreated    interface{}    `json:"date_created"`
  LastPurchase   interface{}    `json:"last_purchase"`
}
