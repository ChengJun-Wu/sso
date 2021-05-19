package router

import (
	"github.com/gin-gonic/gin"
	"sso/handlers"
	"sso/middlewares"
)

type RouteManager struct {
	login handlers.Login
	captcha handlers.Captcha
	route handlers.Route
	auth handlers.Auth
	authRoute handlers.AuthRoute
	user handlers.User
	userAuth handlers.UserAuth
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

		backend.GET("auth", m.auth.Index)
		backend.GET("auth/:id", m.auth.Show)
		backend.PUT("auth/:id", m.auth.Update)
		backend.POST("auth", m.auth.Store)

		backend.GET("auth/:id/route", m.authRoute.Index)
		//backend.GET("auth/:aid/route/:rid", m.authRoute.Show)
		backend.PUT("auth/:id/route/:rid", m.authRoute.Update)
		//backend.POST("auth/:aid/route", m.authRoute.Store)
		backend.DELETE("auth/:id/route/:rid", m.authRoute.Destroy)

		backend.GET("user", m.user.Index)
		backend.GET("user/:id", m.user.Show)
		backend.PUT("user/:id", m.user.Update)
		backend.POST("user", m.user.Store)
		backend.GET("user/:id/auth", m.userAuth.Index)
		//backend.GET("auth/:aid/route/:rid", m.authRoute.Show)
		backend.PUT("user/:id/auth/:aid", m.userAuth.Update)
		//backend.POST("auth/:aid/route", m.authRoute.Store)
		backend.DELETE("user/:id/auth/:aid", m.userAuth.Destroy)

		backend.GET("app", m.app.Index)

	}

}