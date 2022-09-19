package repository

import (
	"apiserver/server/model"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (r *AccountRepository) Select(nickname string) (*model.Account, error) {
	var account = &model.Account{}
	query := "select * from Account where Nickname = ?"

	err := r.db.QueryRow(query, nickname).Scan(&account.Id, &account.Nickname)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r *AccountRepository) Insert(nickname string) (int64, error) {
	query := "insert into Account (Nickname) values (?)"

	result, err := r.db.Exec(query, nickname)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}
