package repository

import (
	"context"
	"github.com/g-vinokurov/pyramidum-backend-service-auth/internal/domain/models"
)

type UserRepository interface {
	SaveUser(
		ctx context.Context,
		email string,
		passHash []byte,
	) (uid int64, err error)

	GetUser(ctx context.Context, email string) (models.User, error)
}
