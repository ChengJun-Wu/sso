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
	route handlers.Route
	app handlers.App
}

func (m *RouteManager) Init(g *gin.Engine) {

	common := g.Group("/")
	{
		common.GET("login", m.login.Index)
		common.POST("login", middlewares.CaptchaMiddleware(), m.login.Store)
		common.DELETE("login", m.login.Destroy)
		common.GET("captcha", m.captcha.Index)
		common.POST("captcha", m.captcha.Store)
	}

	backend := g.Group("backend")
	backend.Use(middlewares.LoginMiddleware())
	{
		backend.GET("route", m.route.Index)
		backend.GET("route/:id", m.route.Show)
		backend.PUT("route/:id", m.route.Update)

		backend.GET("app", m.app.Index)
	}

}