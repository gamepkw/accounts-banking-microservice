package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *AccountHandler) GetCountAccount(c echo.Context) error {
	ctx := c.Request().Context()

	list, err := a.AService.GetCountAccount(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, list)
}
