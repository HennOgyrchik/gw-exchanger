package grpcServer

import (
	"fmt"
	pb "github.com/HennOgyrchik/proto-exchange/exchange"
	"google.golang.org/grpc"
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
	s.server.GracefulStop()
}
