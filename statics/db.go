package statics

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sso/models"
)

var db *gorm.DB

func init()  {
	config := GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Database.Username, config.Database.Password, config.Database.Ip, config.Database.Port, config.Database.Name)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = database
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)

	db.AutoMigrate(&models.User{})
}

func GetDb() *gorm.DB {
	return db
}