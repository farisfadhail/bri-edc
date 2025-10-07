package services

import (
	"bri-edc/api/internal/repositories"

	"gorm.io/gorm"
)

type TransactionService struct {
	db              *gorm.DB
	transactionRepo *repositories.TransactionRepository
	merchantRepo    *repositories.MerchantRepository
	terminalRepo    *repositories.TerminalRepository
}

func NewTransactionService(
	db *gorm.DB,
	transactionRepo *repositories.TransactionRepository,
	merchantRepo *repositories.MerchantRepository,
	terminalRepo *repositories.TerminalRepository,
) *TransactionService {
	return &TransactionService{
		db:              db,
		transactionRepo: transactionRepo,
		merchantRepo:    merchantRepo,
		terminalRepo:    terminalRepo,
	}
}
