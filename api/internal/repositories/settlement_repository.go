package repositories

import (
	"bri-edc/api/models"

	"gorm.io/gorm"
)

type SettlementRepository struct {
	db *gorm.DB
}

func NewSettlementRepository(db *gorm.DB) *SettlementRepository {
	return &SettlementRepository{db: db}
}

func (r *SettlementRepository) Save(settlement *models.Settlement) error {
	return r.db.Save(settlement).Error
}
