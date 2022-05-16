package caserver

import (
	"cademo/config"
	"os"
	"path/filepath"
)

func getHomeDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(pwd, config.C.GetString("cahome"))
}
