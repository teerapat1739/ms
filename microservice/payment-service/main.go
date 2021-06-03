package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/subosito/gotenv"

	paymentrepo "payment-service/business/payment/repository"
	paymentsrc "payment-service/business/payment/service"
	paymenthdlr "payment-service/business/payment/transport"
)

func init() {
	gotenv.Load()
}

func main() {
	resdisConfig := os.Getenv("REDIS_SERVER")
	fmt.Println("Go Redis Tutorial", resdisConfig)

	client := redis.NewClient(&redis.Options{
		Addr:     resdisConfig,
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	e := echo.New()
	e.Use(middleware.Recover())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 6,
	}))
	// api-payment
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/hellworld", func(c echo.Context) error {
		return c.String(http.StatusOK, "hellworld")
	})
	paymentrepo := paymentrepo.NewRepository(client)

	paymentsrc := paymentsrc.NewUserService(paymentrepo, "")

	paymenthdlr.NewpaymentHTTPHandler(e, paymentsrc)

	go func() {
		if err := e.Start(":9999"); err != nil {
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
