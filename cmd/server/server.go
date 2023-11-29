package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kurnhyalcantara/TemanPetani-API/app/config"
	mysqldb "github.com/kurnhyalcantara/TemanPetani-API/app/database/mysql"
	"github.com/kurnhyalcantara/TemanPetani-API/app/logger"
	"github.com/labstack/echo/v4"
)

func Serve() {
	appConfig, dbConfig, errLoadConfig := config.LoadAllConfigs()
	if errLoadConfig != nil {
		log.Fatalf("error load configs: %v", errLoadConfig)
	}
	loggr := logger.SetUpLogger()

	// Connect to database
	errDB := mysqldb.ConnectDB(dbConfig)
	if errDB != nil {
		log.Fatalf("error connect to db: %v", errDB)
	}

	// Create http server
	e := echo.New()

	// go routine signal channel stop server
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	go func() {
		<-sigCh
		loggr.Warnf("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = e.Shutdown(ctx)
	}()

	// Start server
	serverAddr := fmt.Sprintf("%s:%d", appConfig.HOST, appConfig.PORT)
	e.Start(serverAddr)
}
