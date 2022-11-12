package handlers

import (
	"fmt"
	"marketbill-messaging-service/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleSMS(c echo.Context) error {
	err := services.SendDefaultSMS("01091751159", "01091751159", "안녕 난 테스트야")
	fmt.Println(err)
	return c.String(http.StatusOK, "heehhe")
}
