package caclient

import (
	"fmt"
	"path/filepath"

	"github.com/hyperledger/fabric-ca/lib"
)

func Register() error {
	homeDir := "/Users/stephen/develop/gotut/cademo/cahome"
	clientCfg := &lib.ClientConfig{}
	client := lib.Client{
		HomeDir: filepath.Dir(homeDir),
		Config:  clientCfg,
	}

	id, err := client.LoadMyIdentity()
	if err != nil {
		return err
	}

	clientCfg.ID.CAName = clientCfg.CAName
	resp, err := id.Register(&clientCfg.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Password: %s\n", resp.Secret)

	return nil
}
