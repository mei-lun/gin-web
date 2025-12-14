package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}

	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = &Config{}

	err := viper.Unmarshal(AppConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
}
