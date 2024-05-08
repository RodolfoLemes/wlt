package repository

import (
	"wlt/database"
	"wlt/domain"
)

type repository struct {
	account     domain.AccountRepository
	transaction domain.TransactionRepository

	db database.DB
}

func New(db database.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Account() domain.AccountRepository {
	return r.account
}

func (r *repository) Transaction() domain.TransactionRepository {
	return r.transaction
}
