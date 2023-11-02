package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *AccountHandler) GetAccountByAccountNo(c echo.Context) error {
	account_no := c.Param("account_no")
	ctx := c.Request().Context()

	account, err := a.AService.GetAccountByAccountNo(ctx, account_no)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, account)
}
