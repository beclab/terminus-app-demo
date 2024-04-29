package handlers

import (
	"github.com/beclab/terminus-app-demo/pkg/storage"
	"github.com/beclab/terminus-app-demo/pkg/storage/strategies/postgres"
)

type handler struct {
	db storage.Operator
}

func New() *handler {
	return &handler{db: postgres.New()}
}
