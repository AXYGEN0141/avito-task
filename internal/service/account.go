package service

import (
	"avito-task/internal"
	"avito-task/internal/model"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

type AccountService struct {
	accountRepo  internal.AccountBalanceRepositoryInterface
	currencyRepo internal.CurrencyRepositoryInterface
}

func NewAccountService(accountRepo internal.AccountBalanceRepositoryInterface, currencyRepo internal.CurrencyRepositoryInterface) internal.AccountBalanceServiceInterface {
	return &AccountService{accountRepo: accountRepo,
		currencyRepo: currencyRepo}
}

// AddCurrencyAccount adds currency to account.
func (as *AccountService) AddCurrencyAccount(uuid, currencyName string) error {
	id, err := as.currencyRepo.GetCurrencyID(currencyName)
	if err != nil {
		return err
	}
	return as.accountRepo.AddCurrencyAccount(uuid, id)
}

// Convert converts currency.
func (as *AccountService) Convert(acc *model.Account) (err error) {

	acc.CurrencyType, err = as.currencyRepo.GetCurrencyID(acc.Currency)
	if err != nil {
		return err
	}
	amount, err := as.CheckBalance(acc.UUID, acc.CurrencyType)
	if err != nil {
		return err
	}
	log.Print(amount, acc.CurrencyType, 212)

	acc.WalletAmount = amount //all balance, with rub - debit

	if err = as.Debit(acc); err != nil {
		return err
	}
	acc.WalletAmount = acc.ConvertedAmount
	acc.Currency = acc.ConvertCurrency

	if err = as.Add(acc); err != nil {
		return err
	}

	return nil
}

// CheckBalance checks balance of account.
func (as *AccountService) CheckBalance(uuid string, currencyType int) (float64, error) {
	return as.accountRepo.CheckBalance(uuid, currencyType)
}

// Add top-ups account with currency.
func (as *AccountService) Add(account *model.Account) (err error) {
	account.CurrencyType, err = as.currencyRepo.GetCurrencyID(account.Currency)
	if err != nil {
		return err
	}
	// if account.WalletAmount > 0 {
	return as.accountRepo.Add(account)
}

// Debit debits currency.
func (as *AccountService) Debit(account *model.Account) (err error) {
	account.CurrencyType, err = as.currencyRepo.GetCurrencyID(account.Currency)
	if err != nil {
		return err
	}
	balanceAmount, err := as.CheckBalance(account.UUID, account.CurrencyType)
	if err != nil {
		return err
	}

	if balanceAmount-account.WalletAmount < 0 {
		return fmt.Errorf("nedostatochno sredstv")
	}
	return as.accountRepo.Debit(account)
}

// NewAccount creates new account.
func (as *AccountService) NewAccount() (int64, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return 0, err
	}
	return as.accountRepo.CreateAccount(id.String())
}

// Transfer sends money from Sender to Receiver.
func (as *AccountService) Transfer(sender model.Account, receiver model.Account) (err error) {
	if err = as.Debit(&sender); err != nil {
		return err
	}
	if err = as.Add(&receiver); err != nil {
		return err
	}
	return nil
}
