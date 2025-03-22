package services

import (
	"github.com/sriniously/alfred/internal/repositories"
	"github.com/sriniously/alfred/internal/server"
)

type Services struct{}

func NewServices(s *server.Server, repos *repositories.Repositories) *Services {
	return &Services{}
}
