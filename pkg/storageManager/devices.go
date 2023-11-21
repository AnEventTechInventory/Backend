package storageManager

import (
	"errors"
	"fmt"
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/AnEventTechInventory/Backend/pkg/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceStorageManager struct {
	StorageInterface[registry.Device]
	db *gorm.DB
}

func NewDeviceStorageManager() *DeviceStorageManager {
	return &DeviceStorageManager{
		db: database.Get(),
	}
}

func (manager *DeviceStorageManager) Add(device *registry.Device) error {
	// Check if there is a device with the same name
	manager.db.Find(&registry.Device{}, "name = ?", device.Name)
	if !errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
		return errors.New("device with the same name already exists")
	}
	// ignore id if sent. Generate a new unique id.
	device.ID = uuid.New()

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
	// check if the id is valid
	if err := util.ValidateUUID(id); err != nil {
		return nil, err
	}

	// grab the device from the database
	var device = &registry.Device{}
	err := manager.db.First(device, "id = ?", id).Error

	if err != nil {
		return nil, fmt.Errorf("find device with id: %v. Error: %v", id, err.Error())
	}
	return device, nil
}

func (manager *DeviceStorageManager) List() ([]*uuid.UUID, error) {
	devices := make([]*uuid.UUID, 0)
	// only return the IDs
	manager.db.Model(&registry.Device{}).Select("id").Find(&devices)
	if manager.db.Error != nil {
		return nil, manager.db.Error
	}

	return devices, nil
}

func (manager *DeviceStorageManager) Update(device *registry.Device) error {
	// Check if device exists
	oldDevice := &registry.Device{}
	manager.db.First(oldDevice, "id = ?", device.ID)
	if errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("find device with id: %v. Error: Device does not exist", device.ID)
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
	device := &registry.Device{}
	manager.db.First(device, "id = ?", id)
	if errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("find device with id: %v. Error: Device does not exist", device.ID)
	}
	// Delete
	manager.db.Delete(device)
	if manager.db.Error != nil {
		return manager.db.Error
	}
	return nil
}

func (manager *DeviceStorageManager) UpdateContents(id string, contents []string) error {
	// Check if device exists
	var device *registry.Device
	manager.db.First(device, "id = ?", id)
	if errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("find device with id: %v. Error: Device does not exist", device.ID)
	}

	// Check any entry is a valid uuid and exists
	for _, entry := range contents {
		if err := util.ValidateUUID(entry); err != nil {
			return err
		}
		var contentDevice *registry.Device
		manager.db.First(contentDevice, "id = ?", entry)
		if errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("find device with id: %v. Error: Device does not exist", entry)
		}
	}

	// Update
	manager.db.Model(device).Update("contents", contents)
	if manager.db.Error != nil {
		return manager.db.Error
	}
	return nil
}
