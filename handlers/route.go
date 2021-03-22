package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sso/helpers"
	"sso/models"
	"sso/statics"
)

type Route struct {
	Handler
}

type RouteIndexForm struct {
	AppId uint `form:"app_id" json:"app_id" xml:"app_id"`
	Path string `form:"path" json:"path" xml:"path"`
}

func (h *Route) Index(ctx *gin.Context) {
	var (
		form RouteIndexForm
		routes []models.Route
		count int64
	)
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, helpers.ResponseFail(err.Error()))
		return
	}
	db := statics.GetDb()
	query := db.Model(&models.Route{})
	if form.AppId != 0 {
		query.Where("app_id = ?", form.AppId)
	}
	if form.Path != "" {
		query.Where("path like ?", "%" + form.Path + "%")
	}
	query.Count(&count)
	query.Offset(h.Offset(ctx)).Limit(h.Limit(ctx)).Find(&routes)
	ctx.JSON(http.StatusOK, helpers.ResponseDivideData(routes, count))
}

func (h *Route) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	route := models.Route{
		ID: h.StringToUInt(id),
	}
	db := statics.GetDb()
	result := db.Take(&route)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
		return
	}
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess(route))
}

type RouteUpdateForm struct {
	Desc string
}

func (h *Route) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	route := models.Route{
		ID: h.StringToUInt(id),
	}
	db := statics.GetDb()
	result := db.Take(&route)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, helpers.ResponseFail("data not found"))
		return
	}
	var form RouteUpdateForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, helpers.ResponseFail(err.Error()))
		return
	}
	route.Desc = form.Desc
	db.Save(&route)
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}