package controllers

import (
	"net/http"
	"strconv"

	"github.com/juliotorresmoreno/macabro/db"
	"github.com/juliotorresmoreno/macabro/helper"
	"github.com/juliotorresmoreno/macabro/models"
	"github.com/labstack/echo"
)

type businessController struct {
}

func (that businessController) GET(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*models.User)
	if !session.ACL.IsAdmin() && strconv.Itoa(session.ID) != c.Param("user_id") {
		return echo.NewHTTPError(401, "Unauthorized")
	}

	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err))
	}
	defer conn.Close()

	b := &models.Business{}
	_, err = conn.Where("user_id = ?", c.Param("user_id")).Get(b)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err))
	}

	return c.JSON(200, b)
}

func (that businessController) PATCH(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*models.User)
	if !session.ACL.IsAdmin() && strconv.Itoa(session.ID) != c.Param("user_id") {
		return echo.NewHTTPError(401, "Unauthorized")
	}

	actualBusiness := &models.Business{User: &models.User{}}
	updateBusiness := &models.Business{User: &models.User{}}
	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helper.ParseError(err))
	}
	defer conn.Close()
	conn.Where("user_id = ?", session.ID).Get(actualBusiness)
	if err := c.Bind(updateBusiness); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, helper.ParseError(err).Error())
	}
	if actualBusiness.User.ID == 0 {
		actualBusiness.User = session
	}
	if !session.ACL.IsAdmin() && session.ID != actualBusiness.User.ID {
		return echo.NewHTTPError(401, "Unauthorized")
	}

	actualBusiness.Name = updateBusiness.Name
	actualBusiness.Nit = updateBusiness.Nit
	actualBusiness.LegalRepresentative = updateBusiness.LegalRepresentative
	actualBusiness.Website = updateBusiness.Website
	actualBusiness.Address = updateBusiness.Address
	actualBusiness.Country = updateBusiness.Country
	actualBusiness.City = updateBusiness.City
	actualBusiness.EconomicActivity = updateBusiness.EconomicActivity
	actualBusiness.ImgSrc = updateBusiness.ImgSrc

	if err := actualBusiness.Check(); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, helper.ParseError(err).Error())
	}

	if actualBusiness.ID == 0 {
		_, err = conn.InsertOne(actualBusiness)
	} else {
		_, err = conn.Where("id = ?", actualBusiness.ID).Update(actualBusiness)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, helper.ParseError(err).Error())
	}

	return c.String(http.StatusNoContent, "")
}

func BusinessController(g *echo.Group) {
	c := businessController{}
	g.GET("/:user_id", c.GET)
	g.PATCH("/:user_id", c.PATCH)
}
