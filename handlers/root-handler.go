package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func PingPong(c echo.Context) error {
	msg := "Marketbill Messaging Service is running..."
	return c.String(http.StatusOK, msg)
}
