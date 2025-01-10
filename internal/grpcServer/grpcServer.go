package grpcServer

import (
	"context"
	"fmt"
	pb "github.com/HennOgyrchik/proto-exchange/exchange"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"time"
)

func New(addr string, timeout time.Duration, hndlr ExchangeServiceServer) *Server {
	opts := []grpc.ServerOption{grpc.ConnectionTimeout(timeout)}
	grpcSrv := grpc.NewServer(opts...)

	pb.RegisterExchangeServiceServer(grpcSrv, &handler{service: hndlr})

	return &Server{server: grpcSrv, addr: addr}
}

func (s *Server) Run() error {
	const op = "gRPC run"

	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = s.server.Serve(listener)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Server) Stop() {
	s.server.Stop()
}

func (h *handler) GetExchangeRates(ctx context.Context, e *pb.Empty) (*pb.ExchangeRatesResponse, error) {

	rates, err := h.service.GetExchangeRates(ctx, e)
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
