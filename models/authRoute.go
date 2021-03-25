package models

import "gorm.io/gorm"

type AuthRoute struct {
	ID uint `gorm:"primarykey"`
	AuthId uint
	RouteId uint
	gorm.Model
}