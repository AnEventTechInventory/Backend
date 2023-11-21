package registry

import (
	"errors"
	"github.com/AnEventTechInventory/Backend/pkg/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type Device struct {
	ID          uuid.UUID `gorm:"index:;primaryKey;type:char(36)"`
	Name        string    `gorm:"not null; unique"`
	Description string

	Manufacturer *Manufacturer
	Location     *Location
	Type         *Type

	ManufacturerId string `gorm:"not null;type:char(36)"`
	LocationId     string `gorm:"not null;type:char(36)"`
	TypeId         string `gorm:"not null;type:char(36)"`
	Quantity       int    `gorm:"not null; check:quantity > 0"`
	Contents       string `gorm:"type:text"`
	gorm.Model
}

type JsonDevice struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Location     string   `json:"location"`
	Manufacturer string   `json:"manufacturer"`
	Type         string   `json:"type"`
	Quantity     int      `json:"quantity"`
	Contents     []string `json:"contents"`
}

func (device *Device) Validate(db *gorm.DB) error {
	if device.ID == uuid.Nil {
		return util.ErrMissingField("id")
	}
	if err := device.Manufacturer.Validate(db); err != nil {
		return err
	}
	if err := device.Location.Validate(db); err != nil {
		return err
	}
	if err := device.Type.Validate(db); err != nil {
		return err
	}
	if device.Name == "" {
		return util.ErrMissingField("name")
	}
	if device.Quantity < 0 {
		return errors.New("device quantity cannot be negative")
	}
	if err := device.VerifyContents(db); err != nil {
		return err
	}
	return nil
}

func DeviceFromJson(device JsonDevice, db *gorm.DB) (*Device, error) {
	// validate UUID
	id, err := uuid.Parse(device.ID)
	if err != nil {
		return nil, err
	}

	manId, err := uuid.Parse(device.Manufacturer)
	if err != nil {
		return nil, err
	}

	var man *Manufacturer

	db.First(man, "id = ?", manId)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("manufacturer id does not exist")
	}

	locId, err := uuid.Parse(device.Location)
	if err != nil {
		return nil, err
	}

	var loc *Location
	db.First(loc, "id = ?", locId)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("location id does not exist")
	}

	dev := &Device{
		ID:             id,
		Name:           device.Name,
		Description:    device.Description,
		ManufacturerId: device.Manufacturer,
		Manufacturer:   man,
		Location:       loc,
		LocationId:     device.Location,
		Quantity:       device.Quantity,
		Contents:       strings.Join(device.Contents, ","),
	}
	if err := dev.Validate(db); err != nil {
		return nil, err
	}
	return dev, nil
}

func (device *Device) VerifyContents(db *gorm.DB) error {
	// Verify that the device contents exists
	contents := strings.Split(device.Contents, ",")
	for _, content := range contents {
		if err := util.ValidateUUID(content); err != nil {
			return err
		}
		// check if the content ids already exist
		var contentDevice *Device
		db.Find(contentDevice, "id = ?", content)
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return errors.New("device content id does not exist")
		}

		if err := contentDevice.Validate(db); err != nil {
			return err
		}

		// prevent circular references
		if contentDevice.ID == device.ID {
			return errors.New("device cannot contain itself")
		}
	}
	return nil
}
