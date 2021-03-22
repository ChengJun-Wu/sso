package models

import "gorm.io/gorm"

type App struct {
	ID uint `gorm:"primarykey"`
	Name string
	Domain string
	AppKey string
	AppSecret string
	gorm.Model
}
