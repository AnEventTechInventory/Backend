package database

import (
	"errors"
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
	databasePassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%v:%v@tcp(10.0.0.2:3306)/inventory?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword)

	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Logger.Fatal(err)
		return false
	}

	// Migrate the schema
	err = Database.AutoMigrate(&registry.Device{})
	if err != nil {
		logger.Logger.Fatal(err)
		return false
	}
	logger.Logger.Println("Database started successfully")
	return true
}

var MissingError = errors.New("no database connection")

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
