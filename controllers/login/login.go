package login

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/key"
	"server/models/aut"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Pubilc(g *echo.Group) {
	g.POST("/login", Login)
}

func Login(c echo.Context) error {

	data := new(aut.Login)

	

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println("body login err", err)
	}
	json.Unmarshal(body, &data)

	email := data.Email

	
	fmt.Println("data :", data)
	fmt.Println("ee", data.Email )

	db, err := sql.Open("mysql", key.Database)
	if err != nil {
		fmt.Println("Login Error ", err)
	}
	defer db.Close()

	err = db.QueryRow("SELECT id,email, password FROM tb_user WHERE email=?", email).Scan(&data.Id, &data.Email, &data.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": false,
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	_claims := token.Claims.(jwt.MapClaims)
	_claims["exp"] = time.Now().Add(time.Minute * 180).Unix()
	_claims["id"] = data.Id

	tokenString, err := token.SignedString([]byte(key.TOKEN_SECRET))

	// fmt.Println(_claims)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  false,
			"message": "Can't create token.",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"result": tokenString,
	})

}
