package service

import "wlt/domain"

type service struct {
	account     domain.AccountService
	transaction domain.TransactionService
	bank        domain.BankService
}

func New(repository domain.Repository) *service {
	account := newAccountService(repository)
	transaction := newTransactionService(repository)
	bank := newBankService(transaction)

	return &service{
		account:     account,
		transaction: transaction,
		bank:        bank,
	}
}

func (s *service) Account() domain.AccountService {
	return s.account
}

func (s *service) Transaction() domain.TransactionService {
	return s.transaction
}

func (s *service) Bank() domain.BankService {
	return s.bank
}
