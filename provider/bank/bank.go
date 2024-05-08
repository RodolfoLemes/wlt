package bank

import "wlt/domain"

func NewBankProvider(bank string) domain.BankProvider {
	switch bank {
	case "nubank":
		return newNubankProvider()
	default:
		return nil
	}
}
