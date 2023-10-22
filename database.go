package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Database *gorm.DB = nil

func initDatabase() {
	var err error = nil

	databaseUser := os.Getenv("DB_USERNAME")
	databasePassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%v:%v@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword)

	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Fatal(err)
		return
	}
}
