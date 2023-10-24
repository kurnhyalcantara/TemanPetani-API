package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type DBConfig struct {
	DB_HOST string
	DB_PORT	int
	DB_USER	string
	DB_PASS	string
	DB_NAME string
}

func LoadDBConfig() (*DBConfig, error) {
	dbHost, foundDBHost := os.LookupEnv("DB_HOST")
	dbPort, foundDBPort := os.LookupEnv("DB_PORT")
	dbUser, foundDBUser := os.LookupEnv("DB_USER")
	dbPass, foundDBPass := os.LookupEnv("DB_PASS")
	dbName, foundDBName := os.LookupEnv("DB_NAME")

	viper.AddConfigPath(".")
	viper.SetConfigName(".env.local")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if !foundDBHost {
		dbHost = viper.GetString("DB_HOST")
	}

	if !foundDBPort {
		dbPort = viper.GetString("DB_PORT")
	}
	dbPortInt, _ := strconv.Atoi(dbPort)

	if !foundDBUser {
		dbUser = viper.GetString("DB_USER")
	}

	if !foundDBPass {
		dbPass = viper.GetString("DB_PASS")
	}

	if !foundDBName {
		dbName = viper.GetString("DB_NAME")
	}

	return &DBConfig{
		DB_HOST: dbHost,
		DB_PORT: dbPortInt,
		DB_USER: dbUser,
		DB_PASS: dbPass,
		DB_NAME: dbName,
	}, nil
}