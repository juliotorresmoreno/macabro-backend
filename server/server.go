package server

import (
	"fmt"

	"github.com/juliotorresmoreno/macabro/config"
	"github.com/juliotorresmoreno/macabro/controllers"
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
	}))

	api := e.Group("/api")
	controllers.UsersController(api.Group("/users"))
	controllers.AuthController(api.Group("/auth"))
	return &FastServerHTTP{e}
}