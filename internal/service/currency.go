package service

import "avito-task/internal"

type CurrencyService struct {
	repo internal.CurrencyRepositoryInterface
}

// Create creates new currency.
func (cr CurrencyService) Create(name string) error {
	return cr.repo.Create(name)
}

// GetCurrencyID gets currency ID.
func (cr CurrencyService) GetCurrencyID(name string) (int, error) {
	return cr.repo.GetCurrencyID(name)
}

func NewCurrencyService(repo internal.CurrencyRepositoryInterface) internal.CurrencyServiceInterface {
	return &CurrencyService{repo: repo}
}
