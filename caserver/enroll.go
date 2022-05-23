package caserver

import (
	"cademo/config"
	"cademo/message"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/lib"
)

func Enroll(req *message.Enroll) (*lib.EnrollmentResponse, error) {
	identity, err := GetIdentity(req.User)
	if err != nil {
		return nil, err
	}

	csr := clientCSR(identity.Affiliation, req.User, identity.Type)
	cfg := &lib.ClientConfig{
		CSR: *csr,
		TLS: getClientTls(),
		Enrollment: api.EnrollmentRequest{
			Type: "idemix",
		},
	}
	enrollUrl := getEnrollUrl(req.User, defaultIdentityPassword(req.User))
	saveDir := getStoreHome(identity)
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
	if identity.Type == "peer" || identity.Type == "orderer" {
		_, err := EnrollTLS(req)
		if err != nil {
			return nil, err
		}
	}
	if err := storeTlsCA(saveDir); err != nil {
		return nil, err
	}
	return resp, err
}

func ReEnroll(req *message.Enroll) (*lib.EnrollmentResponse, error) {
	identity, err := GetIdentity(req.User)
	if err != nil {
		return nil, err
	}

	home := getStoreHome(identity)
	id, err := loadIdentity(home)
	if err != nil {
		return nil, err
	}
	cfg := id.GetClient().Config
	csr := clientCSR(identity.Affiliation, req.User, identity.Type)
	enrollReq := &api.ReenrollmentRequest{
		CSR: csr,
	}
	resp, err := id.Reenroll(enrollReq)
	if err := storeEnrollment(cfg, resp); err != nil {
		return nil, err
	}
	if err := storeTlsCA(home); err != nil {
		return nil, err
	}
	return resp, err
}

func EnrollTLS(req *message.Enroll) (*lib.EnrollmentResponse, error) {
	identity, err := GetIdentity(req.User)
	if err != nil {
		return nil, err
	}

	host := fmt.Sprintf("%s-%s", req.User, identity.Affiliation)
	cfg := &lib.ClientConfig{
		TLS:    getClientTls(),
		MSPDir: "tls",
		Enrollment: api.EnrollmentRequest{
			Profile: "tls",
		},
		CSR: api.CSRInfo{
			Hosts: []string{host},
		},
	}
	enrollUrl := getEnrollUrl(req.User, defaultIdentityPassword(req.User))
	saveDir := getStoreHome(identity)
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

	_, err := os.Stat(filepath.Join(adminDir, "msp", "signcerts"))
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

// certificate store home
func getStoreHome(identity *api.IdentityInfo) string {
	domain := identity.Affiliation + ".cademo.com"
	storeLocation := filepath.Join(
		getHomeDir(),
		"organizations",
		domain,
		identity.Type+"s",
	)
	lastDir := ""
	switch identity.Type {
	case "peer", "orderer":
		lastDir = identity.ID + "." + domain
	case "user":
		lastDir = identity.ID + "@" + domain
	default:
		lastDir = identity.Type + "_" + identity.ID
	}
	storeLocation = filepath.Join(storeLocation, lastDir)
	return storeLocation
}
