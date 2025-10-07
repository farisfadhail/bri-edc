package models

import "time"

type Merchant struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	MerchantID string `json:"merchant_id" gorm:"type:varchar(100);unique;not null"`
	Name       string `json:"name" gorm:"not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Terminal []Terminal `json:"terminal,omitempty" gorm:"foreignKey:MerchantID;references:MerchantID"`
}
