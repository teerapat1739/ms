package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/subosito/gotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	partyrepo "party-service/business/party/repository"
	partysrc "party-service/business/party/service"
	partyhdlr "party-service/business/party/transport"
)

func init() {
	gotenv.Load()
}

func main() {
	conString := os.Getenv("MYSQL_CONN")
	dbCon, err := gorm.Open(mysql.Open(conString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Recover())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 6,
	}))

	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/hellworld", func(c echo.Context) error {
		return c.String(http.StatusOK, "hellworld")
	})
	partyrepo := partyrepo.NewRepository(dbCon, "table_name")

	partysrc := partysrc.NewUserService(partyrepo, "")

	partyhdlr.NewPartyHTTPHandler(e, partysrc)

	go func() {
		if err := e.Start(":7777"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
