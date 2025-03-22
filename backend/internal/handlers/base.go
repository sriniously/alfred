package handlers

import (
	"github.com/sriniously/alfred/internal/server"
	"github.com/sriniously/alfred/internal/validation"

	"github.com/labstack/echo/v4"
)

// Handler provides base functionality for all handlers
type Handler struct {
	server *server.Server
}

// NewHandler creates a new base handler
func NewHandler(s *server.Server) Handler {
	return Handler{server: s}
}

// HandlerFunc represents a typed handler function that processes a request and returns a response
type HandlerFunc[Req validation.Validatable, Res any] func(c echo.Context, req Req) (Res, error)

// HandlerFuncNoContent represents a typed handler function that processes a request without returning content
type HandlerFuncNoContent[Req validation.Validatable] func(c echo.Context, req Req) error

// Handle wraps a handler with validation and error handling
func Handle[Req validation.Validatable, Res any](
	h Handler,
	handler HandlerFunc[Req, Res],
	status int,
	req Req,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := validation.BindAndValidate(c, req); err != nil {
			h.server.Logger.Err(err).
				Str("handler", c.Path()).
				Msg("request validation failed")
			return err
		}

		res, err := handler(c, req)
		if err != nil {
			return err
		}

		return c.JSON(status, res)
	}
}

// HandleNoContent wraps a handler with validation and error handling for endpoints that don't return content
func HandleNoContent[Req validation.Validatable](
	h Handler,
	handler HandlerFuncNoContent[Req],
	status int,
	req Req,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := validation.BindAndValidate(c, req); err != nil {
			h.server.Logger.Err(err).
				Str("handler", c.Path()).
				Msg("request validation failed")
			return err
		}

		if err := handler(c, req); err != nil {
			return err
		}

		return c.NoContent(status)
	}
}
