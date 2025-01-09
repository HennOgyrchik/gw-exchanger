package grpcServer

import (
	"context"
	pb "github.com/HennOgyrchik/proto-exchange/exchange"
	"google.golang.org/grpc"
)

type ExchangeServiceServer interface {
	GetExchangeRates(context.Context, *pb.Empty) (*pb.ExchangeRatesResponse, error)
	GetExchangeRateForCurrency(context.Context, *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error)
}

type handler struct {
	pb.UnimplementedExchangeServiceServer
	service ExchangeServiceServer
}

type Server struct {
	addr   string
	server *grpc.Server
}
