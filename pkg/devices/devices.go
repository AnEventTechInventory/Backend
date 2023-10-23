package devices

import (
	"errors"
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"regexp"
	"strings"
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

	// verify that the device contents are valid
	// can only a list of valid uuids followed by exactly one ', ' or be the end
	match, err := regexp.Match(`^([a-f\d]{8}(-[a-f\d]{4}){4}[a-f\d]{8}, )*[a-f\d]{8}(-[a-f\d]{4}){4}[a-f\d]{8}$`, []byte(device.Contents))
	if err != nil {
		return err
	}
	if !match {
		return database.InsertError{ErrorMessage: "Device contents must be a list of valid uuids separated by exactly one ', '"}
	}

	for _, content := range strings.Split(device.Contents, ", ") {
		// check if the content ids already exist
		manager.db.Find(&registry.Device{}, "id = ?", content)
		if errors.Is(manager.db.Error, gorm.ErrRecordNotFound) {
			return database.InsertError{ErrorMessage: "Device content id does not exist"}
		}

		// prevent circular references
		if content == device.Id {
			return database.InsertError{ErrorMessage: "Device cannot contain itself"}
		}
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
	var device registry.Device
	manager.db.First(&device, "id = ?", id)

	return nil, nil
}

func (manager *StorageManager) List() ([]*registry.Device, error) {
	// Todo
	return nil, nil
}

func (manager *StorageManager) Update(device *registry.Device) error {
	// Todo
	return nil
}

func (manager *StorageManager) Delete(id string) error {
	// Todo
	return nil
}
