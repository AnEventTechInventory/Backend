package main

import (
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/httpServer"
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

	contentTest := &registry.Device{
		Description:  "blablabla",
		ID:           uuid.New(),
		Name:         "Test_Content",
		Manufacturer: testManufacturer,
		Location:     testLocation,
		Quantity:     10,
		Contents:     "",
	}

	testDevice := &registry.Device{
		Description:  "blablabla",
		ID:           uuid.New(),
		Name:         "Test",
		Manufacturer: testManufacturer,
		Location:     testLocation,
		Quantity:     10,
		Contents:     contentTest.ID.String(),
	}

	database.Get().Migrator().DropTable(&registry.Device{})
	database.Get().Migrator().DropTable(&registry.Location{})
	database.Get().Migrator().DropTable(&registry.Manufacturer{})

	database.Get().AutoMigrate(&registry.Device{})
	database.Get().AutoMigrate(&registry.Location{})
	database.Get().AutoMigrate(&registry.Manufacturer{})

	database.Get().Create(testManufacturer)
	database.Get().Create(testLocation)
	database.Get().Create(contentTest)
	database.Get().Create(testDevice)

	// Your application code here
	httpServer.RunHttpServer()

	stopServer()
}
