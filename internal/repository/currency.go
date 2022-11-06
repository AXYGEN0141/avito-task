package repository

import (
	"avito-task/internal"
	"database/sql"
)

type CurrencyRepo struct {
	db *sql.DB
}

func NewCurrencyRepo(db *sql.DB) internal.CurrencyRepositoryInterface {
	return &CurrencyRepo{db: db}
}

// GetCurrencyID shows currency of given ID
func (cr CurrencyRepo) GetCurrencyID(name string) (res int, err error) {
	query := `SELECT currency_id FROM currency WHERE name = $1`
	err = cr.db.QueryRow(query, name).Scan(&res)
	if err != nil {
		return 0, err
	}
	return
}

// Create creates new currency
func (cr CurrencyRepo) Create(name string) (err error) {
	query := `INSERT INTO currency(name) VALUES ($1)`
	_, err = cr.db.Exec(query, name)
	if err != nil {
		return err
	}
	return nil
}
