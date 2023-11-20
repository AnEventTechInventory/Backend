package registry

import (
	"gorm.io/gorm"
)

type entryInterface interface {
	Validate(db *gorm.DB) error
}

type BaseJson struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
