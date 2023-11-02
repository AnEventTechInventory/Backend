package storageManager

import (
	"errors"
	"fmt"
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceStorageManager struct {
	StorageInterface
	db *gorm.DB
}

func NewDeviceStorageManager() *DeviceStorageManager {
	return &DeviceStorageManager{
		db: database.Database,
	}
}

func (manager *DeviceStorageManager) Add(device *registry.Device) error {
	// Check if there is a device with the same name
	manager.db.Find(&registry.Device{}, "name = ?", device.Name)
	if !errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
		return errors.New("device with the same name already exists")
	}
	// ignore id if sent. Generate a new unique id.
	device.Id = uuid.New().String()

	// verify that the device is valid
	if err := device.Validate(manager.db); err != nil {
		return err
	}

	// insert into database
	manager.db.Create(device)
	if manager.db.Error != nil {
		return manager.db.Error
	}
	return nil
}

func (manager *DeviceStorageManager) Get(id string) (*registry.Device, error) {
	// verify that the id is valid
	if err := validateUUID(id); err != nil {
		return nil, err
	}

	// grab the device from the database
	var device *registry.Device
	manager.db.First(device, "id = ?", id)
	err := manager.db.Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Find device with id: %v. Error: %v", device.Id, err.Error()))
	}
	return device, nil
}

func (manager *DeviceStorageManager) List() ([]*registry.Device, error) {
	var devices []*registry.Device
	if err := manager.db.Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}

func (manager *DeviceStorageManager) Update(device *registry.Device) error {
	// Check if device exists
	var oldDevice *registry.Device
	manager.db.First(oldDevice, "id = ?", device.Id)
	if errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("Find device with id: %v. Error: Device does not exist", device.Id))
	}

	if err := device.Validate(manager.db); err != nil {
		return err
	}

	// Update
	manager.db.Model(oldDevice).Updates(device)
	if manager.db.Error != nil {
		return manager.db.Error
	}
	return nil
}

func (manager *DeviceStorageManager) Delete(id string) error {
	// Check if device exists
	var device *registry.Device
	manager.db.First(device, "id = ?", id)
	if errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("Find device with id: %v. Error: Device does not exist", device.Id))
	}
	// Delete
	manager.db.Delete(device)
	if manager.db.Error != nil {
		return manager.db.Error
	}
	return nil
}
