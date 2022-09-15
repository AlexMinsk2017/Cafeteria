package models

import "gorm.io/gorm"

type Couriers struct {
	gorm.Model
	Name      string `gorm:"size:50"`
	Telephone string `gorm:"size:25"`
	City      string `gorm:"size:50"`
}
