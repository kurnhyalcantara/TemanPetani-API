package mysqldb

import (
	"fmt"
	"log"
	"time"

	"github.com/kurnhyalcantara/TemanPetani-API/app/config"
	"github.com/kurnhyalcantara/TemanPetani-API/app/database/migration"
	"github.com/kurnhyalcantara/TemanPetani-API/app/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(dbConfig *config.DBConfig) error {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.DB_USER, dbConfig.DB_PASS, dbConfig.DB_HOST, dbConfig.DB_PORT, dbConfig.DB_NAME)
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connect to db: %v", err)
	}

	if errInitMigration := migration.InitMigration(db); errInitMigration != nil {
		log.Fatalf("error migrate db: %v", err)
	}
	
	initDb, err := db.DB()
	if err != nil {
		return err
	}
	
	initDb.SetMaxOpenConns(dbConfig.DB_MAX_OPEN_CONNS)
	initDb.SetMaxIdleConns(dbConfig.DB_MAX_IDLE_CONNS)
	initDb.SetConnMaxLifetime(time.Duration(dbConfig.DB_CONN_MAX_LIFETIME))

	if err := initDb.Ping(); err != nil {
		defer initDb.Close()
		return fmt.Errorf("can't send ping to database: %v", err)
	}

	logs := logger.SetUpLogger()
	logs.Println("Database connected...")

	return nil
}