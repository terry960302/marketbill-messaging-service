package handlers

import (
	"marketbill-messaging-service/models"
	"marketbill-messaging-service/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleSMS(c echo.Context) error {
	req := models.DefaultSmsRequest{}
	if err := c.Bind(&req); err != nil {
		c.Error(err)
	}

	res, err := services.SendDefaultSMS(req.To, req.Message)
	if err != nil {
		c.Error(err)
	}

	return c.JSON(http.StatusOK, res)
}
