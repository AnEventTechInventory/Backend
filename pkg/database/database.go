package database

import (
	"fmt"
	"github.com/AnEventTechInventory/Backend/pkg/logger"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Database *gorm.DB = nil

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

	logger.Get().Fatalf("_ %v _ %v _", databaseUser, databasePassword)

	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/inventory?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword, databaseURL)

	logger.Get().Println("Starting database...")

	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Get().Fatal(err)
		return false
	}

	// Migrate the schema
	err = Database.AutoMigrate(&registry.Device{})
	if err != nil {
		logger.Get().Fatal(err)
		return false
	}
	logger.Get().Println("Database started successfully")
	return true
}

type TableError struct {
	TableName string
}

func (e TableError) Error() string {
	return fmt.Sprintf("Database tabble %v does not exist", e.TableName)
}

type QueryError struct {
	Query        string
	ErrorMessage string
}

func (e QueryError) Error() string {
	return fmt.Sprintf("Database query %v failed with error %v", e.Query, e.ErrorMessage)
}

type InsertError struct {
	Query        string
	ErrorMessage string
}

func (e InsertError) Error() string {
	return fmt.Sprintf("Database insert %v failed with error %v", e.Query, e.ErrorMessage)
}

type UpdateError struct {
	Query        string
	ErrorMessage string
}

func (e UpdateError) Error() string {
	return fmt.Sprintf("Database update %v failed with error %v", e.Query, e.ErrorMessage)
}
