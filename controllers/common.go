package controllers

import (
	"github.com/juliotorresmoreno/macabro/models"
	"github.com/labstack/echo"
)

func validateSession(c echo.Context) (*models.User, error) {
	_session := c.Get("session")
	var session *models.User
	if _session == nil {
		return session, echo.NewHTTPError(401, "Unauthorized")
	}
	session = _session.(*models.User)
	return session, nil
}
