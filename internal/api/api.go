package api

import (
	"github.com/jumia-challenger/internal/service"
	"github.com/labstack/echo/v4"
)

type Routers interface {
	RegisterRoutes()
}

type Router struct {
	base    *echo.Echo
	service service.Service
}

func Start(e *echo.Echo) Routers {

	return Router{e, service.New()}
}
