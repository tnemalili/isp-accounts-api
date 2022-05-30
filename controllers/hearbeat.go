package controllers

import (
	"github.com/labstack/echo/v4"
	"isp.accounts.api/core/models"
	"net/http"
)

func HealthCheck(ctx echo.Context) error {

	message := &models.GenericMessage{Message: "Heart rate Ok", Success: true}
	return ctx.JSON(http.StatusOK, message)
}