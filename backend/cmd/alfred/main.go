package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sriniously/alfred/internal/config"
	"github.com/sriniously/alfred/internal/handlers"
	"github.com/sriniously/alfred/internal/logger"
	"github.com/sriniously/alfred/internal/repositories"
	"github.com/sriniously/alfred/internal/router"
	"github.com/sriniously/alfred/internal/server"
	"github.com/sriniously/alfred/internal/services"
)

const DefaultContextTimeout = 30

func main() {
	// load config
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config:" + err.Error())
	}

	logger := logger.NewLogger("debug", cfg.Primary.Env == "production")

	server, err := server.New(cfg, &logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to initialize server")
	}

	// repositories
	repositories := repositories.NewRepositories()

	// services
	services := services.NewServices(server, repositories)

	// handlers
	handlers := handlers.NewHandlers(server, services)

	// router
	router := router.NewRouter(server, handlers)

	server.SetupHTTPServer(router)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	// Start server
	go func() {
		if err = server.Start(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeout*time.Second)

	if err = server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("server forced to shutdown")
	}
	stop()
	cancel()

	log.Info().Msg("server exited properly")
}
