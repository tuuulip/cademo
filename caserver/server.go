package caserver

import (
	"github.com/hyperledger/fabric-ca/lib"
)

func GetServer() *lib.Server {
	homeDir := "/Users/stephen/develop/gotut/cademo/cahome"
	serverCfg := &lib.ServerConfig{}
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
