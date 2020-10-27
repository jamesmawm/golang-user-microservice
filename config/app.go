package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var App Configurations

type Configurations struct {
	Debug    bool
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port                            string
	WaitDurationForGracefulShutdown string
}

type DatabaseConfig struct {
	Connection string
	Host       string
	Name       string
	Username   string
	Password   string
}

func ReadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.BindEnv("database.connection", "DB_CONNECTION")
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.name", "DB_DATABASE")
	viper.BindEnv("database.username", "DB_USERNAME")
	viper.BindEnv("database.password", "DB_PASSWORD")

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&App)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Unable to decode into struct: %v \n", err))
	}

	return
}
