package database

import (
	"fmt"
	"github.com/AnEventTechInventory/Backend/pkg/logger"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var database *gorm.DB = nil

func Get() *gorm.DB {
	if database == nil {
		InitDatabase()
	}
	return database
}

func InitDatabase() bool {
	var err error = nil

	databaseUser := os.Getenv("DB_USERNAME")
	if databaseUser == "" {
		logger.Get().Fatalf("Database username is not set")
		return false
	}

	databasePassword := os.Getenv("DB_PASSWORD")
	if databasePassword == "" {
		logger.Get().Fatalf("Database password is not set")
		return false
	}
	databaseURL := "mysql"

	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/inventory?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword, databaseURL)

	logger.Get().Println("Starting database...")

	database, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{})
	if err != nil {
		logger.Get().Fatal(err)
		return false
	}

	// Migrate the schema
	if err := database.AutoMigrate(&registry.Location{}); err != nil {
		logger.Get().Fatal(err)
		return false
	}

	if err := database.AutoMigrate(&registry.Manufacturer{}); err != nil {
		logger.Get().Fatal(err)
		return false
	}

	if err := database.AutoMigrate(&registry.Type{}); err != nil {
		logger.Get().Fatal(err)
		return false
	}

	if err := database.AutoMigrate(&registry.Device{}); err != nil {
		logger.Get().Fatal(err)
		return false
	}

	database.Statement.RaiseErrorOnNotFound = true
	logger.Get().Println("Database started successfully")
	return true
}
