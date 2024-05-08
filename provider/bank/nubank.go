package bank

import (
	"log"
	"os"

	"wlt/domain"
	"wlt/internal/timehelper"

	"github.com/gocarina/gocsv"
)

type nubankProvider struct{}

func newNubankProvider() *nubankProvider {
	return &nubankProvider{}
}

type nubankCreditCSV struct {
	Date     string `csv:"date"`
	Category string `csv:"category"`
	Title    string `csv:"title"`
	Amount   string `csv:"amount"`
}

func (n *nubankProvider) ListTransactionsFromCSV(r *os.File) ([]domain.Transaction, error) {
	var nus []nubankCreditCSV

	err := gocsv.UnmarshalFile(r, &nus)
	if err != nil {
		return nil, err
	}

	transactions := []domain.Transaction{}
	for _, nu := range nus {
		paymentAt, err := timehelper.ParseDate(nu.Date)
		if err != nil {
			log.Println(err)
		}

		transactions = append(transactions, domain.Transaction{
			Description:   nu.Title,
			Amount:        nu.Amount,
			Currency:      domain.CurrencyBRL,
			Type:          domain.TransactionTypeExpense,
			PaymentMethod: domain.PaymentMethodCredit,
			PaymentAt:     paymentAt,
			Category:      nu.Category,
		})
	}

	return nil, nil
}
