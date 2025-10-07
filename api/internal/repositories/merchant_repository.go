package repositories

import (
	"bri-edc/api/models"

	"gorm.io/gorm"
)

type MerchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) *MerchantRepository {
	return &MerchantRepository{db: db}
}

func (r *MerchantRepository) FindByID(id string) (*models.Merchant, error) {
	var merchant *models.Merchant
	err := r.db.Preload("Terminal").Where("merchant_id = ?", id).First(&merchant).Error
	if err != nil {
		return nil, err
	}

	return merchant, nil
}
