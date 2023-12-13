package config

import (
	sqlite "github.com/bboehmke/gorm-sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
