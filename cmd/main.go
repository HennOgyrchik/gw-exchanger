package main

import (
	"context"
	"gw-exchanger/internal/app"
	"gw-exchanger/internal/config"
	"gw-exchanger/internal/grpcServer"
	"gw-exchanger/internal/storages/postgres"
	"gw-exchanger/pkg/logs"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer stop()

	logger := logs.New(os.Stdout)

	if err := config.LoadConfig("config.env"); err != nil {
		logger.Err("read configuration", err)
		return
	}

	cfg := config.New()

	dbUrl, err := cfg.Postgres.ConnectionURL()
	if err != nil {
		logger.Err("read db url", err)
		return
	}

	db := postgres.New(dbUrl, time.Duration(cfg.Postgres.ConnTimeout)*time.Second)

	if err = db.Start(ctx); err != nil {
		logger.Err("connection db", err)
		return
	}
	defer db.Stop()

	srv := grpcServer.New(cfg.GRPC.ConnectionURL(), time.Duration(cfg.GRPC.Timeout)*time.Second, app.New(logger, db))

	go func() {
		<-ctx.Done()
		srv.Stop()
	}()

	err = srv.Run()
	if err != nil {
		logger.Err("run gRPC", err)
		return
	}

}
