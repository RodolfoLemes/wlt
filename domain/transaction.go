package domain

import "time"

type TransactionService interface {
	BatchCreate(transactions []Transaction) error
}

type TransactionRepository interface {
	BatchCreate(transactions []Transaction) error
}

type Transaction struct {
	ID            int             `db:"id"`
	AccountID     int             `db:"account_id"`
	Description   string          `db:"description"`
	Amount        string          `db:"amount"`
	Currency      Currency        `db:"currency"`
	Type          TransactionType `db:"type"`
	PaymentMethod PaymentMethod   `db:"payment_method"`
	PaymentAt     time.Time       `db:"payment_at"`
	Category      string          `db:"category"`
}

type TransactionType string

const (
	TransactionTypeIncome  TransactionType = "INCOME"
	TransactionTypeExpense TransactionType = "EXPENSE"
)

type PaymentMethod string

const (
	PaymentMethodReceipt  PaymentMethod = "RECEIPT"
	PaymentMethodPix      PaymentMethod = "PIX"
	PaymentMethodDebit    PaymentMethod = "DEBIT"
	PaymentMethodCredit   PaymentMethod = "CREDIT"
	PaymentMethodTransfer PaymentMethod = "TRANSFER"
)

type Currency string

const (
	CurrencyBRL Currency = "BRL"
	CurrencyUSD Currency = "USD"
)
