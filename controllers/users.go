package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/juliotorresmoreno/macabro/db"
	"github.com/juliotorresmoreno/macabro/helper"
	"github.com/juliotorresmoreno/macabro/models"
	"github.com/labstack/echo"
)

type usersController struct {
}

func (that usersController) GET(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*models.User)
	json.NewEncoder(os.Stdout).Encode(session.ACL)
	log.Println(session.ID)
	u := make([]models.User, 0)
	conn, err := db.NewEngigne()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.Find(&u, models.User{ID: session.ID})
	if err != nil {
		return echo.NewHTTPError(501, err.Error())
	}
	return c.JSON(200, u)
}

func (that usersController) PUT(c echo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, err.Error())
	}
	conn, err := db.NewEngigne()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
	}
	defer conn.Close()
	if err := u.Check(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if _, err := conn.InsertOne(u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err).Error())
	}
	return c.JSON(202, u)
}

func (that usersController) PATCH(c echo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, err.Error())
	}
	conn, err := db.NewEngigne()
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			helper.ParseError(err).Error(),
		)
	}
	defer conn.Close()
	if err := u.Check(); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	q := &models.User{ID: 1}
	if _, err := conn.Update(u, q); err != nil {
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
	g.PATCH("", u.PATCH)
}
