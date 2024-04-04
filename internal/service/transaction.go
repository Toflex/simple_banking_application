package service

import "simple_banking_application/internal/repository"

type ITransactionService interface {
}

type transactionService struct {
	transactionRepository repository.ITransactionRepository
}

// NewTransactionService returns a new instance of the user service
func NewTransactionService() ITransactionService {
	return &transactionService{
		transactionRepository: repository.NewTransactionRepo(),
	}
}
