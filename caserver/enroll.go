package caserver

import (
	"bytes"
	"cademo/config"
	"cademo/message"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/lib"
	"github.com/hyperledger/fabric-ca/util"
	"github.com/pkg/errors"
)

func Enroll(req *message.Enroll) (*lib.EnrollmentResponse, error) {
	cfg := &lib.ClientConfig{
		TLS: getClientTls(),
	}
	enrollUrl := getEnrollUrl(req.User, req.Password)
	saveDir := filepath.Join(getHomeDir(), req.Type, req.User)
	resp, err := cfg.Enroll(enrollUrl, saveDir)
	if err != nil {
		return nil, err
	}
	if err := storeEnrollment(cfg, resp); err != nil {
		return nil, err
	}
	if err := storeTlsCA(saveDir); err != nil {
		return nil, err
	}
	return resp, err
}

func EnrollTLS(req *message.Enroll) (*lib.EnrollmentResponse, error) {
	host := fmt.Sprintf("%s-%s", req.User, req.Org)
	cfg := &lib.ClientConfig{
		TLS:    getClientTls(),
		MSPDir: "tls-msp",
		Enrollment: api.EnrollmentRequest{
			Profile: "tls",
		},
		CSR: api.CSRInfo{
			Hosts: []string{host},
		},
	}
	enrollUrl := getEnrollUrl(req.User, req.Password)
	saveDir := filepath.Join(getHomeDir(), req.Type, req.User)
	resp, err := cfg.Enroll(enrollUrl, saveDir)
	if err != nil {
		return nil, err
	}
	if err := storeEnrollment(cfg, resp); err != nil {
		return nil, err
	}
	return resp, err
}

// Enroll admin at server start.
func EnrollAdmin() error {
	adminDir := getAdminDir()

	_, err := os.Stat(adminDir)
	if err == nil {
		logger.Info("Admin already enrolled.")
		return nil
	}

	cfg := &lib.ClientConfig{
		TLS: getClientTls(),
	}
	enrollUrl := getEnrollUrl(config.C.GetString("caadmin.user"), config.C.GetString("caadmin.pass"))
	resp, err := cfg.Enroll(enrollUrl, adminDir)
	if err != nil {
		return err
	}
	if err := storeEnrollment(cfg, resp); err != nil {
		return err
	}
	logger.Info("Enroll admin success.")
	return nil
}

// get enroll url
func getEnrollUrl(user, pass string) string {
	host := config.C.GetString("caserver.host")
	port := config.C.GetInt("caserver.port")
	proto := "http"
	if config.C.GetBool("caserver.tls.enabled") {
		proto = "https"
	}
	url := fmt.Sprintf("%s://%s:%s@%s:%d", proto, user, pass, host, port)
	return url
}

// Store enrollment info to disk
func storeEnrollment(cfg *lib.ClientConfig, enrollment *lib.EnrollmentResponse) error {
	if err := enrollment.Identity.Store(); err != nil {
		return err
	}
	if err := storeCAChain(cfg, &enrollment.CAInfo); err != nil {
		return err
	}
	if err := storeIssuerPublicKey(cfg, &enrollment.CAInfo); err != nil {
		return err
	}
	if err := storeIssuerRevocationPublicKey(cfg, &enrollment.CAInfo); err != nil {
		return err
	}
	return nil
}

// Store tls ca cert
func storeTlsCA(home string) error {
	tlsCaPath := filepath.Join(getHomeDir(), "ca-cert.pem")
	tlsCaCert, err := ioutil.ReadFile(tlsCaPath)
	if err != nil {
		return err
	}
	dir := filepath.Join(home, "msp", "tlscacerts")
	return storeToFile("TLS Ca Cert", dir, "ca.crt", tlsCaCert)
}

