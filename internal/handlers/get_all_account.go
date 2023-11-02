package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *AccountHandler) GetAllAccount(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	cursor := c.QueryParam("cursor")

	ctx := c.Request().Context()
	// listAr, nextCursor, err := a.AService.Fetch(ctx, cursor, int64(num))
	listAr, nextCursor, err := a.AService.GetAllAccount(ctx, cursor, int64(num))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}
