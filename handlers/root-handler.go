package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckServer(c echo.Context) error {
	msg := "Marketbill Messaging Service is running..."
	return c.String(http.StatusOK, msg)
}

func HealthCheck(c echo.Context) error {
	health := map[string]interface{}{
		"HEALTH": "UP",
	}
	return c.JSON(http.StatusOK, health)
}
