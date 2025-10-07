package repositories

import (
	"bri-edc/api/models"

	"gorm.io/gorm"
)

type TerminalRepository struct {
}

func NewTerminalRepository() *TerminalRepository {
	return &TerminalRepository{}
}

func (r *TerminalRepository) GetByID(db *gorm.DB, id string, preload ...string) (*models.Terminal, error) {
	var terminal *models.Terminal
	query := db.Model(&terminal).Where("terminal_id = ?", id)

	for _, p := range preload {
		query = query.Preload(p)
	}

	err := query.First(&terminal).Error
	if err != nil {
		return nil, err
	}

	return terminal, nil
}

func (r *TerminalRepository) GetByMerchantID(db *gorm.DB, merchantID string) ([]models.Terminal, error) {
	var terminals []models.Terminal
	err := db.Preload("Merchant").Where("merchant_id = ?", merchantID).Find(&terminals).Error
	if err != nil {
		return nil, err
	}

	return terminals, nil
}
