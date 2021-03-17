package middlewares

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sso/helpers"
)

func CaptchaMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		frequency := session.Get(helpers.FailedFrequency())
		fmt.Println(frequency)
		if frequency != nil && frequency.(int) > 1 {
			ctx.JSON(http.StatusOK, helpers.ResponseNeedCaptcha())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}