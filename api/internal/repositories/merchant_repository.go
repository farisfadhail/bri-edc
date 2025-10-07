package repositories

import (
	"bri-edc/api/models"

	"gorm.io/gorm"
)

type MerchantRepository struct {
}

func NewMerchantRepository() *MerchantRepository {
	return &MerchantRepository{}
}

func (r *MerchantRepository) GetByID(db *gorm.DB, id string, preload ...string) (*models.Merchant, error) {
	var merchant *models.Merchant
	query := db.Model(&merchant).Where("merchant_id = ?", id)

	for _, p := range preload {
		query = query.Preload(p)
	}

	err := query.First(&merchant).Error
	if err != nil {
		return nil, err
	}

	return merchant, nil
}
