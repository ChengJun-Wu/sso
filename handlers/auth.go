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

type Auth struct {
	Handler
}

type AuthIndexForm struct {
	AuthKey string `form:"auth_key" json:"auth_key" xml:"auth_key"`
	Name string `form:"name" json:"name" xml:"name"`
	Desc string `form:"desc" json:"desc" xml:"desc"`
}

func (h *Auth) Index(ctx *gin.Context) {
	var (
		form AuthIndexForm
		auths []models.Auth
		count int64
	)
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, helpers.ResponseFail(err.Error()))
		return
	}
	db := statics.GetDb()
	query := db.Model(&models.Auth{})
	if form.AuthKey != "" {
		query.Where("auth_key like ?", "%" + form.AuthKey + "%")
	}
	if form.Name != "" {
		query.Where("`name` like ?", "%" + form.Name + "%")
	}
	if form.Desc != "" {
		query.Where("`desc` like ?", "%" + form.Desc + "%")
	}
	query.Count(&count)
	query.Offset(h.Offset(ctx)).Limit(h.Limit(ctx)).Order("id asc").Find(&auths)
	ctx.JSON(http.StatusOK, helpers.ResponseDivideData(auths, count))
}

func (h *Auth) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	auth := models.Auth{
		ID: h.StringToUInt(id),
	}
	db := statics.GetDb()
	result := db.Take(&auth)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
		return
	}
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess(auth))
}

type AuthUpdateForm struct {
	AuthKey string
	Name string
	Desc string
}

func (h *Auth) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	auth := models.Auth{
		ID: h.StringToUInt(id),
	}
	db := statics.GetDb()
	result := db.Take(&auth)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, helpers.ResponseFail("data not found"))
		return
	}
	var form AuthUpdateForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, helpers.ResponseFail(err.Error()))
		return
	}
	auth.AuthKey = form.AuthKey
	auth.Name = form.Name
	auth.Desc = form.Desc
	db.Save(&auth)
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}

type AuthStoreForm struct {
	AuthKey string
	Name string
	Desc string
}

func (h *Auth) Store(ctx *gin.Context) {
	var form AuthStoreForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, helpers.ResponseFail(err.Error()))
		return
	}
	auth := models.Auth{
		AuthKey: form.AuthKey,
		Name: form.Name,
		Desc: form.Desc,
	}
	db := statics.GetDb()
	db.Create(&auth)
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}