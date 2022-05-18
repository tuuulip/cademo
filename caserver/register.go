package caserver

import (
	"cademo/config"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"

	"github.com/grantae/certinfo"
	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/lib"
	"github.com/pkg/errors"
)

func Register(req *api.RegistrationRequest) (string, error) {
	id, err := getAdminIdentity()
	if err != nil {
		return "", err
	}

	resp, err := id.Register(req)
	if err != nil {
		return "", err
	}
	logger.Infof("Register %s success.", req.Name)
	return resp.Secret, nil
}

// get all identities
func GetAllIdentities() ([]api.IdentityInfo, error) {
	id, err := getAdminIdentity()
	if err != nil {
		return nil, err
	}
	identities := []api.IdentityInfo{}
	err = id.GetAllIdentities("", func(d *json.Decoder) error {
		var id api.IdentityInfo
		err := d.Decode(&id)
		if err != nil {
			return err
		}
		identities = append(identities, id)
		return nil
	})
	return identities, err
}

// get all certificates
func GetAllCertificates() ([]x509.Certificate, []string, error) {
	id, err := getAdminIdentity()
	if err != nil {
		return nil, nil, err
	}
	req := &api.GetCertificatesRequest{}
	certs := []x509.Certificate{}
	certsDisplay := []string{}
	err = id.GetCertificates(req, func(d *json.Decoder) error {
		type certPEM struct {
			PEM string `db:"pem"`
		}
		var cert certPEM
		err := d.Decode(&cert)
		if err != nil {
			return err
		}
		block, rest := pem.Decode([]byte(cert.PEM))
		if block == nil || len(rest) > 0 {
			return errors.New("Certificate decoding error")
		}
		certificate, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return err
		}
		certs = append(certs, *certificate)

		result, err := certinfo.CertificateText(certificate)
		if err != nil {
			return err
		}
		certsDisplay = append(certsDisplay, result)
		return nil
	})
	return certs, certsDisplay, err
}

// load admin identity
func getAdminIdentity() (*lib.Identity, error) {
	homeDir := getAdminDir()
	caurl := getRegisterUrl()

	clientCfg := &lib.ClientConfig{
		URL: caurl,
		TLS: getClientTls(),
	}
	client := lib.Client{
		HomeDir: homeDir,
		Config:  clientCfg,
	}

	return client.LoadMyIdentity()
}

// combine register url
func getRegisterUrl() string {
	proto := "http"
	if config.C.GetBool("caserver.tls.enabled") {
		proto = "https"
	}
	return fmt.Sprintf(
		"%s://%s:%d",
		proto,
		config.C.GetString("caserver.host"),
		config.C.GetInt("caserver.port"),
	)
}
