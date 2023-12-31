package registry

import (
	"github.com/AnEventTechInventory/Backend/pkg/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	ID          uuid.UUID `gorm:"primaryKey; not null;type:char(36)"`
	Name        string    `gorm:"not null; unique"`
	Description string
	gorm.Model
}

type JsonLocation BaseJson

func (location *Location) Validate(db *gorm.DB) error {
	if location.ID == uuid.Nil {
		return util.ErrMissingField("id")
	}
	if location.Name == "" {
		return util.ErrMissingField("name")
	}
	return nil
}
