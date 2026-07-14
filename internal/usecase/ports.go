package usecase

import "baut001/backend/internal/entity"

type AuthRepositoryPort interface {
	FindUserByUsername(username string) (entity.User, error)
	ListScopes(userID string) ([]string, error)
	GetUserByID(userID string) (entity.User, error)
}
