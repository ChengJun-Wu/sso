package handlers

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

type Login struct {
	Handler
}

type LoginFrom struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

func (c *Login) Store(ctx *gin.Context)  {
	var form LoginFrom
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, helpers.ResponseFail(err.Error()))
		return
	}
	db := statics.GetDb()
	user := models.User{
		Username: form.Username,
	}
	result := db.Take(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.IncrFailedFrequency(ctx)
		ctx.JSON(http.StatusOK, helpers.ResponseFail("账号或密码不正确"))
		return
	}
	if !helpers.PasswordVerify(form.Password, user.Password) {
		c.IncrFailedFrequency(ctx)
		ctx.JSON(http.StatusOK, helpers.ResponseFail("账号或密码不正确"))
		return
	}
	session := sessions.Default(ctx)
	session.Set("uid", user.ID)
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}