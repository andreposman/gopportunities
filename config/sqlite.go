package config

import (
	"os"

	"github.com/andreposman/gopportunities/internal/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	logger := InitLogger("- SQLite -")
	dbPath := "./internal/db/main.db"

	//? Check if DB exists before
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating...")
		if err = os.MkdirAll("./internal/db", os.ModePerm); err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	//? create and conntect to db
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errf("SQLite opening error: %v", err)
		return nil, err
	}

	//? Migrate schema
	if err = db.AutoMigrate(&schemas.Opening{}); err != nil {
		logger.Errf("SQLite auto migration error: %v", err)
		return nil, err
	}

	return db, nil
}
