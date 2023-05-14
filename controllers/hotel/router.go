package hotel

import "github.com/labstack/echo"

func Private(g *echo.Group) {
	g.GET("/hotel/select_room", SelectRoom)
}
