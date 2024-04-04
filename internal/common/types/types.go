package types

type PaymentRequest struct {
	Amount    float64
	Reference string
	Currency  string
}

type CurrencyCode string
