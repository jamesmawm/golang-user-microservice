package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return
}

func GetPort() string {
	return viper.GetString("APP_PORT")
}

func GetSecret() string {
	return viper.GetString("APP_SECRET")
}
