package seeders

import (
	"bri-edc/api/utils"

	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) {
	password, _ := utils.HashPassword("password")

	users := []map[string]interface{}{
		{"id": 1, "username": "admin", "password": password},
	}

	for _, u := range users {
		err := db.Table("users").Create(u).Error
		if err != nil {
			panic("Failed to seed users: " + err.Error())
		}
	}
}
