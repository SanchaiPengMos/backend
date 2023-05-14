package register

import "github.com/labstack/echo"

func Pubilc(g *echo.Group) {
	g.GET("/register", Register)
}
