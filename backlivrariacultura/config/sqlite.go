package config

import (
	"os"

	"github.com/rafabcanedo/LibraryCulture/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if the database file exists
	_, error := os.Stat(dbPath)
	if os.IsNotExist(error) {
		logger.Info("database file not found, creating...")
		error = os.MkdirAll("./db", os.ModePerm)
		if error != nil {
			return nil, error
		}
		file, error := os.Create(dbPath)
		if error != nil {
			return nil, error
		}
		// Import closed file !!
		file.Close()
	}

	db, error := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if error != nil {
		logger.Errorf("sqlite books error: %v", error)
		return nil, error
	}
	error = db.AutoMigrate(&schemas.Book{})
	if error != nil {
		logger.Errorf("sqlite automigration error: %v", error)
		return nil, error
	}

	return db, nil
}
