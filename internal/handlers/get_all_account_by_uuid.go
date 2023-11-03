package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *AccountHandler) GetAllAccountByUuid(c echo.Context) error {
	uuid := c.Get("tel").(string)
	ctx := c.Request().Context()

	accounts, err := a.accountService.GetAllAccountByUuid(ctx, uuid)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, accounts)
}
