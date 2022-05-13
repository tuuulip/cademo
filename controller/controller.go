package controller

import "github.com/gin-gonic/gin"

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Enroll(ctx *gin.Context) {
	ResponseSuccess(ctx, "ok")
}
