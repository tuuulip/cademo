package caserver

import (
	"path/filepath"

	"github.com/hyperledger/fabric-ca/lib"
)

func Register() (string, error) {
	homeDir := getHomeDir()
	clientCfg := &lib.ClientConfig{}
	client := lib.Client{
		HomeDir: filepath.Dir(homeDir),
		Config:  clientCfg,
	}

	id, err := client.LoadMyIdentity()
	if err != nil {
		return "", err
	}

	clientCfg.ID.CAName = clientCfg.CAName
	resp, err := id.Register(&clientCfg.ID)
	if err != nil {
		return "", err
	}
	logger.Info("Register success")
	return resp.Secret, nil
}
