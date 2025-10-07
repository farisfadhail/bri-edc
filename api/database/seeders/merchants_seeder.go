package seeders

import "gorm.io/gorm"

func MerchantsSeeder(db *gorm.DB) {
	merchants := []map[string]interface{}{
		{"id": 1, "merchant_id": "MCH123", "name": "Merchant One"},
		{"id": 2, "merchant_id": "MCH124", "name": "Merchant Two"},
		{"id": 3, "merchant_id": "MCH125", "name": "Merchant Three"},
	}

	for _, m := range merchants {
		err := db.Table("merchants").Create(m).Error
		if err != nil {
			panic("Failed to seed merchants: " + err.Error())
		}
	}
}
