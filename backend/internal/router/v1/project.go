package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/sriniously/alfred/internal/handlers"
)

func registerProjectRoutes(r *echo.Group, h *handlers.Handlers) {
	projectRouter := r.Group("/projects")

	// projectRouter.GET("", h.ProjectHandler.GetProjects)
}
