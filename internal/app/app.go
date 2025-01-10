package app

import (
	"context"
	"fmt"
	pb "github.com/HennOgyrchik/proto-exchange/exchange"
	"gw-exchanger/internal/storages"
	"gw-exchanger/pkg/logs"
)

func New(logger *logs.Log, storage storages.Storage) *App {
	return &App{
		log:     logger,
		storage: storage,
	}
}

func (a *App) GetExchangeRates(ctx context.Context, _ *pb.Empty) (*pb.ExchangeRatesResponse, error) {
	const op = "App GetExchangeRates"

	result, err := a.storage.GetExchangeRates(ctx)
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
	}

	return &pb.ExchangeRatesResponse{Rates: result.Rates}, err
}

func (a *App) GetExchangeRateForCurrency(ctx context.Context, req *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error) {
	const op = "App GetExchangeRateForCurrency"

	result, err := a.storage.GetExchangeRateForCurrency(ctx, storages.CurrencyRequest{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
	})
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
	}

	return &pb.ExchangeRateResponse{
		FromCurrency: result.FromCurrency,
		ToCurrency:   result.ToCurrency,
		Rate:         result.Rate,
	}, err
}
