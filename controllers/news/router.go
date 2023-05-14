package news

import "github.com/labstack/echo"

func Private(g *echo.Group) {
	g.GET("/news/select_news", NewsTopic)
}
