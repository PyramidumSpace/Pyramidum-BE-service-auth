package app

import (
	"fmt"
	"log/slog"
	"time"

	grpcapp "github.com/pyramidum-space/backend-service-auth/internal/app/grpc"
	"github.com/pyramidum-space/backend-service-auth/internal/services/auth"
	"github.com/pyramidum-space/backend-service-auth/internal/storage/postgres"
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
) (*App, error) {
	const op = "app.New"
	storage, err := postgres.New(storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	authService := auth.New(log, storage, storage, tokenTTL, secret)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{GRPCServer: grpcApp}, nil
}
