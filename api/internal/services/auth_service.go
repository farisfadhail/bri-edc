package services

import (
	"bri-edc/api/internal/repositories"

	"gorm.io/gorm"
)

type AuthService struct {
	db       *gorm.DB
	authRepo *repositories.AuthRepository
}

func NewAuthService(db *gorm.DB, authRepo *repositories.AuthRepository) *AuthService {
	return &AuthService{
		db:       db,
		authRepo: authRepo,
	}
}
