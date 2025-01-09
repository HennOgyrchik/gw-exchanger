package app

import (
	"context"
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

func (a *App) GetExchangeRates(ctx context.Context, empty *pb.Empty) (*pb.ExchangeRatesResponse, error) {
	//TODO
	return nil, nil
}

func (a *App) GetExchangeRateForCurrency(ctx context.Context, req *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error) {
	//TODO
	return nil, nil
}
