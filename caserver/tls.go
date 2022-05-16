package caserver

import (
	"cademo/config"
	"path/filepath"

	"github.com/hyperledger/fabric-ca/lib/tls"
)

func getServerTls() tls.ServerTLSConfig {
	tls := tls.ServerTLSConfig{
		Enabled:  config.C.GetBool("caserver.tls.enabled"),
		CertFile: "tls-cert.pem",
	}
	return tls
}

func getClientTls() tls.ClientTLSConfig {
	file := filepath.Join(getHomeDir(), "tls-cert.pem")
	tls := tls.ClientTLSConfig{
		Enabled:   config.C.GetBool("caserver.tls.enabled"),
		CertFiles: []string{file},
	}
	return tls
}
