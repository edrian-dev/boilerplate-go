package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_api "github.com/nomada-sh/levita-stp/api"
	_stpDelivery "github.com/nomada-sh/levita-stp/api/stp/delivery/restful"
	_userDelivery "github.com/nomada-sh/levita-stp/api/user/delivery/restful"
)

const version = "0.1"

func NewRouter(port string) {
	server := echo.New()
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderAuthorization, echo.HeaderContentType, echo.HeaderXRequestedWith},
		AllowMethods: []string{echo.POST, echo.GET, echo.PUT, echo.DELETE, echo.HEAD},
	}))

	server.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, version)
	})

	_userDelivery.NewRouter(_userDelivery.Input{
		Router: server,
		User:   _api.User,
	})

	_stpDelivery.NewRouter(_stpDelivery.Input{
		Router: server,
		STP:    _api.STP,
	})

	server.Start(port)
}
