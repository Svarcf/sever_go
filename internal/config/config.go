package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DB_HOST  string
	DB_PORT  string
	DB_PASS  string
	DB_USER  string
	DB_NAME  string
	APP_PORT string
	APP_NAME string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("APP_NAME", "sever-go")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}
	config := &Config{
		DB_HOST:  viper.GetString("DB_HOST"),
		DB_PORT:  viper.GetString("DB_PORT"),
		DB_PASS:  viper.GetString("DB_PASS"),
		DB_USER:  viper.GetString("DB_USER"),
		DB_NAME:  viper.GetString("DB_NAME"),
		APP_PORT: viper.GetString("APP_PORT"),
		APP_NAME: viper.GetString("APP_NAME"),
	}

	return config, nil
}
