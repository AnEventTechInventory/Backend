package registry

import (
	"github.com/AnEventTechInventory/Backend/pkg/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type entryInterface interface {
	Validate(db *gorm.DB) error
}

type entry struct {
	entryInterface
	Id          uuid.UUID `json:"id" gorm:"primaryKey; not null"`
	Name        string    `json:"name" gorm:"not null; unique"`
	Description string    `json:"description"`
	gorm.Model
}

func (entr *entry) Validate(db *gorm.DB) error {
	if entr.Id == uuid.Nil {
		return util.ErrMissingField("id")
	}
	if entr.Name == "" {
		return util.ErrMissingField("name")
	}
	return nil
}
