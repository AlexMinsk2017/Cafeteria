package models

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	ClientRefer     uint
	Client          Clients `gorm:"foreignKey:ClientRefer"`
	CourierRefer    uint
	Courier         Couriers `gorm:"foreignKey:CourierRefer"`
	RestaurantRefer uint
	Restaurant      Restaurants `gorm:"foreignKey:RestaurantRefer"`
	City            string      `gorm:"size:50"`
	Summa           float64
}
