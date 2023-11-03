package main

import (
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/httpServer"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func main() {
	startServer()

	testManufacturer := registry.Manufacturer{
		Id:          uuid.New().String(),
		Name:        "Test Manufacturer",
		Description: "testing manufacturer",
	}

	testLocation := registry.Location{
		Id:   uuid.New().String(),
		Name: "test location",
	}

	testDevice := registry.Device{
		Id:           uuid.New().String(),
		Name:         "Test",
		Manufacturer: testManufacturer.Id,
		Description:  "blablabla",
		Quantity:     10,
		Contents:     nil,
		Location:     testLocation.Id,
		Model:        gorm.Model{},
	}

	database.Database.Create(testDevice)

	// Your application code here
	httpServer.RunHttpServer()

	stopServer()
}
