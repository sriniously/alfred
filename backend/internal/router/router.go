package router

import (
	"github.com/labstack/echo/v4"
	"github.com/sriniously/alfred/internal/handlers"
	"github.com/sriniously/alfred/internal/middlewares"
	v1 "github.com/sriniously/alfred/internal/router/v1"
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

	v1Router := router.Group("/api/v1")
	v1.RegisterV1Routes(v1Router, h, middlewares)

	return router
}
