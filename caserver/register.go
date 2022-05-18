package caserver

import (
	"cademo/config"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/lib"
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

// get identity
func loadIdentity(home string) (*lib.Identity, error) {
	caurl := getRegisterUrl()

	clientCfg := &lib.ClientConfig{
		URL: caurl,
		TLS: getClientTls(),
	}
	client := lib.Client{
		HomeDir: home,
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
