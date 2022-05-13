package controller

import (
	"cademo/caclient"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-ca/lib"
)

type Controller struct {
	caServer *lib.Server
}

func NewController(caServer *lib.Server) *Controller {
	return &Controller{
		caServer: caServer,
	}
}

func (c *Controller) CaInfo(ctx *gin.Context) {
	ca, err := c.caServer.GetCA("")
	if err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, ca)
}

func (c *Controller) Enroll(ctx *gin.Context) {
	ResponseSuccess(ctx, "ok")
}

func (c *Controller) Register(ctx *gin.Context) {
	if err := caclient.Register(); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, "ok")
}
