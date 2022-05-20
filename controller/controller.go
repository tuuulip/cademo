package controller

import (
	"cademo/caserver"
	"cademo/message"
	"crypto/x509"

	"github.com/cloudflare/cfssl/helpers"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/lib"
	"github.com/hyperledger/fabric-ca/util"
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

func (c *Controller) ReEnroll(ctx *gin.Context) {
	req := &message.Enroll{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}

	resp, err := caserver.ReEnroll(req)
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

func (c *Controller) DeleteIdentity(ctx *gin.Context) {
	req := &api.RemoveIdentityRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	if err := caserver.DeleteIdentity(req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, "ok")
}

func (c *Controller) RevokeIdentity(ctx *gin.Context) {
	req := &api.RevocationRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	if err := caserver.RevokeIdentity(req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, "ok")
}

func (c *Controller) CertificateList(ctx *gin.Context) {
	req := &api.GetCertificatesRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}

	certs, err := caserver.GetCertificateList(req)
	if err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	msgCerts := []message.Certificate{}
	for _, cert := range certs {
		msgCert := c.parseCertificate(&cert)
		msgCerts = append(msgCerts, *msgCert)
	}
	ResponseSuccess(ctx, msgCerts)
}

func (c *Controller) parseCertificate(raw *x509.Certificate) *message.Certificate {
	pem := helpers.EncodeCertificatePEM(raw)
	cert := &message.Certificate{
		Id:           raw.Subject.CommonName,
		SerialNumber: util.GetSerialAsHex(raw.SerialNumber),
		Pem:          string(pem),
		NotBefore:    raw.NotBefore,
		NotAfter:     raw.NotAfter,
	}
	return cert
}

func (c *Controller) AllAffiliations(ctx *gin.Context) {
	info, err := caserver.AllAffiliations()
	if err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, info)
}
