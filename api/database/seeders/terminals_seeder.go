package seeders

import "gorm.io/gorm"

func TerminalsSeeder(db *gorm.DB) {
	terminals := []map[string]interface{}{
		{"id": 1, "terminal_id": "T01", "merchant_id": "MCH123", "location": "Location A"},
		{"id": 2, "terminal_id": "T02", "merchant_id": "MCH123", "location": "Location B"},
		{"id": 3, "terminal_id": "T03", "merchant_id": "MCH124", "location": "Location C"},
	}

	for _, t := range terminals {
		err := db.Table("terminals").Create(t).Error
		if err != nil {
			panic("Failed to seed terminals: " + err.Error())
		}
	}
}
