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
	// password := data.Password

	// fmt.Println(email)
	// fmt.Println(data)

	db, err := sql.Open("mysql", key.Database)
	if err != nil {
		fmt.Println("Login Error ", err)
	}
	defer db.Close()

	err = db.QueryRow("SELECT id,email, password FROM tb_user WHERE email=?", email).Scan(&data.Id, &data.Email, &data.Password)

	if err != nil {
		fmt.Println("err scan", err)
		return c.JSON(http.StatusOK, echo.Map{
			"status": false,
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	_claims := token.Claims.(jwt.MapClaims)
	_claims["exp"] = time.Now().Add(time.Minute * 180).Unix()
	_claims["id"] = data.Id

	tokenString, err := token.SignedString([]byte(key.TOKEN_SECRET))
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

func ValidateToken(token string) (aut.JwtClaims, error) {

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])

		}
		return []byte(key.TOKEN_SECRET), nil
	})

	var result aut.JwtClaims
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		result.Id = int(claims["id"].(float64))
		return result, nil

	} else {
		return result, err
	}

}
