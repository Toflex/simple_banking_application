package models

import "simple_banking_application/cmd/internal/common/types"

// Account keeps track of users balance
type Account struct {
	Base

	UserID   string
	Currency types.CurrencyCode `gorm:"default:'NGN'"`
	Balance  float64
	User     User
}
