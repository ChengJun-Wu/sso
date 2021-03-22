package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sso/helpers"
	"sso/models"
	"sso/statics"
)

type App struct {
	Handler
}

func (h *App) Index(ctx *gin.Context) {
	var (
		apps []models.App
		count int64
	)
	db := statics.GetDb()
	query := db.Model(models.App{})
	query.Count(&count)
	query.Offset(h.Offset(ctx)).Limit(h.Limit(ctx)).Find(&apps)
	ctx.JSON(http.StatusOK, helpers.ResponseDivideData(apps, count))
}
