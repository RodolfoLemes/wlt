package domain

type Repository interface {
	Account() AccountRepository
	Transaction() TransactionRepository
}
