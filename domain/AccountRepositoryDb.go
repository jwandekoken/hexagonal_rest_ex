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

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	var LastInsertedId int64

	// starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank account transaction
	tx.QueryRow(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values ($1, $2, $3, $4) RETURNING transaction_id`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate).Scan(&LastInsertedId)

	// updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - $1 where account_id = $2`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + $1 where account_id = $2`, t.Amount, t.AccountId)
	}

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// 0 is the zero-value of int vars
	if LastInsertedId == 0 {
		logger.Error("Error while getting the last transaction id")
		return nil, errs.NewUnexpectedError("Could not obtain Transaction id")
	}

	// Getting the latest account information from the accounts table
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(LastInsertedId, 10)

	// updating the transaction struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = $1"
	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}
