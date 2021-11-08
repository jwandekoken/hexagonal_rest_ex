package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/jwandekoken/golang_rest-server/errs"
	"github.com/jwandekoken/golang_rest-server/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	var id int64

	err := d.client.Get(&id, "insert into accounts (customer_id, opening_date, account_type, amount, status) values ($1, $2, $3, $4, $5) returning account_id", a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating an Account: " + err.Error())
		return nil, errs.NewUnexpectedError("Error while creating new account")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}
