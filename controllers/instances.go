package controllers

import (
	"github.com/juliotorresmoreno/macabro/db"
	"github.com/juliotorresmoreno/macabro/helper"
	"github.com/juliotorresmoreno/macabro/models"
	"github.com/labstack/echo"
)

type instancesController struct {
}

func (el *instancesController) GETALL(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*models.User)

	u := make([]models.Instance, 0)
	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}
	defer conn.Close()

	if err = conn.Find(&u); err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}

	return c.JSON(200, u)
}

func (el *instancesController) GET(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*models.User)

	u := new(models.Instance)
	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}
	defer conn.Close()

	id := c.Param("id")
	_, err = conn.Where("id = ?", id).Get(u)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}

	return c.JSON(200, u)
}

func (el *instancesController) POST(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*models.User)

	u := new(models.Instance)
	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}
	defer conn.Close()

	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(406, helper.ParseError(err))
	}
	u.User = session

	if err = u.Check(); err != nil {
		return echo.NewHTTPError(406, helper.ParseError(err))
	}
	id := c.Param("id")
	_, err = conn.Where("id = ?", id).Update(u)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}

	return c.JSON(200, u)
}

func (el *instancesController) PUT(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*models.User)

	u := new(models.Instance)
	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}
	defer conn.Close()

	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(406, helper.ParseError(err))
	}
	u.User = session

	if err = u.Check(); err != nil {
		return echo.NewHTTPError(406, helper.ParseError(err))
	}
	_, err = conn.Insert(u)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}

	return c.JSON(200, u)
}

func (el *instancesController) DELETE(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*models.User)

	u := new(models.Instance)
	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}
	defer conn.Close()

	id := c.Param("id")
	_, err = conn.Where("id = ?", id).Delete(u)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}

	return c.JSON(200, u)
}

func InstancesController(g *echo.Group) {
	c := &instancesController{}

	g.PUT("", c.PUT)
	g.GET("", c.GETALL)
	g.GET("/:id", c.GET)
	g.POST("/:id", c.POST)
	g.DELETE("/:id", c.DELETE)
}
