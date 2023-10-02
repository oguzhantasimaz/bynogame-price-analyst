package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	glog "github.com/labstack/gommon/log"
	mw "github.com/oguzhantasimaz/bynogame-price-analyst/core/api/middleware"
	"github.com/oguzhantasimaz/bynogame-price-analyst/core/api/route"
	"github.com/oguzhantasimaz/bynogame-price-analyst/core/bootstrap"
	"time"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	echoInstance := echo.New()

	echoInstance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderAccessControlAllowOrigin},
	}))
	echoInstance.Use(middleware.Logger())
	echoInstance.Use(middleware.RateLimiterWithConfig(mw.CustomRateLimiterConfig))
	echoInstance.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Request timeout",
		Timeout:      30 * time.Second,
	}))
	echoInstance.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  glog.ERROR,
	}))

	route.Setup(db, echoInstance)

	echoInstance.Logger.Fatal(echoInstance.Start(env.ServerAddress))
}
