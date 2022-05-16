package caserver

import (
	"bytes"
	"cademo/config"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/hyperledger/fabric-ca/lib"
	"github.com/hyperledger/fabric-ca/util"
	"github.com/pkg/errors"
)

func Enroll(user, pass string) (*lib.EnrollmentResponse, error) {
	adminDir := getAdminDir()
	cfg := &lib.ClientConfig{}
	enrollUrl := fmt.Sprintf(
		"http://%s:%s@%s:%d",
		user, pass,
		config.C.GetString("caserver.host"),
		config.C.GetInt("caserver.port"),
	)
	resp, err := cfg.Enroll(enrollUrl, adminDir)
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

	cfg := &lib.ClientConfig{}
	enrollUrl := fmt.Sprintf(
		"http://%s:%s@%s:%d",
		config.C.GetString("caadmin.user"),
		config.C.GetString("caadmin.pass"),
		config.C.GetString("caserver.host"),
		config.C.GetInt("caserver.port"),
	)
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
