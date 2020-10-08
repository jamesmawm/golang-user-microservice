package config

import "github.com/spf13/viper"

func GetDatabase() string {
	return viper.GetString("DB_CONNECTION")
}

func GetDatabaseHost() string {
	return viper.GetString("DB_HOST")
}

func GetDatabasePort() string {
	return viper.GetString("DB_PORT")
}

func GetDatabaseName() string {
	return viper.GetString("DB_DATABASE")
}

func GetDatabaseUser() string {
	return viper.GetString("DB_USERNAME")
}

func GetDatabasePassword() string {
	return viper.GetString("DB_PASSWORD")
}
