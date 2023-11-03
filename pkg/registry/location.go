package registry

import (
	"gorm.io/gorm"
)

type Location struct {
	entry
	gorm.Model
}
