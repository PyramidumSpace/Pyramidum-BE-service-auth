package app

import (
	grpcapp "github.com/g-vinokurov/pyramidum-backend-service-auth/internal/app/api"
	"log/slog"
	"time"

	"github.com/g-vinokurov/pyramidum-backend-service-auth/internal/services/auth"
	"github.com/g-vinokurov/pyramidum-backend-service-auth/internal/storage/postgres"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
	secret string,
) *App {
	storage, err := postgres.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, tokenTTL, secret)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{GRPCServer: grpcApp}
}
