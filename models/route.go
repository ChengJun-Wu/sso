package models

import "gorm.io/gorm"

type Route struct {
	ID uint `gorm:"primarykey"`
	AppId uint `gorm:"index:idx_appId_path_method,unique"`
	Path string `gorm:"index:idx_appId_path_method,unique"`
	Method string `gorm:"index:idx_appId_path_method,unique"`
	Desc string
	Enable int
	gorm.Model
}
