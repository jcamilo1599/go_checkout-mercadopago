package models

type PaymentMethodsModel struct {
  DefaultCardID          interface{}            `json:"default_card_id"`
  DefaultPaymentMethodID interface{}            `json:"default_payment_method_id"`
  ExcludedPaymentMethods []ExcludedPaymentModel `json:"excluded_payment_methods"`
  ExcludedPaymentTypes   []ExcludedPaymentModel `json:"excluded_payment_types"`
  Installments           interface{}            `json:"installments"`
  DefaultInstallments    interface{}            `json:"default_installments"`
}
