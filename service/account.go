package service

import "wlt/domain"

type accountService struct {
	repository domain.Repository
}

func newAccountService(repository domain.Repository) *accountService {
	return &accountService{
		repository: repository,
	}
}

func (a *accountService) List() ([]domain.Account, error) {
	return a.repository.Account().List()
}

func (a *accountService) Create(account domain.Account) error {
	return a.repository.Account().Create(account)
}
