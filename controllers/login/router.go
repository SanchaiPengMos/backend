package login

import "github.com/labstack/echo"

func Pubilc(g *echo.Group) {
	g.POST("/login", Login)
}
