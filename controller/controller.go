package controller

import (
	"cademo/caserver"
	"cademo/message"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-ca/api"
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
	req := &message.Enroll{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}

	resp, err := caserver.Enroll(req)
	if err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, resp)
}

func (c *Controller) EnrollTls(ctx *gin.Context) {
	req := &message.Enroll{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}

	resp, err := caserver.EnrollTLS(req)
	if err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, resp)
}

func (c *Controller) Register(ctx *gin.Context) {
	req := &api.RegistrationRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	passwd, err := caserver.Register(req)
	if err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, passwd)
}

func (c *Controller) AllIdentities(ctx *gin.Context) {
	ids, err := caserver.GetAllIdentities()
	if err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, ids)
}

func (c *Controller) AllCertificates(ctx *gin.Context) {
	certs, displays, err := caserver.GetAllCertificates()
	_ = displays
	if err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, certs)
}
