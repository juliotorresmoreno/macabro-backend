package server

import (
	"fmt"

	"github.com/juliotorresmoreno/macabro/config"
	"github.com/juliotorresmoreno/macabro/controllers"
	"github.com/juliotorresmoreno/macabro/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// FastServerHTTP s
type FastServerHTTP struct {
	*echo.Echo
}

// Listen s
func (that FastServerHTTP) Listen() {
	conf := config.GetConfig()
	host := fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	that.Echo.Start(host)
}

// NewFastServerHTTP s
func NewFastServerHTTP() *FastServerHTTP {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
		},
		AllowMethods: []string{
			echo.GET, echo.POST, echo.PATCH,
			echo.PUT, echo.DELETE, echo.HEAD,
			echo.OPTIONS,
		},
	}))

	api := e.Group("/api")
	api.Use(middlewares.Session)
	controllers.UsersController(api.Group("/users"))
	controllers.AuthController(api.Group("/auth"))
	controllers.InstancesController(api.Group("/instances"))
	controllers.BusinessController(api.Group("/business"))
	controllers.PaymentMethodsController(api.Group("/payment-methods"))

	return &FastServerHTTP{e}
}
