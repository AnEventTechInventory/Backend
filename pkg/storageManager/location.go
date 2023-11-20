package storageManager

import (
	"errors"
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LocationStorageManager struct {
	StorageInterface
	db *gorm.DB
}

func NewLocationStorageManager() *LocationStorageManager {
	return &LocationStorageManager{
		db: database.Get(),
	}
}

func (manager *LocationStorageManager) Add(location *registry.Location) error {
	// Verify name is unique
	manager.db.Find(&registry.Location{}, "name = ?", location.Name)
	if !errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
		return errors.New("location already exists")
	}

	location.ID = uuid.New()

	if err := location.Validate(nil); err != nil {
		return err
	}

	manager.db.Create(location)
	if manager.db.Error != nil {
		return manager.db.Error
	}
	return nil
}

func (manager *LocationStorageManager) Get(id string) (*registry.Location, error) {
	var location = &registry.Location{}
	manager.db.First(location, "id = ?", id)
	if manager.db.Error != nil {
		return nil, manager.db.Error
	}
	return location, nil
}

func (manager *LocationStorageManager) List() ([]*registry.Location, error) {
	var locations []*registry.Location
	manager.db.Find(&locations)
	if manager.db.Error != nil {
		return nil, manager.db.Error
	}
	return locations, nil
}

func (manager *LocationStorageManager) Update(location *registry.Location) error {
	manager.db.Save(location)
	if manager.db.Error != nil {
		return manager.db.Error
	}
	return nil
}

func (manager *LocationStorageManager) Delete(id string) error {
	manager.db.Delete(&registry.Location{}, "id = ?", id)
	if manager.db.Error != nil {
		return manager.db.Error
	}
	return nil
}
