package handlers

import (
	"github.com/sriniously/alfred/internal/server"
	"github.com/sriniously/alfred/internal/services"
)

type Handlers struct {
	HealthHandler *HealthHandler
}

func NewHandlers(server *server.Server, services *services.Services) *Handlers {
	return &Handlers{HealthHandler: NewHealthHandler(server)}
}
