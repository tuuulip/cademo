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
