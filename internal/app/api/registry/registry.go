package registry

import (
	"github.com/g-vinokurov/pyramidum-backend-service-auth/internal/domain/service"
	"github.com/g-vinokurov/pyramidum-backend-service-auth/internal/infrastructure/persistence/postgres"
)

type Registry struct {
	Services *Services
	Storage  *postgres.Storage
}

func NewRegistry(authService service.Auth, storage *postgres.Storage) *Registry {
	return &Registry{Services: newServices(authService), Storage: storage}
}

type Services struct {
	Auth service.Auth
}

func newServices(auth service.Auth) *Services {
	return &Services{Auth: auth}
}
