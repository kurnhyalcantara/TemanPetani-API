package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type DBConfig struct {
	DB_HOST              string
	DB_PORT              int
	DB_USER              string
	DB_PASS              string
	DB_NAME              string
	DB_MAX_OPEN_CONNS    int
	DB_MAX_IDLE_CONNS    int
	DB_CONN_MAX_LIFETIME int
}

func LoadDBConfig() (*DBConfig, error) {
	dbHost, foundDBHost := os.LookupEnv("DB_HOST")
	dbPort, foundDBPort := os.LookupEnv("DB_PORT")
	dbUser, foundDBUser := os.LookupEnv("DB_USER")
	dbPass, foundDBPass := os.LookupEnv("DB_PASS")
	dbName, foundDBName := os.LookupEnv("DB_NAME")
	dbMaxOpenConns, foundDBMaxOpenConns := os.LookupEnv("DB_MAX_OPEN_CONNS")
	dbMaxIdleConns, foundDBMaxIdleConns := os.LookupEnv("DB_MAX_IDLE_CONNS")
	dbConnMaxLifetime, foundDBConnMaxLifetime := os.LookupEnv("DB_CONN_MAX_LIFETIME")

	if !foundDBHost {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env.local")
		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	if !foundDBHost {
		dbHost = viper.GetString("DB_HOST")
	}

	if !foundDBPort {
		dbPort = viper.GetString("DB_PORT")
	}

	if !foundDBUser {
		dbUser = viper.GetString("DB_USER")
	}

	if !foundDBPass {
		dbPass = viper.GetString("DB_PASS")
	}

	if !foundDBName {
		dbName = viper.GetString("DB_NAME")
	}

	if !foundDBMaxOpenConns {
		dbMaxOpenConns = viper.GetString("DB_MAX_OPEN_CONNS")
	}

	if !foundDBMaxIdleConns {
		dbMaxIdleConns = viper.GetString("DB_MAX_IDLE_CONNS")
	}

	if !foundDBConnMaxLifetime {
		dbConnMaxLifetime = viper.GetString("DB_CONN_MAX_LIFETIME")
	}

	dbPortInt, _ := strconv.Atoi(dbPort)
	dbMaxOpenConnsInt, _ := strconv.Atoi(dbMaxOpenConns)
	dbMaxIdleConnsInt, _ := strconv.Atoi(dbMaxIdleConns)
	dbConnMaxLifetimeInt, _ := strconv.Atoi(dbConnMaxLifetime)

	return &DBConfig{
		DB_HOST:              dbHost,
		DB_PORT:              dbPortInt,
		DB_USER:              dbUser,
		DB_PASS:              dbPass,
		DB_NAME:              dbName,
		DB_MAX_OPEN_CONNS:    dbMaxOpenConnsInt,
		DB_MAX_IDLE_CONNS:    dbMaxIdleConnsInt,
		DB_CONN_MAX_LIFETIME: dbConnMaxLifetimeInt,
	}, nil
}
