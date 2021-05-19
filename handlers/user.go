package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sso/helpers"
	"sso/models"
	"sso/statics"
)

type User struct {
	Handler
}

type UserIndexForm struct {
	Name string `form:"auth_key" json:"auth_key" xml:"auth_key"`
	Username string `form:"name" json:"name" xml:"name"`
}

func (h *User) Index(ctx *gin.Context) {
	var (
		form UserIndexForm
		users []models.User
		count int64
	)
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, helpers.ResponseFail(err.Error()))
		return
	}
	db := statics.GetDb()
	query := db.Model(&models.User{})
	if form.Name != "" {
		query.Where("name like ?", "%" + form.Name + "%")
	}
	if form.Username != "" {
		query.Where("`username` like ?", "%" + form.Username + "%")
	}
	query.Count(&count)
	query.Offset(h.Offset(ctx)).Limit(h.Limit(ctx)).Order("id asc").Find(&users)
	for idx, _ := range users {
		users[idx].Password = ""
	}
	ctx.JSON(http.StatusOK, helpers.ResponseDivideData(users, count))
}

func (h *User) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	db := statics.GetDb()
	result := db.Where("id", h.StringToUInt(id)).Take(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
		return
	}
	user.Password = ""
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess(user))
}

type UserUpdateForm struct {
	Name string
	Password string
}

func (h *User) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	db := statics.GetDb()
	result := db.Where("id", h.StringToUInt(id)).Take(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, helpers.ResponseFail("data not found"))
		return
	}
	var form UserUpdateForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, helpers.ResponseFail(err.Error()))
		return
	}
	user.Name = form.Name
	if user.Password != "" {
		user.Password = helpers.PasswordHash(form.Password)
	}
	db.Save(&user)
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}

type UserStoreForm struct {
	Name string
	Username string
	Password string
}

func (h *User) Store(ctx *gin.Context) {
	var form UserStoreForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, helpers.ResponseFail(err.Error()))
		return
	}
	user := models.User{
		Name: form.Name,
		Username: form.Username,
		Password: helpers.PasswordHash(form.Password),
	}
	db := statics.GetDb()
	db.Create(&user)
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}