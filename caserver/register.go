package caserver

import (
	"cademo/config"
	"fmt"

	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/lib"
)

func Register() (string, error) {
	homeDir := getAdminDir()
	caurl := fmt.Sprintf(
		"http://%s:%d",
		config.C.GetString("caserver.host"),
		config.C.GetInt("caserver.port"),
	)
	clientCfg := &lib.ClientConfig{
		URL: caurl,
	}
	client := lib.Client{
		HomeDir: homeDir,
		Config:  clientCfg,
	}

	id, err := client.LoadMyIdentity()
	if err != nil {
		return "", err
	}

	req := &api.RegistrationRequest{
		Name:        "peer1",
		Type:        "peer",
		Affiliation: "org1.department1",
		Secret:      "peer1pw",
		CAName:      clientCfg.CAName,
	}
	resp, err := id.Register(req)
	if err != nil {
		return "", err
	}
	logger.Infof("Register %s success.", req.Name)
	return resp.Secret, nil
}
