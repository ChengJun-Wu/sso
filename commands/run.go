package commands

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sso/helpers"
	"sso/statics"
	"sso/models"
	"sso/router"
)

type Run struct {
}

func (command *Run) Run()  {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("sso-session", store))

	routeManager := router.RouteManager{}
	routeManager.Init(r)
	go checkInit()
	r.Run(":19284") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func checkInit()  {
	db := statics.GetDb()
	var user models.User
	result := db.Where("name", "admin").Take(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		user = models.User{
			Name: "admin",
			Username: "admin",
			Password: helpers.PasswordHash("admin"),
		}
		db.Create(&user)
	}
	var count int64
	db.Model(models.App{}).Count(&count)
	if count == 0 {
		app := models.App{
			Name: "sso",
			AppKey: helpers.RandomString(16),
			AppSecret: helpers.RandomString(16),
		}
		db.Create(&app)
	}
}