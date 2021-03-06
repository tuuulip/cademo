package caserver

import (
	"fmt"

	"cademo/config"
	"cademo/log"

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

	// set admin info
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
	// set db config
	db := getServerDBConfig()
	// set affiliations
	affiliations := map[string]interface{}{
		"org1": []string{"department1", "department2"},
		"org2": []string{"department1", "department2"},
	}
	// set csr info
	csr := serverCSR()

	// set server config
	serverCfg := &lib.ServerConfig{
		CAcfg: lib.CAConfig{
			CA: lib.CAInfo{
				Name: config.C.GetString("caname"),
			},
			DB:           db,
			Affiliations: affiliations,
			Registry: lib.CAConfigRegistry{
				MaxEnrollments: -1,
				Identities:     []lib.CAConfigIdentity{id},
			},
			CSR: *csr,
		},
		TLS: getServerTls(),
	}
	serverCfg.CAcfg.Cfg.Identities.AllowRemove = true
	return serverCfg
}

func getServerDBConfig() lib.CAConfigDB {
	dbsrc := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.C.GetString("cadb.host"),
		config.C.GetInt("cadb.port"),
		config.C.GetString("cadb.user"),
		config.C.GetString("cadb.password"),
		config.C.GetString("cadb.dbname"),
		config.C.GetString("cadb.sslmode"),
	)
	db := lib.CAConfigDB{
		Type:       "postgres",
		Datasource: dbsrc,
	}
	return db
}
