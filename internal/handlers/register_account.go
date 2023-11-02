package handler

import (
	"net/http"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
	"github.com/labstack/echo/v4"
)

func (a *AccountHandler) RegisterAccount(c echo.Context) (err error) {

	uuid := c.Get("tel").(string)

	var account model.Account

	if err = c.Bind(&account); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	account.Uuid = uuid

	var ok bool
	if ok, err = isRequestValid(&account); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()

	if err = a.AService.RegisterAccount(ctx, &account); err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	res, err := a.AService.GetAccountByAccountNo(ctx, account.AccountNo)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}
