package config

import (
	"gorm.io/gorm"
)

var (
	// DB is the connection string to the database
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
