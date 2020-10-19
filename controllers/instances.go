package controllers

import (
	"strconv"

	"github.com/juliotorresmoreno/macabro/db"
	"github.com/juliotorresmoreno/macabro/helper"
	"github.com/juliotorresmoreno/macabro/models"
	"github.com/labstack/echo"
)

type instancesController struct {
}

func (el *instancesController) GETALL(c echo.Context) error {
	session, err := validateSession(c)
	if err != nil {
		return err
	}

	instances := make([]*models.InstanceWithDate, 0)
	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}
	defer conn.Close()

	if err = conn.Find(&instances); err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}

	return c.JSON(200, instances)
}

func (el *instancesController) GET(c echo.Context) error {
	session, err := validateSession(c)
	if err != nil {
		return err
	}

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
	session, err := validateSession(c)
	if err != nil {
		return err
	}

	instanceUpdate := new(models.Instance)
	instanceActual := new(models.Instance)
	conn, err := db.NewEngigneWithSession(session.Username, session.ACL.Group)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}
	defer conn.Close()
	id, _ := strconv.Atoi(c.Param("id"))
	instanceActual.ID = id
	_, err = conn.Get(instanceActual)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}

	if err = c.Bind(instanceUpdate); err != nil {
		return echo.NewHTTPError(406, helper.ParseError(err))
	}
	instanceActual.User = session
	instanceActual.IsCloud = instanceUpdate.IsCloud
	instanceActual.Name = instanceUpdate.Name
	instanceActual.Type = instanceUpdate.Type
	instanceActual.Replicas = instanceUpdate.Replicas
	instanceActual.AutoScaling = instanceUpdate.AutoScaling
	instanceActual.AllowDeletion = instanceUpdate.AllowDeletion
	instanceActual.BackupPeriodicity = instanceUpdate.BackupPeriodicity
	instanceActual.URL = instanceUpdate.URL
	instanceActual.Username = instanceUpdate.Username
	instanceActual.Password = instanceUpdate.Password

	if err = instanceActual.Check(); err != nil {
		return echo.NewHTTPError(406, helper.ParseError(err))
	}
	_, err = conn.Where("id = ?", id).Update(instanceActual)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err))
	}

	return c.String(204, "")
}

func (el *instancesController) PUT(c echo.Context) error {
	session, err := validateSession(c)
	if err != nil {
		return err
	}

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

	return c.String(204, "")
}

func (el *instancesController) DELETE(c echo.Context) error {
	session, err := validateSession(c)
	if err != nil {
		return err
	}

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

	return c.String(204, "")
}

func InstancesController(g *echo.Group) {
	c := &instancesController{}

	g.PUT("", c.PUT)
	g.GET("", c.GETALL)
	g.GET("/:id", c.GET)
	g.POST("/:id", c.POST)
	g.DELETE("/:id", c.DELETE)
}
