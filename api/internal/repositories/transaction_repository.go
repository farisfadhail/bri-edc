package repositories

import (
	"bri-edc/api/internal/validations"
	"bri-edc/api/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(transaction validations.CreateTransactionRequest) error {
	return r.db.Create(&transaction).Error
}

func (r *TransactionRepository) GetByID(id string) (*models.Transaction, error) {
	var transaction *models.Transaction

	err := r.db.Preload("Merchant").Preload("Terminal").Where("transaction_id = ?", id).First(&transaction).Error

	return transaction, err
}
