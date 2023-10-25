package storageManager

import (
	"errors"
	"fmt"
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceStore interface {
	Add(device *registry.Device) error
	Get(id string) (*registry.Device, error)
	List() ([]*registry.Device, error)
	Update(device *registry.Device) error
	Delete(id string) error
}

type StorageManager struct {
	db *gorm.DB
}

func NewDeviceStorageManager(db *gorm.DB) *StorageManager {
	return &StorageManager{db: db}
}

func (manager *StorageManager) Add(device *registry.Device) error {
	// Check if there is a device with the same name
	manager.db.Find(&registry.Device{}, "name = ?", device.Name)
	if !errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
		return database.InsertError{ErrorMessage: "Device with the same name already exists"}
	}
	// ignore id if sent. Generate a new unique id.
	device.Id = uuid.New().String()

	// verify that the device is valid
	if device.Name == "" {
		return database.InsertError{ErrorMessage: "Device name cannot be empty"}
	}
	if device.Quantity < 0 {
		return database.InsertError{ErrorMessage: "Device quantity cannot be negative"}
	}

	// insert into database
	manager.db.Create(device)
	if manager.db.Error != nil {
		return database.InsertError{ErrorMessage: manager.db.Error.Error()}
	}
	return nil
}

func (manager *StorageManager) Get(id string) (*registry.Device, error) {
	// verify that the id is valid
	if _, err := uuid.Parse(id); err != nil {
		return nil, database.QueryError{ErrorMessage: "Invalid device id"}
	}

	// grab the device from the database
	var device *registry.Device
	manager.db.First(device, "id = ?", id)
	err := manager.db.Error
	if err != nil {
		return nil, database.QueryError{Query: fmt.Sprintf("Find device with id: %v", device.Id), ErrorMessage: err.Error()}
	}
	return device, nil
}

func (manager *StorageManager) List() ([]*registry.Device, error) {
	var devices []*registry.Device
	if err := manager.db.Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}

func (manager *StorageManager) Update(device *registry.Device) error {
	// Check if device exists
	var oldDevice *registry.Device
	manager.db.First(oldDevice, "id = ?", device.Id)
	if errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
		return database.QueryError{Query: fmt.Sprintf("Find device with id: %v", device.Id), ErrorMessage: "Device does not exist"}
	}
	// Update
	manager.db.Model(oldDevice).Updates(device)
	if manager.db.Error != nil {
		return database.UpdateError{ErrorMessage: manager.db.Error.Error()}
	}
	return nil
}

func (manager *StorageManager) Delete(id string) error {
	// Check if device exists
	var device *registry.Device
	manager.db.First(device, "id = ?", id)
	if errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
		return database.QueryError{Query: fmt.Sprintf("Find device with id: %v", id), ErrorMessage: "Device does not exist"}
	}
	// Delete
	manager.db.Delete(device)
	if manager.db.Error != nil {
		return database.DeleteError{ErrorMessage: manager.db.Error.Error()}
	}
	return nil
}
