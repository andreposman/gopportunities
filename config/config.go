package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	logger *Logger
	db     *gorm.DB
)

func Init() error {

	return errors.New("Fake Error")
}

func InitLogger(prefix string) *Logger {
	//Init new logger
	logger = NewLogger(prefix)
	return logger
}
