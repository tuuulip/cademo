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
	return filepath.Join(pwd, config.Configer.GetString("cahome"))
}

func getAdminDir() string {
	return filepath.Join(getHomeDir(), "admin")
}
