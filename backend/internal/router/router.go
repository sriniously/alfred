package router

import (
	"github.com/labstack/echo/v4"
	"github.com/sriniously/alfred/internal/handlers"
	"github.com/sriniously/alfred/internal/middlewares"
	"github.com/sriniously/alfred/internal/server"
)

func NewRouter(s *server.Server, h *handlers.Handlers) *echo.Echo {
	router := echo.New()

	middlewares := middlewares.NewMiddlewares(s)

	router.HTTPErrorHandler = middlewares.Global.GlobalErrorHandler

	router.Use(
		middlewares.Global.CORS(),
		middlewares.Global.Secure(),
		middlewares.Global.RequestLogger(),
		middlewares.Global.Recover(),
	)

	registerSystemRoutes(router, h)

	return router
}
