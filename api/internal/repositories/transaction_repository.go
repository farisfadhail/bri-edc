package repositories

import (
	"bri-edc/api/models"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository struct {
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) Create(db *gorm.DB, transaction models.Transaction) error {
	return db.Create(&transaction).Error
}

func (r *TransactionRepository) GetByID(db *gorm.DB, id string, preload ...string) (*models.Transaction, error) {
	var transaction *models.Transaction

	query := db.Model(&transaction).Where("transaction_id = ?", id)

	for _, p := range preload {
		query = query.Preload(p)
	}

	err := query.First(&transaction).Error

	return transaction, err
}

func (r *TransactionRepository) GetBatchTransactions(db *gorm.DB, date string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	startOfDay, err := time.Parse("2006-01-02", date)
	if err != nil {
		return transactions, err
	}
	endOfDay := startOfDay.Add(24 * time.Hour)

	err = db.Model(&transactions).
		Where("timestamp >= ? AND timestamp < ?", startOfDay, endOfDay).
		Where("is_settled = ?", false).
		Find(&transactions).Error

	return transactions, err
}

func (r *TransactionRepository) Update(db *gorm.DB, transaction *models.Transaction) error {
	return db.Save(transaction).Error
}
