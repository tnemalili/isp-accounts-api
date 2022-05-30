package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"isp.accounts.api/core"
	"isp.accounts.api/core/models"
	"net/http"
)

func CREATEAccount(ctx echo.Context) error {

	var request models.CreateAccountModel
	err := ctx.Bind(&request)
	if err != nil {
		message := &models.GenericMessage{Message: "Bad Request"}
		return ctx.JSON(http.StatusBadRequest, message)
	}
	success, err := core.DBClient.CreateAccount(&request)
	if err != nil {
		return core.HTTPErrorHandler(ctx, err, http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, &models.GenericMessage{Message: "Created", Success: success})
}

func FETCHAccounts(ctx echo.Context) error {

	accounts, err := core.DBClient.FetchAccounts()
	if err != nil {
		message := &models.GenericMessage{Message: err.Error()}
		return ctx.JSON(http.StatusNotFound, message)
	}

	return ctx.JSON(http.StatusOK, accounts)
}

func FETCHAccount(ctx echo.Context) error {

	id := ctx.Param("id")

	account, err := core.DBClient.FetchAccount(id)
	if err != nil {
		message := &models.GenericMessage{Message: err.Error()}
		return ctx.JSON(http.StatusNotFound, message)
	}
	// NOW ACCOUNT IS FOUND, LET's FETCH THE USER
	fetchAccountHolder, err := core.HTTPClient.FetchCustomerHandler(account.CustomerID)
	if err != nil {
		message := &models.GenericMessage{Message: err.Error()}
		return ctx.JSON(http.StatusNotFound, message)
	}
	account.AccountHolder = core.HTTPResponseHandler(fetchAccountHolder)
	return ctx.JSON(http.StatusOK, account)
}

func UPDATEAccount(ctx echo.Context) error {

	id := ctx.Param("id")
	name := ctx.QueryParam("status")
	if name == "" {
		message := &models.GenericMessage{Message: "Missing Status", Success: false}
		return ctx.JSON(http.StatusBadRequest, message)
	}
	//TODO: HandleError Here
	updated := core.DBClient.UpdateAccountStatus(id, name)
	if !updated {
		message := &models.GenericMessage{Message: "Failed to update an account", Success: updated}
		return ctx.JSON(http.StatusNotFound, message)
	}
	msg := &models.GenericMessage{Message: "Account Updated", Success: updated}
	return ctx.JSON(http.StatusOK, msg)
}

func FUNDAccount(ctx echo.Context)  error {

	id := ctx.Param("id")
	var request models.Amount
	err := ctx.Bind(&request)
	if err != nil { return core.HTTPErrorHandler(ctx, err, http.StatusBadRequest) }
	success := core.DBClient.FundAccount(id, request.Value)
	return ctx.JSON(http.StatusOK, gin.H{"success": success})
}
