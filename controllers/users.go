package controllers

import (
	"net/http"

	"github.com/juliotorresmoreno/macabro/db"
	"github.com/juliotorresmoreno/macabro/helper"
	"github.com/juliotorresmoreno/macabro/models"
	"github.com/labstack/echo"
)

type usersController struct {
}

func (that usersController) GET(c echo.Context) error {
	u := make([]models.User, 0)
	conn, err := db.NewEngigne()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.Find(&u)
	if err != nil {
		return &echo.HTTPError{
			Code:     501,
			Message:  err.Error(),
			Internal: err,
		}
	}
	return c.JSON(200, u)
}

func (that usersController) PUT(c echo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusNotAcceptable,
			Message: err.Error(),
		}
	}
	conn, err := db.NewEngigne()
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: helper.ParseError(err).Error(),
		}
	}
	defer conn.Close()
	if err := u.Check(); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	if _, err := conn.InsertOne(u); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: helper.ParseError(err).Error(),
		}
	}
	return c.JSON(202, u)
}

// UsersController s
func UsersController(g *echo.Group) {
	u := usersController{}
	g.GET("", u.GET)
	g.PUT("", u.PUT)
}
