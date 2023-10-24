package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type AppConfig struct {
	HOST	string
	PORT	int
}

func LoadAppConfig() (*AppConfig, error) {
	host, foundAppHost := os.LookupEnv("APP_HOST")
	port, foundAppPort := os.LookupEnv("APP_PORT")

	if !foundAppHost || !foundAppPort {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env.local")
		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}

		if !foundAppHost {
			host = viper.GetString("APP_HOST")
		}

		if !foundAppPort {
			port = viper.GetString("APP_PORT")
		}
	}

	portInt, _ := strconv.Atoi(port)

	return &AppConfig{
		HOST: host,
		PORT: portInt,
	}, nil
}