package caserver

import (
	"cademo/config"

	"github.com/cloudflare/cfssl/csr"
	"github.com/hyperledger/fabric-ca/api"
)

func serverCSR() *api.CSRInfo {
	csrName := csr.Name{
		C:  "CN",
		ST: "Guang Dong",
		L:  "Guang Zhou",
		O:  "Ca",
	}
	csr := api.CSRInfo{
		CN:    "fabric-ca-server",
		Names: []csr.Name{csrName},
		Hosts: []string{config.C.GetString("caserver.host")},
		KeyRequest: &api.KeyRequest{
			ReuseKey: true,
		},
	}
	return &csr
}

func clientCSR(organization, domain string) *api.CSRInfo {
	csrName := csr.Name{
		C:  "CN",
		ST: "Guang Dong",
		L:  "Guang Zhou",
		O:  organization,
	}
	hosts := []string{
		domain,
		config.C.GetString("caserver.host"),
	}
	csr := api.CSRInfo{
		CN:    domain,
		Names: []csr.Name{csrName},
		Hosts: hosts,
		KeyRequest: &api.KeyRequest{
			ReuseKey: true,
		},
	}
	return &csr
}
