package registry

import (
	"github.com/AnEventTechInventory/Backend/pkg/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Manufacturer struct {
	entryInterface
	ID          uuid.UUID `json:"id" gorm:"primaryKey; not null;type:char(36)"`
	Name        string    `json:"name" gorm:"not null; unique"`
	Description string    `json:"description"`
	gorm.Model
}

type JsonManufacturer BaseJson

func (manufacturer *Manufacturer) Validate(db *gorm.DB) error {
	if manufacturer.ID == uuid.Nil {
		return util.ErrMissingField("id")
	}
	if manufacturer.Name == "" {
		return util.ErrMissingField("name")
	}
	return nil
}
