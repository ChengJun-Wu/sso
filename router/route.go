package router

import (
	"github.com/gin-gonic/gin"
	"sso/handlers"
	"sso/middlewares"
)

type RouteManager struct {
	login handlers.Login
	captcha handlers.Captcha
	test handlers.Test
}

func (m *RouteManager) Init(g *gin.Engine) {
	g.GET("login", m.login.Index)
	g.POST("login", middlewares.CaptchaMiddleware(), m.login.Store)
	g.DELETE("login", m.login.Destroy)
	g.GET("captcha", m.captcha.Index)
	g.POST("captcha", m.captcha.Store)
}