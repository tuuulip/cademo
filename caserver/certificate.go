package caserver

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"

	"github.com/hyperledger/fabric-ca/api"
	"github.com/pkg/errors"
)

// Get certificates
func GetCertificateList(req *api.GetCertificatesRequest) ([]x509.Certificate, error) {
	id, err := getAdminIdentity()
	if err != nil {
		return nil, err
	}
	certs := []x509.Certificate{}
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
		return nil
	})
	return certs, err
}
