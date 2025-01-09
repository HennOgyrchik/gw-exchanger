package app

import (
	"gw-exchanger/internal/storages"
	"gw-exchanger/pkg/logs"
)

type App struct {
	log     *logs.Log
	storage storages.Storage
}
