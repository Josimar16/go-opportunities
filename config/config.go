package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	// DB is the connection string to the database
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing SQLite: %v", err)
	}

	return err
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
