package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ExampleHandler interface {
	Handler1() echo.HandlerFunc
}

type exampleHandler struct {
}

func NewHandler() ExampleHandler {
	return &exampleHandler{}
}

func (r *exampleHandler) Handler1() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok lah")
	}
}
