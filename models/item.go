package models

type ItemModel struct {
  ID          string `json:"id"`
  CategoryID  string `json:"category_id"`
  CurrencyID  string `json:"currency_id"`
  Description string `json:"description"`
  Title       string `json:"title"`
  Quantity    int64  `json:"quantity"`
  UnitPrice   int64  `json:"unit_price"`
}
