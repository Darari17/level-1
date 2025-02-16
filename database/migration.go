package database

import (
	"fmt"
	"level-1/model/domain"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if db == nil {
		return fmt.Errorf("DB Connection is nil")
	}

	var book domain.Book

	err := db.AutoMigrate(&book)
	return err
}
