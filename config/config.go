package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	logger *Logger
	db     *gorm.DB
)

func Init() error {
	var err error

	db, err = InitSQLite()
	if err != nil {
		return fmt.Errorf("error initializing SQLite in config: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func InitLogger(prefix string) *Logger {
	//Init new logger
	logger = NewLogger(prefix)
	return logger
}
