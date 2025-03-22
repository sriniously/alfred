package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sriniously/alfred/internal/server"
)

type HealthHandler struct {
	server *server.Server
}

func NewHealthHandler(server *server.Server) *HealthHandler {
	return &HealthHandler{server: server}
}

func (h *HealthHandler) CheckHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