// Store the CAChain in the CACerts folder of MSP (Membership Service Provider)
// The root cert in the chain goes into MSP 'cacerts' directory.
// The others (if any) go into the MSP 'intermediatecerts' directory.
func storeCAChain(config *lib.ClientConfig, si *lib.GetCAInfoResponse) error {
	mspDir := config.MSPDir
	// Get a unique name to use for filenames
	serverURL, err := url.Parse(config.URL)
	if err != nil {
		return err
	}
	fname := serverURL.Host
	if config.CAName != "" {
		fname = fmt.Sprintf("%s-%s", fname, config.CAName)
	}
	fname = strings.Replace(fname, ":", "-", -1)
	fname = strings.Replace(fname, ".", "-", -1) + ".pem"
	tlsfname := fmt.Sprintf("tls-%s", fname)

	rootCACertsDir := path.Join(mspDir, "cacerts")
	intCACertsDir := path.Join(mspDir, "intermediatecerts")
	tlsRootCACertsDir := path.Join(mspDir, "tlscacerts")
	tlsIntCACertsDir := path.Join(mspDir, "tlsintermediatecerts")

	var rootBlks [][]byte
	var intBlks [][]byte
	chain := si.CAChain
	for len(chain) > 0 {
		var block *pem.Block
		block, chain = pem.Decode(chain)
		if block == nil {
			break
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return errors.Wrap(err, "Failed to parse certificate in the CA chain")
		}

		if !cert.IsCA {
			return errors.New("A certificate in the CA chain is not a CA certificate")
		}

		// If authority key id is not present or if it is present and equal to subject key id,
		// then it is a root certificate
		if len(cert.AuthorityKeyId) == 0 || bytes.Equal(cert.AuthorityKeyId, cert.SubjectKeyId) {
			rootBlks = append(rootBlks, pem.EncodeToMemory(block))
		} else {
			intBlks = append(intBlks, pem.EncodeToMemory(block))
		}
	}

	// Store the root certificates in the "cacerts" msp folder
	certBytes := bytes.Join(rootBlks, []byte(""))
	if len(certBytes) > 0 {
		if config.Enrollment.Profile == "tls" {
			err := storeToFile("TLS root CA certificate", tlsRootCACertsDir, tlsfname, certBytes)
			if err != nil {
				return err
			}
		} else {
			err = storeToFile("root CA certificate", rootCACertsDir, fname, certBytes)
			if err != nil {
				return err
			}
		}
	}

	// Store the intermediate certificates in the "intermediatecerts" msp folder
	certBytes = bytes.Join(intBlks, []byte(""))
	if len(certBytes) > 0 {
		if config.Enrollment.Profile == "tls" {
			err = storeToFile("TLS intermediate certificates", tlsIntCACertsDir, tlsfname, certBytes)
			if err != nil {
				return err
			}
		} else {
			err = storeToFile("intermediate CA certificates", intCACertsDir, fname, certBytes)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func storeIssuerPublicKey(config *lib.ClientConfig, si *lib.GetCAInfoResponse) error {
	if len(si.IssuerPublicKey) > 0 {
		err := storeToFile("Issuer public key", config.MSPDir, "IssuerPublicKey", si.IssuerPublicKey)
		if err != nil {
			return err
		}
	}
	return nil
}

func storeIssuerRevocationPublicKey(config *lib.ClientConfig, si *lib.GetCAInfoResponse) error {
	if len(si.IssuerRevocationPublicKey) > 0 {
		err := storeToFile("Issuer revocation public key", config.MSPDir, "IssuerRevocationPublicKey", si.IssuerRevocationPublicKey)
		if err != nil {
			return err
		}
	}
	return nil
}

func storeToFile(what, dir, fname string, contents []byte) error {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return errors.Wrapf(err, "Failed to create directory for %s at '%s'", what, dir)
	}
	fpath := path.Join(dir, fname)
	err = util.WriteFile(fpath, contents, 0644)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("Failed to store %s at '%s'", what, fpath))
	}
	fmt.Printf("Stored %s at %s", what, fpath)
	return nil
}
