package handler

import (
	"net/http"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
	"github.com/labstack/echo/v4"
)

func (a *AccountHandler) UpdateAccount(c echo.Context) (err error) {
	var account model.Account
	err = c.Bind(&account)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&account); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.AService.UpdateAccount(ctx, &account)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, AccountResponse{Message: "Update account successfully", Body: &account})
}
