package handler

import (
	"net/http"
	"strconv"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
	"github.com/labstack/echo/v4"
)

func (a *AccountHandler) CloseAccount(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("sender"))
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	sender := string(idP)
	ctx := c.Request().Context()

	err = a.accountService.DeleteAccount(ctx, sender)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
