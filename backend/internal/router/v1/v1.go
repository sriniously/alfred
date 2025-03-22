package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/sriniously/alfred/internal/handlers"
	"github.com/sriniously/alfred/internal/middlewares"
)

func RegisterV1Routes(r *echo.Group, h *handlers.Handlers, middlware *middlewares.Middlewares) {
	registerProjectRoutes(r, h)
}
