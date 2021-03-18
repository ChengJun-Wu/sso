package commands

import (
	"github.com/gin-gonic/gin"
	"sso/models"
	"sso/router"
	"sso/statics"
)

type Route struct {
}

func (command *Route) Run()  {
	r := gin.Default()
	routeManager := router.RouteManager{}
	routeManager.Init(r)
	db := statics.GetDb()
	var enableIds []uint
	var dbRoute models.Route
	for _, route := range r.Routes(){
		dbRoute = models.Route{
			Path: route.Path,
			Method: route.Method,
		}
		db.FirstOrCreate(&dbRoute, models.Route{Path: route.Path, Method: route.Method})
		enableIds = append(enableIds, dbRoute.ID)
	}
	db.Model(models.Route{}).Where("id IN ?", enableIds).Updates(models.Route{Enable: 1})
	db.Model(models.Route{}).Where("id NOT IN ?", enableIds).Updates(models.Route{Enable: 0})
}
