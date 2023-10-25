package config

func LoadAllConfigs() (*AppConfig, *DBConfig, error) {
	appConfig, errApp := LoadAppConfig()
	if errApp != nil {
		return nil, nil, errApp
	}

	dbConfig, errDB := LoadDBConfig()
	if errDB != nil {
		return nil, nil, errDB
	}

	return appConfig, dbConfig, nil
}