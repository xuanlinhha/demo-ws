package delivery

import (
	"demo-ws/example-1/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler1 interface {
	Handler1() echo.HandlerFunc
}

type handler1 struct {
	Service1 service.Service1
}

func NewHandler1(service1 service.Service1) Handler1 {
	return &handler1{Service1: service1}
}

func (r *handler1) Handler1() echo.HandlerFunc {
	return func(c echo.Context) error {
		count, err := r.Service1.GetRowsNo()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "GetRowsNo error")
		} else {
			return c.JSON(http.StatusOK, count)
		}
	}
}
