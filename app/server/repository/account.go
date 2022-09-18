package repository

import (
	"apiserver/server/model"
)

type AccountRepository struct {
	//TODO:  dbAccessor...
}

func (r *AccountRepository) Select(nickname string) (*model.Account, error) {
	panic("Sorry")
	//return nil, errors.New("Not implement")
}
