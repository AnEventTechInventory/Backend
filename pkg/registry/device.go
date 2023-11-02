package registry

import (
	"errors"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

type Device struct {
	Id           string `json:"id" gorm:"primaryKey; not null"`
	Name         string `json:"name" gorm:"not null"`
	Manufacturer string `json:"manufacturer" gorm:"not null"`
	Description  string `json:"description"`
	Quantity     int    `json:"quantity" gorm:"not null; check:quantity > 0"`
	Contents     string `json:"contents"`
	gorm.Model
}

func (device *Device) VerifyContents(db *gorm.DB) error {
	// verify that the device contents are valid
	// can only a list of valid uuids followed by exactly one ', ' or be the end
	match, err := regexp.Match(`^([a-f\d]{8}(-[a-f\d]{4}){4}[a-f\d]{8}, )*[a-f\d]{8}(-[a-f\d]{4}){4}[a-f\d]{8}$`, []byte(device.Contents))
	if err != nil {
		return err
	}
	if !match {
		return errors.New("device contents must be a list of valid uuids separated by exactly one ', '")
	}

	// Verify that the device contents exists
	for _, content := range strings.Split(device.Contents, ", ") {
		// check if the content ids already exist
		db.Find(&Device{}, "id = ?", content)
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return errors.New("device content id does not exist")
		}

		// prevent circular references
		if content == device.Id {
			return errors.New("device cannot contain itself")
		}
	}

	return nil
}

func (device *Device) Validate(db *gorm.DB) error {
	if device.Name == "" {
		return errors.New("device name cannot be empty")
	}
	if device.Quantity < 0 {
		return errors.New("device quantity cannot be negative")
	}
	if device.Manufacturer == "" {
		return errors.New("device manufacturer cannot be empty")
	}
	err := device.VerifyContents(db)
	if err != nil {
		return err
	}
	return nil
}
