package repository

import (
	"context"

	"wlt/database"
	"wlt/domain"
)

type transactionRepository struct {
	driver database.Driver
}

func newTransactionRepository(driver database.Driver) *transactionRepository {
	return &transactionRepository{
		driver: driver,
	}
}

func (r *transactionRepository) BatchCreate(transactions []domain.Transaction) error {
	_, err := r.driver.NamedExec(
		context.Background(),
		`INSERT INTO transaction 
		(account_id, description, amount, currency, type, payment_method, payment_at, category) 
		VALUES (:account_id, :description, :amount, :currency, :type, :payment_method, :payment_at, :categorymi)`,
		transactions,
	)

	return err
}
