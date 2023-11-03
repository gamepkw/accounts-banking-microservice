package handler

import (
	"net/http"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
	"github.com/labstack/echo/v4"
)

func (a *AccountHandler) ElasticSearchAccountByAccountNo(c echo.Context) error {

	var account model.ElasticSearchAccount

	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()

	response, err := a.accountService.ElasticSearchAccountByAccountNo(ctx, account)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}
