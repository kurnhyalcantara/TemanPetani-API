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
		return err
	}

	if errInitMigration := migration.InitMigration(db); errInitMigration != nil {
		return errInitMigration
	}

	initDb, errInitDB := db.DB()
	if errInitDB != nil {
		log.Fatalf("error init db: %v", errInitDB)
	}

	initDb.SetMaxOpenConns(dbConfig.DB_MAX_OPEN_CONNS)
	initDb.SetMaxIdleConns(dbConfig.DB_MAX_IDLE_CONNS)
	initDb.SetConnMaxLifetime(time.Duration(dbConfig.DB_CONN_MAX_LIFETIME))

	if err := initDb.Ping(); err != nil {
		defer initDb.Close()
		log.Fatalf("can't send ping to database: %v", err)
	}

	logs := logger.SetUpLogger()
	logs.Println("Database connected...")

	return nil
}

func InitDB() *gorm.DB {
	dbConfig, errDb := config.LoadDBConfig()
	if errDb != nil {
		log.Fatalf("error init db config: %v", errDb)
	}
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.DB_USER, dbConfig.DB_PASS, dbConfig.DB_HOST, dbConfig.DB_PORT, dbConfig.DB_NAME)
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("error init db open: %v", err)
	}
	return db
}
