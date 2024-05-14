package main

import (
	"fmt"
	"github.com/pyramidum-space/backend-service-auth/internal/app"
	"github.com/pyramidum-space/backend-service-auth/internal/config"
	"github.com/pyramidum-space/backend-service-auth/internal/env"
	"github.com/pyramidum-space/backend-service-auth/internal/lib/logger/sl"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	env.MustLoadEnv()

	fmt.Println("Starting service-auth...")
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	application, err := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL, cfg.AppSecretKey)
	if err != nil {
		log.Error("unable to create application", sl.Err(err))
		os.Exit(1)
	}

	go func() {
		application.GRPCServer.MustRun()
	}()

	// Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// Waiting for SIGINT (pkill -2) or SIGTERM
	<-stop

	// initiate graceful shutdown
	application.GRPCServer.Stop() // Assuming GRPCServer has Stop() method for graceful shutdown
	log.Info("Gracefully stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
