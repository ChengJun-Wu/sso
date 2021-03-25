package models

import "gorm.io/gorm"

type Auth struct {
	ID uint `gorm:"primarykey"`
	AuthKey string
	Name string
	Desc string
	gorm.Model
}