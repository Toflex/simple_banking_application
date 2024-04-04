package repository

import (
	"gorm.io/gorm"
	"simple_banking_application/internal/database"
)

type ITransactionRepository interface {
	WithTx(tx *gorm.DB) ITransactionRepository
}

type transactionRepo struct {
	db *gorm.DB
}

func (r *transactionRepo) WithTx(tx *gorm.DB) ITransactionRepository {
	return &transactionRepo{db: tx}
}

// NewTransactionRepo will instantiate Transaction Repository
func NewTransactionRepo() ITransactionRepository {
	return &transactionRepo{db: database.DB()}
}
