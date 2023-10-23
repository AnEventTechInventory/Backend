package registry

import "gorm.io/gorm"

type Device struct {
	Id           string `json:"id" gorm:"primaryKey;type:string; not null"`
	Name         string `json:"name" gorm:"type: string; not null"`
	Manufacturer string `json:"manufacturer" gorm:"type:string;not null"`
	Description  string `json:"description" gorm:"type:string"`
	Quantity     int    `json:"quantity" gorm:"not null; check:quantity > 0"`
	Contents     string `json:"contents"`
	gorm.Model
}
