package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type config struct {
	Port    int
	Name    string
	PathMap string `mapstructure:"path_map"`
}

// C is global configuration
var C config

func InitConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	// Find home directory.
	// 	home, err := homedir.Dir()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// 	// Search config in home directory with name ".cobra" (without extension).
	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigName(".cobra")
	// }

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
