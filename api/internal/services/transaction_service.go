package services

import (
	"bri-edc/api/internal/repositories"
	"bri-edc/api/internal/validations"
	"bri-edc/api/models"
	"bri-edc/api/resources"
	"bri-edc/api/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TransactionService struct {
	db              *gorm.DB
	transactionRepo *repositories.TransactionRepository
	merchantRepo    *repositories.MerchantRepository
	terminalRepo    *repositories.TerminalRepository
	settlementRepo  *repositories.SettlementRepository
}

func NewTransactionService(
	db *gorm.DB,
	transactionRepo *repositories.TransactionRepository,
	merchantRepo *repositories.MerchantRepository,
	terminalRepo *repositories.TerminalRepository,
	settlementRepo *repositories.SettlementRepository,
) *TransactionService {
	return &TransactionService{
		db:              db,
		transactionRepo: transactionRepo,
		merchantRepo:    merchantRepo,
		terminalRepo:    terminalRepo,
		settlementRepo:  settlementRepo,
	}
}

func (s *TransactionService) Sale(req validations.CreateTransactionRequest) (*resources.SaleResource, error) {
	tx := s.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	_, err := s.merchantRepo.GetByID(tx, req.MerchantID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = s.terminalRepo.GetByID(tx, req.TerminalID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	transactionID := utils.GenerateTransactionNumber()
	timestamp, err := time.Parse(time.RFC3339, req.Timestamp)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	msg := fmt.Sprintf("%s|%s|%s|%d|%s", transactionID, req.MerchantID, req.TerminalID, req.Amount, timestamp.UTC().Format(time.RFC3339))
	hmacValue := utils.GenerateHMAC(msg)

	newTransaction := &models.Transaction{
		TransactionID: transactionID,
		MerchantID:    req.MerchantID,
		TerminalID:    req.TerminalID,
		Amount:        req.Amount,
		CardNumber:    req.CardNumber,
		Status:        "approved",
		Timestamp:     timestamp,
		HMAC:          hmacValue,
	}

	if err := s.transactionRepo.Create(tx, *newTransaction); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &resources.SaleResource{
		TransactionID: transactionID,
		Status:        newTransaction.Status,
		Message:       "Transaction authorized",
	}, nil
}

func (s *TransactionService) Settlement() (*resources.SettlementResource, error) {
	tx := s.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	currentDate := time.Now().Format("2006-01-02")

	transactions, err := s.transactionRepo.GetBatchTransactions(tx, currentDate)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	batchID := utils.GenerateBatchNumber()
	totalCount := len(transactions)
	approvedCount := 0
	declinedCount := 0
	totalAmount := 0
	for _, t := range transactions {
		msg := fmt.Sprintf("%s|%s|%s|%d|%s", t.TransactionID, t.MerchantID, t.TerminalID, t.Amount, t.Timestamp.UTC().Format(time.RFC3339))
		expectedHMAC := utils.GenerateHMAC(msg)
		if t.HMAC != expectedHMAC {
			tx.Rollback()
			return nil, fmt.Errorf("invalid HMAC for transaction %s %s %s", t.TransactionID, t.HMAC, expectedHMAC)
		}

		if t.Status == "approved" {
			approvedCount++
			totalAmount += t.Amount
		}

		if t.Status == "declined" {
			declinedCount++
		}

		t.IsSettled = true
		_ = s.transactionRepo.Update(tx, &t)
	}

	newSettlement := &models.Settlement{
		BatchID:     batchID,
		TotalCount:  totalCount,
		Approved:    approvedCount,
		Declined:    declinedCount,
		TotalAmount: totalAmount,
	}

	err = s.settlementRepo.Save(tx, newSettlement)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &resources.SettlementResource{
		BatchID:     newSettlement.BatchID,
		TotalCount:  newSettlement.TotalCount,
		Approved:    newSettlement.Approved,
		Declined:    newSettlement.Declined,
		TotalAmount: newSettlement.TotalAmount,
	}, nil
}
