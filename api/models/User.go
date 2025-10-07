package models

import "time"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"password" gorm:"not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
