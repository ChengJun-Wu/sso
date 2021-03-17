package handlers

import (
	"bytes"
	"encoding/base64"
	"github.com/afocus/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"image/png"
	"net/http"
	"sso/helpers"
)

type Captcha struct {
	Handler
}

func (c *Captcha) Index(ctx *gin.Context) {
	cap := captcha.New()
	cap.SetSize(164, 64)
	cap.SetFont("assets/comic.ttf")
	img, code := cap.Create(6, captcha.NUM)

	session := sessions.Default(ctx)
	session.Set("captcha", code)
	session.Save()

	buff := bytes.NewBuffer(nil)
	png.Encode(buff, img)
	str := base64.StdEncoding.EncodeToString(buff.Bytes())
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess(str))
}

type captchaForm struct {
	Captcha string `form:"captcha" json:"captcha" xml:"captcha"  binding:"required"`
}

func (c *Captcha) Store(ctx *gin.Context) {
	var form captchaForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, helpers.ResponseFail(err.Error()))
		return
	}
	session := sessions.Default(ctx)
	captcha := session.Get("captcha")
	if captcha != form.Captcha {
		ctx.JSON(http.StatusOK, helpers.ResponseFail("验证码不正确"))
		return
	}
	session.Delete("captcha")
	c.DecrFailedFrequency(ctx)
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}