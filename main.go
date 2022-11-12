package main

import (
	"marketbill-messaging-service/config"
	"marketbill-messaging-service/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ReadConfig()

	e := echo.New()
	e.GET("/", handlers.PingPong)
	e.POST("/messaging/sms", handlers.HandleSMS)
	e.Logger.Fatal(e.Start(":" + config.C.Server.Port))
}
