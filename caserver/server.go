package caserver

import (
	"fmt"

	"github.com/hyperledger/fabric-ca/lib"
)

func GetServer() *lib.Server {
	homeDir := "/Users/stephen/develop/gotut/cademo/cahome"
	dbsrc := fmt.Sprintf(
		"host=%s port=5432 user=%s password=%s dbname=%s sslmode=%s",
		"10.10.7.35",
		"root",
		"baas123",
		"fabric_ca_zoo",
		"disable",
	)
	id := lib.CAConfigIdentity{
		Name:           "admin",
		Pass:           "adminpw",
		MaxEnrollments: -1,
	}
	db := lib.CAConfigDB{
		Type:       "postgres",
		Datasource: dbsrc,
	}
	serverCfg := &lib.ServerConfig{
		CAcfg: lib.CAConfig{
			DB: db,
			Registry: lib.CAConfigRegistry{
				MaxEnrollments: -1,
				Identities:     []lib.CAConfigIdentity{id},
			},
		},
	}
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
