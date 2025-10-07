package seeders

import (
	"bri-edc/api/config"

	"gorm.io/gorm"
)

func RunAllSeeder(db *gorm.DB) {
	config.LoadEnv()

	MerchantsSeeder(db)
	TerminalsSeeder(db)
	UserSeeder(db)
}
