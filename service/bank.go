package service

import (
	"os"

	"wlt/domain"
	"wlt/provider/bank"
)

type bankService struct {
	transactionService domain.TransactionService
}

func newBankService(
	transactionService domain.TransactionService,
) *bankService {
	return &bankService{
		transactionService: transactionService,
	}
}

func (b *bankService) ListTransactionsFromCSV(filename string, bankName string, accountID int) ([]domain.Transaction, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path + filename)
	if err != nil {
		return nil, err
	}

	bankProvider := bank.NewBankProvider(bankName)

	transactions, err := bankProvider.ListTransactionsFromCSV(file)
	if err != nil {
		return nil, err
	}

	for i := range transactions {
		transactions[i].AccountID = accountID
	}

	err = b.transactionService.BatchCreate(transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
