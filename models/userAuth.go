package models

import "gorm.io/gorm"

type UserAuth struct {
	ID uint `gorm:"primarykey"`
	UserId uint
	AuthId uint
	gorm.Model
}