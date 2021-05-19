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

func (c *Login) Index(ctx *gin.Context) {
	session := sessions.Default(ctx)
	uid := session.Get("uid")
	if uid == nil {
		ctx.JSON(http.StatusOK, helpers.ResponseNeedLogin())
		return
	}
	var user models.User
	db := statics.GetDb()
	result := db.Where("id", uid.(uint)).Take(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, helpers.ResponseNeedLogin())
		return
	}
	user.Password = ""
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess(user))
}

type LoginFrom struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

func (c *Login) Store(ctx *gin.Context) {
	var form LoginFrom
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, helpers.ResponseFail(err.Error()))
		return
	}
	db := statics.GetDb()
	var user models.User
	result := db.Where("username", form.Username).Take(&user)
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
	session.Save()
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}

func (c *Login) Destroy(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("uid")
	session.Save()
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}