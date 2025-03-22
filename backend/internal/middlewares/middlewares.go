package middlewares

import "github.com/sriniously/alfred/internal/server"

type Middlewares struct {
	Global *GlobalMiddlewares
}

func NewMiddlewares(s *server.Server) *Middlewares {
	return &Middlewares{
		Global: NewGlobalMiddlewares(s),
	}
}
