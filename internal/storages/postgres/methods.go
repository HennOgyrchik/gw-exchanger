package postgres

import (
	"context"
	"fmt"
	"gw-exchanger/internal/storages"
)

type rateOfCurrency struct {
	currency string
	rate     float32
}

func (p *PSQL) GetExchangeRates(ctx context.Context) (storages.ExchangeRatesResponse, error) {
	const op = "PSQL GetExchangeRates"

	var result storages.ExchangeRatesResponse

	ctxWithTimeout, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	rows, err := p.pool.Query(ctxWithTimeout, "select * from exchange")
	if err != nil {
		return result, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var roc rateOfCurrency

		if err = rows.Scan(&roc.currency, &roc.rate); err != nil {
			return result, fmt.Errorf("%s: %w", op, err)
		}

		result.Rates[roc.currency] = roc.rate
	}

	return result, nil
}

func (p *PSQL) GetExchangeRateForCurrency(ctx context.Context, request storages.CurrencyRequest) (storages.ExchangeRateResponse, error) {
	const op = "PSQL GetExchangeRateForCurrency"

	var result storages.ExchangeRateResponse

	result.FromCurrency = request.FromCurrency
	result.ToCurrency = request.ToCurrency

	ctxWithTimeout, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	rows, err := p.pool.Query(ctxWithTimeout, "select ratio from exchange where currency=$1 or currency=$2", request.FromCurrency, request.FromCurrency)
	if err != nil {
		return result, fmt.Errorf("%s: %w", op, err)
	}

	var rates map[string]float32

	for rows.Next() {
		var roc rateOfCurrency

		if err = rows.Scan(&roc.currency, &roc.rate); err != nil {
			return result, fmt.Errorf("%s: %w", op, err)
		}

		rates[roc.currency] = roc.rate
	}

	rate, err := calculateRate(rates, request.FromCurrency, request.ToCurrency)
	if err != nil {
		return result, fmt.Errorf("%s: %w", op, err)
	}

	result.Rate = rate

	return result, nil
}

func calculateRate(rates map[string]float32, fromCurrency string, toCurrency string) (float32, error) {
	const op = "PSQL calculateRate"

	var (
		from, to float32
		ok       bool
	)

	if from, ok = rates[fromCurrency]; !ok || from == 0 {
		return 0, fmt.Errorf("%s: %s", op, "invalid rates map")
	}

	if to, ok = rates[toCurrency]; !ok || to == 0 {
		return 0, fmt.Errorf("%s: %s", op, "invalid rates map")
	}

	return from / to, nil

}
