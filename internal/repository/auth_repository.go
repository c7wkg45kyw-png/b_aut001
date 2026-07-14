package repository

import (
	"baut001/backend/internal/entity"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository { return &AuthRepository{db: db} }

func (r *AuthRepository) FindUserByUsername(username string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ? AND status = ?", username, "active").First(&user).Error
	return user, err
}

func (r *AuthRepository) ListScopes(userID string) ([]string, error) {
	var rows []entity.UserScope
	if err := r.db.Where("user_id = ?", userID).Order("scope ASC").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]string, 0, len(rows))
	for _, row := range rows {
		out = append(out, row.Scope)
	}
	return out, nil
}

func (r *AuthRepository) GetUserByID(userID string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("id::text = ? AND status = ?", userID, "active").First(&user).Error
	return user, err
}
