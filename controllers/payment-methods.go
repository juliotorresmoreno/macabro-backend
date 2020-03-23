package controllers

import (
	"strconv"

	"github.com/juliotorresmoreno/macabro/db"
	"github.com/juliotorresmoreno/macabro/helper"
	"github.com/juliotorresmoreno/macabro/models"
	"github.com/labstack/echo"
)

type paymentMethodsController struct {
}

func (that paymentMethodsController) GET(c echo.Context) error {
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
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	defer conn.Close()

	p := &models.PaymentMethods{}
	_, err = conn.Where("user_id = ?", session.ID).Get(p)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	p.Decrypt()

	return c.JSON(200, p)
}

func (that paymentMethodsController) PUT(c echo.Context) error {
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
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	defer conn.Close()

	p := &models.PaymentMethods{}
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}

	_, err = conn.InsertOne(p)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	p.Encrypt()

	return c.JSON(202, p)
}

func PaymentMethodsController(g *echo.Group) {
	c := paymentMethodsController{}
	g.GET("/:user_id", c.GET)
	g.PUT("/:user_id", c.PUT)
}
