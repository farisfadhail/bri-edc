package services

import (
	"bri-edc/api/internal/repositories"
	"bri-edc/api/internal/validations"
	"bri-edc/api/utils"

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

func (s *AuthService) Login(req validations.AuthRequest) (string, error) {
	user, err := s.authRepo.GetByUsername(s.db, req.Username)
	if err != nil {
		return "", err
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return "", err
	}

	isValid := utils.CheckPasswordHash(req.Password, hashPassword)
	if !isValid {
		return "", err
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
