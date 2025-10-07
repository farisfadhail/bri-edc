package models

import "time"

type Settlement struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	BatchID     string `json:"batch_id" gorm:"type:varchar(100);unique;not null"`
	TotalCount  int    `json:"total_count" gorm:"not null"`
	Approved    int    `json:"approved" gorm:"not null"`
	Declined    int    `json:"declined" gorm:"not null"`
	TotalAmount int    `json:"total_amount" gorm:"not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
