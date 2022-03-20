package main

import (
	"net/http"
	"server/controllers/address"
	"server/controllers/location"
	"server/controllers/login"
	"server/controllers/register"
	"server/controllers/user"
	"server/key"
	"server/models/aut"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${remote_ip} ${user_agent} status=${status} error:${error}\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Use(middleware.Recover())

	var IsToken = middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &aut.JwtClaims{},
		SigningKey: []byte(key.TOKEN_SECRET),
	})

	_publicAPI := e.Group("/public/api/v1")
	_PrivateAPI := e.Group("/private/api/v1", IsToken)

	login.Pubilc(_publicAPI)
	address.Pubilc(_publicAPI)
	register.Pubilc(_publicAPI)

	user.Privet(_PrivateAPI)
	location.Privet(_PrivateAPI)

	e.Start(":9999")

}
