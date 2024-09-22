package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST  string
	DB_PORT  string
	DB_PASS  string
	DB_USER  string
	DB_NAME  string
	DB_INIT  bool
	APP_PORT string
	APP_NAME string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	viper.SetDefault("app.port", "8080")
	viper.SetDefault("app.name", "sever-go")
	viper.SetDefault("db.init", false)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Error reading config file, %s", err))
	}

	config := &Config{
		DB_HOST:  viper.GetString("db.host"),
		DB_PORT:  viper.GetString("db.port"),
		DB_PASS:  viper.GetString("db.pass"),
		DB_USER:  viper.GetString("db.user"),
		DB_NAME:  viper.GetString("db.name"),
		DB_INIT:  viper.GetBool("db.init"),
		APP_PORT: viper.GetString("app.port"),
		APP_NAME: viper.GetString("app.name"),
	}

	return config
}
