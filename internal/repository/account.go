package repository

import (
	"avito-task/internal"
	"avito-task/internal/model"
	"database/sql"
	"log"
)

type AccountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) internal.AccountBalanceRepositoryInterface {
	return &AccountRepo{db: db}
}

// AddCurrencyAccount creates new currency account.
func (ar *AccountRepo) AddCurrencyAccount(uuid string, currencyID int) (err error) {
	query := `INSERT INTO account(uuid, balance, currency_type) VALUES ($1, $2, $3) RETURNING account_id`
	_, err = ar.db.Exec(query, uuid, 0, currencyID)
	if err != nil {
		return err
	}
	return nil
}

// CreateAccount creates new account.
func (ar *AccountRepo) CreateAccount(uuid string) (res int64, err error) {

	query := `INSERT INTO account(uuid, balance, currency_type) VALUES ($1, $2, $3) RETURNING account_id`
	err = ar.db.QueryRow(query, uuid, 0, 1).Scan(&res)
	if err != nil {
		return 0, err
	}
	return res, nil
}

// Add top-ups balance of account.
func (ar *AccountRepo) Add(account *model.Account) error {
	log.Println("update add balance repo", account.WalletAmount, account.CurrencyType)

	query := "UPDATE account SET balance = balance + $1 WHERE uuid = $2 AND currency_type=$3"
	_, err := ar.db.Exec(query, account.WalletAmount, account.UUID, account.CurrencyType)
	if err != nil {
		return err
	}
	return nil
}

// CheckBalance checks balance of account.
func (ar *AccountRepo) CheckBalance(uuid string, currencyType int) (amount float64, err error) {
	query := `SELECT balance FROM account WHERE uuid = $1 AND currency_type=$2`
	err = ar.db.QueryRow(query, uuid, currencyType).Scan(&amount)
	if err != nil {
		return 0, err
	}
	return
}

// Debit debits currency.
func (ar *AccountRepo) Debit(account *model.Account) error {
	query := "UPDATE account SET balance = balance - $1 WHERE uuid = $2 AND currency_type=$3"
	_, err := ar.db.Exec(query, account.WalletAmount, account.UUID, account.CurrencyType)
	if err != nil {
		return err
	}
	return nil
}
