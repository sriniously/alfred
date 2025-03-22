package router

import (
	"github.com/labstack/echo/v4"
	"github.com/sriniously/alfred/internal/handlers"
)

func registerSystemRoutes(r *echo.Echo, h *handlers.Handlers) {
	r.GET("/health", h.HealthHandler.CheckHealth)
}
