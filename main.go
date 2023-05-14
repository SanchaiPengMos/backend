package main

import (
	connectdb "backend/connectDB"
	"backend/controllers/booking"
	"backend/controllers/hotel"
	"backend/controllers/login"
	"backend/controllers/news"
	"backend/controllers/register"
	"backend/controllers/user"
	"backend/key"
	"backend/models"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {

	connectdb.Initialize()

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	var IsToken = middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &models.JwtClaims{},
		SigningKey: []byte(key.TOKEN_SECRET),
	})

	_publicAPI := e.Group("/public/api/v1")
	_PrivateAPI := e.Group("/private/api/v1", IsToken)

	booking.Private(_PrivateAPI)
	hotel.Private(_PrivateAPI)
	login.Pubilc(_publicAPI)
	news.Private(_PrivateAPI)
	register.Pubilc(_publicAPI)
	user.Private(_PrivateAPI)

	e.Start(":9999")

}
