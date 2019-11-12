package main

import (
	"fmt"
	"os"
	"strings"

	"backend-kendo-tutorial/databases"
	"backend-kendo-tutorial/router"

	"github.com/spf13/viper"
)

func main() {

	LoadConfig()

	databases.Open()
	defer databases.Close()

	// 啟動Gin
	app := router.InitRoute()
	app.Run(viper.GetString("server.port"))
}

func LoadConfig() {
	if _, err := os.Stat("./config.yml"); os.IsNotExist(err) {
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	// viper.WatchConfig()
}
