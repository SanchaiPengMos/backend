package main

import (
	"server/controllers/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// const (
// 	username = "root"
// 	password = "root"
// 	hostname = "127.0.0.1:3306"
// 	dbname   = "projectasd"
// )

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	_publicAPI := e.Group("/public/api/v1")

	user.Privet(_publicAPI)

	e.Start(":9999")

}
