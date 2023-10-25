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
	"github.com/labstack/echo/v4"
)

func Serve() {
	appConfig, dbConfig, err := config.LoadAllConfigs();
	if err != nil {
		log.Fatalf("error load configs: %v", err)
	}

	// Connect to database
	if errDB := mysqldb.ConnectDB(dbConfig); errDB != nil {
		log.Fatalf("error connect to db: %v", err)
	}

	// Create http server
	e := echo.New()
	serverAddr := fmt.Sprintf("%s:%d", appConfig.HOST, appConfig.PORT)
	go func() {
		if err := e.Start(serverAddr); err != nil {
			log.Fatalf("Opps... init server error: %v", err)
		}
	}()

	// go routine signal channel stop server
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sigCh
	fmt.Printf("\nShutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}	
}