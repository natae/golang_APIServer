package service

import (
	"apiserver/server/model"
	"apiserver/server/repository"
	"database/sql"
)

type AccountService struct {
	repo *repository.AccountRepository
}

func NewAccountService(db *sql.DB) *AccountService {
	return &AccountService{repo: repository.NewAccountRepository(db)}
}

func (s *AccountService) GetAccount(nickname string) (*model.Account, error) {
	return s.repo.Select(nickname)
}

func (s *AccountService) CreateAccount(nickname string) (int64, error) {
	return s.repo.Insert(nickname)
}
