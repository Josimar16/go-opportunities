package config

import (
	"github.com/Josimar16/go-opportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database does not exist, creating a new one")

		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			logger.Errorf("Error creating the database: %v", err)
			return nil, err
		}

		defer file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error initializing the database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opportunity{})

	if err != nil {
		logger.Errorf("Error migrating the database: %v", err)
		return nil, err
	}

	return db, nil
}
