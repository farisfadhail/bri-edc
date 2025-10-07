package models

import "time"

type Transaction struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	TransactionID string    `json:"transaction_id" gorm:"type:varchar(100);unique"`
	MerchantID    string    `json:"merchant_id" gorm:"type:varchar(100);unique"`
	TerminalID    string    `json:"terminal_id" gorm:"type:varchar(100);unique"`
	Amount        int       `json:"amount"`
	CardNumber    string    `json:"card_number" gorm:"type:varchar(20)"`
	Status        string    `json:"status" gorm:"type:varchar(50)"`
	IsSettled     bool      `json:"is_settled" gorm:"default:false"`
	Timestamp     time.Time `json:"timestamp" gorm:"type:datetime"`
	HMAC          string    `json:"hmac" gorm:"type:varchar(255)"`

	CreatedAt time.Time `json:"created_at"`

	Merchant Merchant `json:"merchant,omitempty" gorm:"foreignKey:MerchantID;references:MerchantID"`
	Terminal Terminal `json:"terminal,omitempty" gorm:"foreignKey:TerminalID;references:TerminalID"`
}
