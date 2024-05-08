package repository

import (
	"context"

	"wlt/database"
	"wlt/domain"
)

type accountRepository struct {
	driver database.Driver
}

func newAccountRepository(driver database.Driver) *accountRepository {
	return &accountRepository{
		driver: driver,
	}
}

func (a *accountRepository) List() ([]domain.Account, error) {
	var accounts []domain.Account
	err := a.driver.Select(context.Background(), &accounts, "SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (a *accountRepository) Create(account domain.Account) error {
	_, err := a.driver.NamedExec(
		context.Background(),
		"INSERT INTO account (name, bank_name, owner_name) VALUES (:name, :bank_name, :owner_name)",
		account,
	)

	return err
}
