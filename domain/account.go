package domain

type AccountService interface {
	List() ([]Account, error)
	Create(account Account) error
}

type AccountRepository interface {
	List() ([]Account, error)
	Create(account Account) error
}

type Account struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	BankName  string `db:"bank_name"`
	OwnerName string `db:"owner_name"`
}
