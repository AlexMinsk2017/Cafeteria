package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	User     string `gorm:"size:25"`
	Password string `gorm:"size:50"`
}
