package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var Configer = viper.GetViper()

func init() {
	Configer.AutomaticEnv()
	Configer.SetConfigName("config")
	Configer.SetConfigType("yaml")

	// current dir or parent dir
	Configer.AddConfigPath(".")
	Configer.AddConfigPath("..")

	cp := os.Getenv("CONFIG_PATH")
	if cp != "" {
		Configer.AddConfigPath(cp)
	}

	Configer.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := Configer.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
