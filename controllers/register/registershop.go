package register

import (
	"database/sql"
	"fmt"
	"net/http"
	"server/key"
	validatetoken "server/validateToken"
	"strings"

	"github.com/labstack/echo"
)

func RegisterShop(c echo.Context) error {

	db, err := sql.Open("mysql", key.Database)
	if err != nil {
		fmt.Println("error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": false,
		})
	}
	defer db.Close()

	token := c.Request().Header.Get("Authorization")

	jwtString := strings.Split(token, "Bearer ")[1]

	decodeToken, err := validatetoken.ValidateToken(jwtString)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": false,
		})
	}

	fmt.Println(decodeToken.Id)

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
	})

}
