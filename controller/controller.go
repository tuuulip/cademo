package controller

import (
	"cademo/caserver"
	"cademo/db"
	"cademo/message"
	"crypto/x509"
	"encoding/hex"

	"github.com/cloudflare/cfssl/helpers"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/lib"
	"github.com/hyperledger/fabric-ca/util"
)

type Controller struct {
	caServer *lib.Server
	dbClient *db.DBClient
}

func NewController(caServer *lib.Server) *Controller {
	dbClinet, err := db.NewDBClient()
	if err != nil {
		panic(err)
	}
	return &Controller{
		caServer: caServer,
		dbClient: dbClinet,
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

func (c *Controller) DeleteCertificate(ctx *gin.Context) {
	req := &db.Certificates{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}

	if err := c.dbClient.DeleteCertificate(req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}

	ResponseSuccess(ctx, "ok")
}

func (c *Controller) parseCertificate(raw *x509.Certificate) *message.Certificate {
	pem := helpers.EncodeCertificatePEM(raw)
	cert := &message.Certificate{
		Id:             raw.Subject.CommonName,
		AuthorityKeyId: hex.EncodeToString(raw.AuthorityKeyId),
		SerialNumber:   util.GetSerialAsHex(raw.SerialNumber),
		Pem:            string(pem),
		NotBefore:      raw.NotBefore,
		NotAfter:       raw.NotAfter,
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

func (c *Controller) AddAffiliation(ctx *gin.Context) {
	req := &api.AddAffiliationRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	if err := caserver.AddAffiliation(req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, "ok")

}

func (c *Controller) DelAffiliation(ctx *gin.Context) {
	req := &api.RemoveAffiliationRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	if err := caserver.DeleteAffiliation(req); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, "ok")

}

func (c *Controller) GetUserState(ctx *gin.Context) {
	states, err := c.dbClient.QueryIdentityStates()
	if err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, states)
}

func (c *Controller) GetCredential(ctx *gin.Context) {
	cred, err := c.dbClient.QueryCredentials()
	if err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, cred)
}

func (c *Controller) DeleteCredential(ctx *gin.Context) {
	cred := &db.Credentials{}
	if err := ctx.ShouldBindJSON(cred); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	if err := c.dbClient.DeleteCredential(cred); err != nil {
		ResponseFail(ctx, err.Error())
		return
	}
	ResponseSuccess(ctx, "ok")
}
