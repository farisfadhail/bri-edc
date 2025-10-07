package models

import "time"

type Terminal struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	TerminalID string `json:"terminal_id" gorm:"type:varchar(100);unique;not null"`
	MerchantID string `json:"merchant_id" gorm:"type:varchar(100);unique;not null"`
	Name       string `json:"name" gorm:"not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`

	Merchant Merchant `json:"merchant" gorm:"foreignKey:MerchantID;references:MerchantID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
