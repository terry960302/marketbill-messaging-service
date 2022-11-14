package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func CheckServer(c echo.Context) error {
	profile := os.Getenv("PROFILE")
	msg := fmt.Sprintf("[%s]", profile) + "Marketbill Messaging Service is running..."
	return c.String(http.StatusOK, msg)
}

func HealthCheck(c echo.Context) error {
	health := map[string]interface{}{
		"status": "UP",
	}
	return c.JSON(http.StatusOK, health)
}
