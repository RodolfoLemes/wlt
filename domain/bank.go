package domain

import (
	"os"
)

type BankService interface {
	ListTransactionsFromCSV(filename string, bankName string, accountID int) ([]Transaction, error)
}

type BankProvider interface {
	ListTransactionsFromCSV(r *os.File) ([]Transaction, error)
}
