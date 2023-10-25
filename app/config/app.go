package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type AppConfig struct {
	HOST	string
	PORT	int
	DEBUG bool
	JWT_SECRET_KEY string
	JWT_REFRESH_KEY	string
}

var appConfig = &AppConfig{}

func LoadAppConfig() (*AppConfig, error) {
	host, foundAppHost := os.LookupEnv("APP_HOST")
	port, foundAppPort := os.LookupEnv("APP_PORT")
	debug, foundDebug := os.LookupEnv("APP_DEBUG")
	jwtSecretKey, foundJwtSecretKey := os.LookupEnv("JWT_SECRET_KEY")
	jwtRefreshKey, foundJwtRefreshKey := os.LookupEnv("JWT_REFRESH_KEY")

	if !foundAppHost || !foundAppPort || !foundDebug || !foundJwtSecretKey || !foundJwtRefreshKey {
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

		if !foundDebug {
			debug = viper.GetString("APP_DEBUG")
		}

		if !foundJwtSecretKey {
			jwtSecretKey = viper.GetString("JWT_SECRET_KEY")
		}

		if !foundJwtRefreshKey {
			jwtRefreshKey = viper.GetString("JWT_REFRESH_KEY")
		}
	}

	portInt, _ := strconv.Atoi(port)
	debugBool, _ := strconv.ParseBool(debug)

	return &AppConfig{
		HOST: host,
		PORT: portInt,
		DEBUG: debugBool,
		JWT_SECRET_KEY: jwtSecretKey,
		JWT_REFRESH_KEY: jwtRefreshKey,
	}, nil
}

func GetAppConfig() *AppConfig {
	return appConfig
}