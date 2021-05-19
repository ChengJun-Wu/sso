package middlewares

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sso/helpers"
	"sso/models"
	"sso/statics"
)

func LoginMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		uid := session.Get("uid")
		if uid == nil {
			ctx.JSON(http.StatusOK, helpers.ResponseNeedLogin())
			ctx.Abort()
			return
		}
		var user models.User
		db := statics.GetDb()
		result := db.Where("id", uid.(uint)).Take(&user)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusOK, helpers.ResponseNeedLogin())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}