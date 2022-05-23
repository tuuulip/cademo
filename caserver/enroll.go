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
	csr := clientCSR(req.Organization, req.User, req.Type)
	cfg := &lib.ClientConfig{
		CSR: *csr,
		TLS: getClientTls(),
	}
	enrollUrl := getEnrollUrl(req.User, defaultIdentityPassword(req.User))
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

func ReEnroll(req *message.Enroll) (*lib.EnrollmentResponse, error) {
	home := filepath.Join(getHomeDir(), req.Type, req.User)
	id, err := loadIdentity(home)
	if err != nil {
		return nil, err
	}
	cfg := id.GetClient().Config
	csr := clientCSR(req.Organization, req.User, req.Type)
	enrollReq := &api.ReenrollmentRequest{
		CSR: csr,
	}
	resp, err := id.Reenroll(enrollReq)
	if err := storeEnrollment(cfg, resp); err != nil {
		return nil, err
	}
	saveDir := filepath.Join(getHomeDir(), req.Type, req.User)
	if err := storeTlsCA(saveDir); err != nil {
		return nil, err
	}
	return resp, err
}

func EnrollTLS(req *message.Enroll) (*lib.EnrollmentResponse, error) {
	host := fmt.Sprintf("%s-%s", req.User, req.Organization)
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
	enrollUrl := getEnrollUrl(req.User, defaultIdentityPassword(req.User))
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
