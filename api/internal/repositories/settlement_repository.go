package repositories

import (
	"bri-edc/api/models"

	"gorm.io/gorm"
)

type SettlementRepository struct {
}

func NewSettlementRepository() *SettlementRepository {
	return &SettlementRepository{}
}

func (r *SettlementRepository) Save(db *gorm.DB, settlement *models.Settlement) error {
	return db.Save(settlement).Error
}
