package storages

import "context"

type Storage interface {
	GetExchangeRates(ctx context.Context) (ExchangeRatesResponse, error)
	GetExchangeRateForCurrency(ctx context.Context, request CurrencyRequest) (ExchangeRateResponse, error)
}
