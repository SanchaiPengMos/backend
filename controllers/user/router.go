package user

import "github.com/labstack/echo"

func Private(g *echo.Group) {
	g.GET("/user/all_user", AllUser)
	g.GET("/user/:id", GetUser)
	g.POST("/user/edit_user/:id", EditUser)

}
