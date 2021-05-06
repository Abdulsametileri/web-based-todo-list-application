package config

import (
	"github.com/spf13/viper"
	"log"
)

func Setup() {
	if IsDebug {
	} else {
		viper.SetConfigFile(".env")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error while reading .env file %s", err)
		}
	}
}
