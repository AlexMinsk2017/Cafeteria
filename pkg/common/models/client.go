package models

import "gorm.io/gorm"

type Clients struct {
	gorm.Model
	Name      string `gorm:"size:50"`
	Telephone string `gorm:"size:25"`
	Mail      string `gorm:"size:50"`
	City      string `gorm:"size:50"`
}
