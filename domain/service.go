package domain

type Service interface {
	Account() AccountService
	Transaction() TransactionService
	Bank() BankService
}
