package caserver

import (
	"fmt"
	"os"
	"path/filepath"

	"cademo/config"

	"github.com/hyperledger/fabric-ca/lib"
)

func GetServer() *lib.Server {
	homeDir := getCaHome()
	dbsrc := fmt.Sprintf(
		"host=%s port=5432 user=%s password=%s dbname=%s sslmode=%s",
		config.Configer.GetString("cadb.host"),
		config.Configer.GetString("cadb.user"),
		config.Configer.GetString("cadb.password"),
		config.Configer.GetString("cadb.dbname"),
		config.Configer.GetString("cadb.sslmode"),
	)
	id := lib.CAConfigIdentity{
		Name:           config.Configer.GetString("caadmin.user"),
		Pass:           config.Configer.GetString("caadmin.pass"),
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

func getCaHome() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(pwd, config.Configer.GetString("cahome"))
}
