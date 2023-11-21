package registry

import (
	"github.com/AnEventTechInventory/Backend/pkg/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Type struct {
	ID          uuid.UUID `gorm:"primaryKey; not null;type:char(36)"`
	Name        string    `gorm:"not null; unique"`
	Description string
	gorm.Model
}

func (t *Type) BeforeCreate(db *gorm.DB) (err error) {
	t.ID = uuid.New()
	return t.Validate(db)
}

type JsonType BaseJson

func (t *Type) Validate(db *gorm.DB) error {
	if t.ID == uuid.Nil {
		return util.ErrMissingField("id")
	}
	if t.Name == "" {
		return util.ErrMissingField("name")
	}
	return nil
}
