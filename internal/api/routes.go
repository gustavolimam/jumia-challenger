package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r Router) RegisterRoutes() {

	r.base.GET("/phones", r.ReadPhoneNumbers)
}

func (r Router) ReadPhoneNumbers(c echo.Context) error {

	res, err := r.service.ReadPhoneNumbers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
