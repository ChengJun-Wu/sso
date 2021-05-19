package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sso/helpers"
	"sso/models"
	"sso/statics"
)

type UserAuth struct {
	Handler
}

func (h *UserAuth) Index(ctx *gin.Context) {
	userId := ctx.Param("id")
	var (
		userAuth []models.UserAuth
		count int64
	)
	db := statics.GetDb()
	query := db.Model(&models.UserAuth{})
	query.Where("user_id = ?", userId)
	query.Count(&count)
	query.Offset(h.Offset(ctx)).Limit(h.Limit(ctx)).Find(&userAuth)
	ctx.JSON(http.StatusOK, helpers.ResponseDivideData(userAuth, count))
}

func (h *UserAuth) Update(ctx *gin.Context) {
	userId := ctx.Param("id")
	authId := ctx.Param("aid")
	db := statics.GetDb()
	userAuth := models.UserAuth{
		UserId: h.StringToUInt(userId),
		AuthId: h.StringToUInt(authId),
	}
	db.Create(&userAuth)
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}

func (h *UserAuth) Destroy(ctx *gin.Context) {
	userId := ctx.Param("id")
	authId := ctx.Param("aid")
	db := statics.GetDb()
	db.Where("user_id", userId).Where("auth_id", authId).Delete(&models.UserAuth{})
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}