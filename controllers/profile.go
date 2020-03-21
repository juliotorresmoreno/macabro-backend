package controllers

import (
	"github.com/labstack/echo"
)

type profileController struct {
}

func (that profileController) PATCH(c echo.Context) error {
	r := Response{"success":true}
	return c.JSON(200, r)
}


// ProfileController s
func ProfileController(g *echo.Group) {
	c := profileController{}
	g.PATCH("", c.PATCH)
}
