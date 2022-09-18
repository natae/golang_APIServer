package service

import (
	"apiserver/server/model"
	"apiserver/server/repository"
)

type AccountService struct {
	repo *repository.AccountRepository
}

func (s *AccountService) GetAccount(nickname string) (*model.Account, error) {
	return s.repo.Select(nickname)
}
