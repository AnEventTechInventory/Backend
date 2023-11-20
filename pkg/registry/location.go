package registry

import (
	"github.com/AnEventTechInventory/Backend/pkg/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	entryInterface
	ID          uuid.UUID `json:"id" gorm:"primaryKey; not null;type:char(36)"`
	Name        string    `json:"name" gorm:"not null; unique"`
	Description string    `json:"description"`
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
