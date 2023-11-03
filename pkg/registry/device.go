package registry

import (
	"errors"
	"github.com/AnEventTechInventory/Backend/pkg/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Device struct {
	entry
	Manufacturer Manufacturer `json:"manufacturer" gorm:"not null; foreignKey:Id"`
	Location     Location     `json:"location" gorm:"not null; foreignKey:Id"`
	Quantity     int          `json:"quantity" gorm:"not null; check:quantity > 0"`
	Contents     []*Device    `json:"contents" gorm:"many2many:device_contents;"`
	gorm.Model
	entryInterface
}

func (device *Device) VerifyContents(db *gorm.DB) error {
	// Verify that the device contents exists
	for _, content := range device.Contents {
		// check if the content ids already exist
		db.Find(&Device{}, "id = ?", content.Id)
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return errors.New("device content id does not exist")
		}

		if err := content.Validate(db); err != nil {
			return err
		}

		// prevent circular references
		if content.Id == device.Id {
			return errors.New("device cannot contain itself")
		}
	}

	return nil
}

func (device *Device) Validate(db *gorm.DB) error {
	if device.Id == uuid.Nil {
		return util.ErrMissingField("id")
	}
	if err := device.Manufacturer.Validate(db); err != nil {
		return err
	}
	if err := device.Location.Validate(db); err != nil {
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
