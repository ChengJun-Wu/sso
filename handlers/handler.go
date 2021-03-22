package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sso/helpers"
	"strconv"
)

type HandlerInterface interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Destroy(ctx *gin.Context)
}

type Handler struct {
}

type Form struct {
}

func (h *Handler) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}

func (h *Handler) Show(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}

func (h *Handler) Store(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}

func (h *Handler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}

func (h *Handler) Destroy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess())
}

func (h *Handler) UidString(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	return strconv.Itoa(session.Get("uid").(int))
}

func (h *Handler) IncrFailedFrequency(ctx *gin.Context) {
	session := sessions.Default(ctx)
	key := helpers.FailedFrequency()
	var count int
	v := session.Get(key)
	if v == nil {
		count = 1
	} else {
		count = v.(int)
		count++
	}
	session.Set(key, count)
	session.Save()
}

func (h *Handler) DecrFailedFrequency(ctx *gin.Context) {
	session := sessions.Default(ctx)
	key := helpers.FailedFrequency()
	var count int
	v := session.Get(key)
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		if count > 0 {
			count--
		}
	}
	session.Set(key, count)
	session.Save()
}

func (h *Handler) Page(ctx *gin.Context) int {
	page := ctx.DefaultQuery("page", "1")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	return pageInt
}

func (h *Handler) Limit(ctx *gin.Context) int {
	limit := ctx.DefaultQuery("limit", "10")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}
	return limitInt
}

func (h *Handler) Offset(ctx *gin.Context) int {
	return (h.Page(ctx) - 1) * h.Limit(ctx)
}

func (h *Handler) StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func (h *Handler) StringToUInt(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}