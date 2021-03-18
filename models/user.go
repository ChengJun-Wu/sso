package models

import "gorm.io/gorm"

type User struct {
	ID uint `gorm:"primarykey"`
	Name string
	Username string
	Password string
	gorm.Model
}
