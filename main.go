package main

import (
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/httpServer"
	"github.com/AnEventTechInventory/Backend/pkg/logger"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/google/uuid"
)

func main() {
	startServer()

	testManufacturer := &registry.Manufacturer{
		ID:          uuid.New(),
		Name:        "Test Manufacturer",
		Description: "testing manufacturer",
	}

	testLocation := &registry.Location{
		ID:          uuid.New(),
		Name:        "test location",
		Description: "testing location",
	}

	testType := &registry.Type{
		ID:          uuid.New(),
		Name:        "Test Type",
		Description: "testing type",
	}

	contentTest := &registry.Device{
		Description:  "blablabla",
		ID:           uuid.New(),
		Name:         "Test_Content",
		Manufacturer: testManufacturer,
		Location:     testLocation,
		Type:         testType,
		Quantity:     10,
		Contents:     "",
	}

	testDevice := &registry.Device{
		Description:  "blablabla",
		ID:           uuid.New(),
		Name:         "Test",
		Manufacturer: testManufacturer,
		Location:     testLocation,
		Type:         testType,
		Quantity:     10,
		Contents:     contentTest.ID.String(),
	}

	if contentTest == nil {
		logger.Get().Fatal("test was nil")
		return
	}
	contentTest.Validate(database.Get())
	testDevice.Validate(database.Get())

	err := database.Get().Migrator().DropTable(
		&registry.Location{},
		&registry.Manufacturer{},
		&registry.Type{},
		&registry.Device{},
	)
	if err != nil {
		logger.Get().Fatal(err)
		return
	}

	err = database.Get().AutoMigrate(
		&registry.Location{},
		&registry.Manufacturer{},
		&registry.Type{},
		&registry.Device{},
	)
	if err != nil {
		logger.Get().Fatal(err)
		return
	}

	database.Get().Create(testManufacturer)
	database.Get().Create(testLocation)
	database.Get().Create(testType)
	database.Get().Create(contentTest)
	database.Get().Create(testDevice)

	// Your application code here
	httpServer.RunHttpServer()

	stopServer()
}
