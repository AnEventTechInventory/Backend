package storageManager

import (
	"errors"
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/AnEventTechInventory/Backend/pkg/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ManufacturerStorageManager struct {
	StorageInterface
	db *gorm.DB
}

func NewManufacturerStorageManager() *ManufacturerStorageManager {
	return &ManufacturerStorageManager{
		db: database.Get(),
	}
}

func (manager *ManufacturerStorageManager) Add(manufacturer *registry.Manufacturer) error {
	// Check if manufacturer already exists
	manager.db.Find(&manufacturer, "name = ?", manufacturer.Name)
	if !errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
		return errors.New("manufacturer already exists")
	}

	// Add manufacturer
	manufacturer.ID = uuid.New()

	// verify manufacturer is valid
	if err := manufacturer.Validate(nil); err != nil {
		return err
	}

	manager.db.Create(manufacturer)
	if manager.db.Error != nil {
		return manager.db.Error
	}
	return nil
}

func (manager *ManufacturerStorageManager) Get(id string) (*registry.Manufacturer, error) {
	// verify that the id is valid
	if err := util.ValidateUUID(id); err != nil {
		return nil, err
	}

	// grab the manufacturer from the database
	manufacturer := &registry.Manufacturer{}
	manager.db.First(manufacturer, "id = ?", id)
	if manager.db.Error != nil {
		return nil, manager.db.Error
	}
	return manufacturer, nil
}

func (manager *ManufacturerStorageManager) List() ([]*registry.Manufacturer, error) {
	// grab all manufacturers from the database
	var manufacturers []*registry.Manufacturer
	manager.db.Find(&manufacturers)
	if manager.db.Error != nil {
		return nil, manager.db.Error
	}
	return manufacturers, nil
}

func (manager *ManufacturerStorageManager) Update(manufacturer *registry.Manufacturer) error {
	// Check if manufacturer exists
	var oldManufacturer *registry.Manufacturer
	manager.db.Find(&manufacturer, "id = ?", manufacturer.ID)
	if errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
		return errors.New("manufacturer does not exist")
	}

	// Update
	if err := manufacturer.Validate(nil); err != nil {
		return err
	}

	manager.db.Model(oldManufacturer).Updates(manufacturer)
	if manager.db.Error != nil {
		return manager.db.Error
	}
	return nil
}

func (manager *ManufacturerStorageManager) Delete(id string) error {
	// Check if manufacturer exists
	var manufacturer *registry.Manufacturer
	manager.db.First(manufacturer, "id = ?", id)
	if errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
		return errors.New("manufacturer does not exist")
	}
	// Delete
	manager.db.Delete(manufacturer)
	if manager.db.Error != nil {
		return manager.db.Error
	}
	return nil
}
