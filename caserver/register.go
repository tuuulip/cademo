package caserver

import (
	"cademo/config"
	"fmt"

	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/lib"
	"github.com/hyperledger/fabric-ca/lib/tls"
)

func Register(req *api.RegistrationRequest) (string, error) {
	homeDir := getAdminDir()
	caurl := getRegisterUrl()

	tls := tls.ClientTLSConfig{
		Enabled:   config.C.GetBool("caserver.tls.enabled"),
		CertFiles: []string{"../tls-cert.pem"},
	}
	clientCfg := &lib.ClientConfig{
		URL: caurl,
		TLS: tls,
	}
	client := lib.Client{
		HomeDir: homeDir,
		Config:  clientCfg,
	}

	id, err := client.LoadMyIdentity()
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
