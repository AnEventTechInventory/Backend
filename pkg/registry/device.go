package registry

import "gorm.io/gorm"

type Device struct {
	Id           string `json:"id" gorm:"primaryKey; not null"`
	Name         string `json:"name" gorm:"not null"`
	Manufacturer string `json:"manufacturer" gorm:"not null"`
	Description  string `json:"description"`
	Quantity     int    `json:"quantity" gorm:"not null; check:quantity > 0"`
	Contents     string `json:"contents"`
	gorm.Model
}
