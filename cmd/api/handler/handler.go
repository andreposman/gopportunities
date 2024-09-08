package handler

import (
	"github.com/andreposman/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.InitLogger("- handler - ")
	db = config.GetSQLite()
}
