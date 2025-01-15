package grpcServer

import (
	"context"
	pb "github.com/HennOgyrchik/proto-exchange/exchange"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) GetExchangeRates(ctx context.Context, _ *pb.Empty) (*pb.ExchangeRatesResponse, error) {

	rates, err := h.service.GetExchangeRates(ctx, &pb.Empty{})
	if err != nil {
		err = status.Errorf(codes.Internal, "Internal error")
	}

	return rates, err
}

func (h *handler) GetExchangeRateForCurrency(ctx context.Context, req *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error) {

	currency, err := h.service.GetExchangeRateForCurrency(ctx, req)
	if err != nil {
		err = status.Errorf(codes.Internal, "Internal error")
	}
	return currency, err
}
