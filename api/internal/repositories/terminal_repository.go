package repositories

import (
	"bri-edc/api/models"

	"gorm.io/gorm"
)

type TerminalRepository struct {
	db *gorm.DB
}

func NewTerminalRepository(db *gorm.DB) *TerminalRepository {
	return &TerminalRepository{db: db}
}

func (r *TerminalRepository) GetByID(id string) (*models.Terminal, error) {
	var terminal *models.Terminal
	err := r.db.Preload("Merchant").First(&terminal, "terminal_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return terminal, nil
}

func (r *TerminalRepository) GetByMerchantID(merchantID string) ([]models.Terminal, error) {
	var terminals []models.Terminal
	err := r.db.Preload("Merchant").Where("merchant_id = ?", merchantID).Find(&terminals).Error
	if err != nil {
		return nil, err
	}

	return terminals, nil
}
