package models

import "simple_banking_application/cmd/internal/common/types"

const (
	NGN types.CurrencyCode = "NGN"
)

type Transaction struct {
	Base

	Currency  types.CurrencyCode
	Amount    float64
	Reason    float64
	Reference string
}
