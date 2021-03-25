package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sso/helpers"
	"sso/models"
	"sso/statics"
)

type AuthRoute struct {
	Handler
}

func (h *AuthRoute) Index(ctx *gin.Context) {
	authId := ctx.Param("id")
	var (
		authRoutes []models.AuthRoute
		count int64
	)
	db := statics.GetDb()
	query := db.Model(&models.AuthRoute{})
	query.Where("auth_id = ?", authId)
	query.Count(&count)
	query.Offset(h.Offset(ctx)).Limit(h.Limit(ctx)).Find(&authRoutes)
	ctx.JSON(http.StatusOK, helpers.ResponseDivideData(authRoutes, count))
}

func (h *AuthRoute) Update(ctx *gin.Context) {
	authId := ctx.Param("id")
	routeId := ctx.Param("rid")
	db := statics.GetDb()
	authRoute := models.AuthRoute{
		AuthId: h.StringToUInt(authId),
		RouteId: h.StringToUInt(routeId),
	}
	db.Create(&authRoute)
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}

func (h *AuthRoute) Destroy(ctx *gin.Context) {
	authId := ctx.Param("id")
	routeId := ctx.Param("rid")
	db := statics.GetDb()
	db.Where("auth_id", authId).Where("route_id", routeId).Delete(&models.AuthRoute{})
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}