package registry

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Manufacturer struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique" gorm:"not null"`
	Description string `json:"description"`
	gorm.Model
}

func (manufacturer *Manufacturer) Validate() error {
	if manufacturer.Name == "" {
		return errors.New("manufacturer name cannot be empty")
	}

	// verify uuid is valid
	if _, err := uuid.Parse(manufacturer.Id); err != nil {
		return errors.New("manufacturer id is not a valid uuid")
	}

	return nil
}
