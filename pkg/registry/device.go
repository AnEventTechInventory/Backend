package registry

import "gorm.io/gorm"

type Device struct {
	Id           string   `json:"id" gorm:"primaryKey;type:uuid;"`
	Name         string   `json:"name" gorm:"not null"`
	Manufacturer string   `json:"manufacturer" gorm:"type:uuid;not null"`
	Description  string   `json:"description"`
	Quantity     int      `json:"quantity" gorm:"not null; check:quantity > 0"`
	Contents     []string `json:"contents"`
	gorm.Model
}
