package service

import "wlt/domain"

type transactionService struct {
	repository domain.Repository
}

func newTransactionService(repository domain.Repository) *transactionService {
	return &transactionService{
		repository: repository,
	}
}

func (t *transactionService) BatchCreate(transactions []domain.Transaction) error {
	return t.repository.Transaction().BatchCreate(transactions)
}
