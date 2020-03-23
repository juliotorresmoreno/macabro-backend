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

func (that paymentMethodsController) GETALL(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*models.User)
	userID := c.QueryParam("user_id")
	if !session.ACL.IsAdmin() && strconv.Itoa(session.ID) != userID {
		return echo.NewHTTPError(401, "Unauthorized")
	}

	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	defer conn.Close()

	p := make([]*models.PaymentMethods, 0)
	err = conn.Where("user_id = ?", userID).Find(&p)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	for _, c := range p {
		c.Decrypt(session.Username)
	}

	return c.JSON(200, p)
}

func (that paymentMethodsController) PUT(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*models.User)

	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	defer conn.Close()

	p := &models.PaymentMethods{}
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	if p.UserID == 0 {
		p.UserID = session.ID
	}
	if !session.ACL.IsAdmin() && session.ID != p.UserID {
		return echo.NewHTTPError(401, "Unauthorized")
	}

	s, _ := conn.Where("user_id = ?", p.UserID).Table(p.TableName()).Count()
	if s == 0 {
		p.Default = 1
	}

	p.Encrypt(session.Username)
	_, err = conn.InsertOne(p)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}

	return c.JSON(202, p)
}

func (that paymentMethodsController) PATCH(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*models.User)

	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	defer conn.Close()

	p := &models.PaymentMethods{}
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	if p.Default == 0 {
		return c.String(204, "")
	}
	id := c.ParamValues()[0]

	dbSession := conn.NewSession()
	dbSession.Begin()
	_, err = dbSession.
		Table(p.TableName()).
		Where("user_id = ? and id != ?", session.ID, id).
		Update(map[string]interface{}{"default": 0})
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	_, err = conn.Where("id = ?", id).
		Table(p.TableName()).
		Update(map[string]interface{}{"default": p.Default})

	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	dbSession.Commit()

	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}

	return c.String(204, "")
}

func (that paymentMethodsController) DELETE(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*models.User)

	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	defer conn.Close()
	id := c.ParamValues()[0]

	p := &models.PaymentMethods{}
	conn.Where("id = ?", id).Get(p)
	if !session.ACL.IsAdmin() && session.ID != p.UserID {
		return echo.NewHTTPError(401, "Unauthorized")
	}

	if p.Default == 1 {
		return echo.NewHTTPError(406, "No se puede eliminar el metodo de pago por defecto")
	}

	conn.Where("id = ? and \"default\" = 0", id).Delete(models.PaymentMethods{})

	return c.String(204, "")
}

func PaymentMethodsController(g *echo.Group) {
	c := paymentMethodsController{}
	g.GET("", c.GETALL)
	g.PUT("", c.PUT)
	g.PATCH("/:id", c.PATCH)
	g.DELETE("/:id", c.DELETE)
}
