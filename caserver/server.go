package caserver

import (
	"fmt"

	"cademo/config"
	"cademo/log"

	"github.com/cloudflare/cfssl/csr"
	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/lib"
)

var logger = log.GetLogger("info")

func GetServer() *lib.Server {
	homeDir := getHomeDir()

	serverCfg := getServerCfg()
	blockingStart := false
	return &lib.Server{
		HomeDir:       homeDir,
		Config:        serverCfg,
		BlockingStart: blockingStart,
		CA: lib.CA{
			Config:         &serverCfg.CAcfg,
			ConfigFilePath: homeDir,
		},
	}
}

func getServerCfg() *lib.ServerConfig {
	dbsrc := fmt.Sprintf(
		"host=%s port=5432 user=%s password=%s dbname=%s sslmode=%s",
		config.C.GetString("cadb.host"),
		config.C.GetString("cadb.user"),
		config.C.GetString("cadb.password"),
		config.C.GetString("cadb.dbname"),
		config.C.GetString("cadb.sslmode"),
	)
	id := lib.CAConfigIdentity{
		Name:           config.C.GetString("caadmin.user"),
		Pass:           config.C.GetString("caadmin.pass"),
		Type:           "client",
		MaxEnrollments: -1,
		Attrs: map[string]string{
			"hf.Registrar.Roles":         "*",
			"hf.Registrar.DelegateRoles": "*",
			"hf.Revoker":                 "true",
			"hf.IntermediateCA":          "true",
			"hf.GenCRL":                  "true",
			"hf.Registrar.Attributes":    "*",
			"hf.AffiliationMgr":          "true",
		},
	}
	db := lib.CAConfigDB{
		Type:       "postgres",
		Datasource: dbsrc,
	}
	affiliations := map[string]interface{}{
		"org1": []string{"department1", "department2"},
		"org2": []string{"department1", "department2"},
	}
	csrName := csr.Name{
		C:  "China",
		ST: "Guang Dong",
		L:  "Guang Zhou",
		O:  "My Company",
		OU: "Development",
	}
	csr := api.CSRInfo{
		CN:    "fabric-ca-server",
		Names: []csr.Name{csrName},
	}
	serverCfg := &lib.ServerConfig{
		CAcfg: lib.CAConfig{
			DB:           db,
			Affiliations: affiliations,
			Registry: lib.CAConfigRegistry{
				MaxEnrollments: -1,
				Identities:     []lib.CAConfigIdentity{id},
			},
			CSR: csr,
		},
	}
	return serverCfg
}
