package controllers

import (
	"net/http"
	"strconv"

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
	if !session.ACL.IsAdmin() && strconv.Itoa(session.ID) != c.Param("user_id") {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	u := &models.User{}
	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Where("id = ?", c.Param("user_id")).Get(u)
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
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*models.User)
	if !session.ACL.IsAdmin() && strconv.Itoa(session.ID) != c.Param("user_id") {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	actualUser := &models.User{}
	updateUser := &models.User{}
	if err := c.Bind(updateUser); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, err.Error())
	}
	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			helper.ParseError(err).Error(),
		)
	}
	defer conn.Close()
	conn.Get(actualUser)
	if actualUser.ID == 0 {
		return echo.NewHTTPError(401, "unautorized")
	}
	actualUser.Password = ""
	actualUser.ValidPassword = ""
	actualUser.Name = newValueString(updateUser.Name, actualUser.Name)
	actualUser.LastName = newValueString(updateUser.LastName, actualUser.LastName)

	actualUser.DocumentType = newValueString(updateUser.DocumentType, actualUser.DocumentType)
	actualUser.Expedite = newValueTime(updateUser.Expedite, actualUser.Expedite)
	actualUser.Document = newValueString(updateUser.Document, actualUser.Document)
	actualUser.DateBirth = newValueTime(updateUser.DateBirth, actualUser.DateBirth)
	actualUser.ImgSrc = newValueString(updateUser.ImgSrc, actualUser.ImgSrc)
	actualUser.Country = newValueString(updateUser.Country, actualUser.Country)
	actualUser.Nationality = newValueString(updateUser.Nationality, actualUser.Nationality)
	actualUser.Facebook = newValueString(updateUser.Facebook, actualUser.Facebook)
	actualUser.Linkedin = newValueString(updateUser.Linkedin, actualUser.Linkedin)

	if err := actualUser.Check(); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			helper.ParseError(err).Error(),
		)
	}
	if _, err := conn.Where("id = ?", actualUser.ID).Update(actualUser); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			helper.ParseError(err).Error(),
		)
	}
	return c.JSON(http.StatusAccepted, actualUser)
}

// UsersController s
func UsersController(g *echo.Group) {
	u := usersController{}
	g.GET("/:user_id", u.GET)
	g.PUT("", u.PUT)
	g.PATCH("/:user_id", u.PATCH)
}
