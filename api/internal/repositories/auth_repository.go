package repositories

import (
	"bri-edc/api/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (r *AuthRepository) GetByUsername(db *gorm.DB, username string) (*models.User, error) {
	var user *models.User

	err := db.Model(&user).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
