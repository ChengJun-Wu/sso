package models

import "gorm.io/gorm"

type Route struct {
	ID uint `gorm:"primarykey"`
	Path string `gorm:"index:idx_path_method,unique"`
	Method string `gorm:"index:idx_path_method,unique"`
	Desc string
	Enable int
	gorm.Model
}
