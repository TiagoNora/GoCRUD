package config

import (
	"os"

	"github.com/TiagoNora/GoCRUD/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	err = db.AutoMigrate(&schemas.Product{}, &schemas.User{})
	user := schemas.User{}
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}
	result := db.Find(&user)
	rows := result.RowsAffected
	if rows == 0 {
		user := schemas.User{
			Name:     "Admin",
			Username: "Admin Admin",
			Email:    "Admin",
			Password: "adminPassword",
			Role:     "ADMIN",
		}

		if err := user.HashPassword(user.Password); err != nil {
			logger.Errorf("validation error: %v", err.Error())
			return nil, err
		}

		if err := db.Create(&user).Error; err != nil {
			logger.Errorf("error creating user: %v", err.Error())
			return nil, err
		}
	}
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}
