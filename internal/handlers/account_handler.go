package handler

import (
	"net/http"

	"github.com/gamepkw/accounts-banking-microservice/internal/middleware"
	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
	accountService "github.com/gamepkw/accounts-banking-microservice/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

type AccountHandler struct {
	AService accountService.AccountService
}

func NewAccountHandler(e *echo.Echo, us accountService.AccountService) {
	handler := &AccountHandler{
		AService: us,
	}
	restrictedGroup := e.Group("/users/accounts")
	restrictedGroup.Use(middleware.CustomJWTMiddleware)

	e.GET("/accounts", handler.GetAllAccount)
	// e.POST("/accounts/register", handler.RegisterAccount)
	e.GET("/accounts/:account_no", handler.GetAccountByAccountNo)
	e.GET("/accounts-limit/:account_no", handler.GetDailyLimit)
	e.GET("/accounts-daily-limit/:account_no", handler.GetSumDailyTransaction)
	e.PUT("/accounts/:account_no", handler.UpdateAccount)
	e.PUT("/accounts/:account_no", handler.CloseAccount)
	e.GET("/accounts/get-count-by-status", handler.GetCountAccount)

	restrictedGroup.GET("/get-all-account", handler.GetAllAccountByUuid)
	restrictedGroup.POST("/register", handler.RegisterAccount)
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AccountResponse struct {
	Message string         `json:"message"`
	Body    *model.Account `json:"body,omitempty"`
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case model.ErrInternalServerError:
		return http.StatusInternalServerError
	case model.ErrNotFound:
		return http.StatusNotFound
	case model.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func isRequestValid(m *model.Account) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
