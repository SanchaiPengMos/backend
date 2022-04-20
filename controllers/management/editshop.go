package management

import (
	"database/sql"
	"fmt"
	"net/http"
	"server/key"
	"server/models/shop"
	validatetoken "server/validateToken"
	"strings"

	"github.com/labstack/echo"
)

func EditShop(c echo.Context) error {
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

	data := new(shop.ShopManagement)

	if err := c.Bind(data); err != nil {
		return err
	}

	idtoken := string(decodeToken.Id)

	id := c.Param(idtoken)

	fmt.Println("data::::",data)

	stmt, err := db.Prepare("UPDATE tb_shop SET tel=$1, image=$2,nameshop=$3,address=$4,province=$5,amper=$6,district=$7,lat=$8,lng=$9 ,id_user=10, is_open=$11 WHERE id_user= $12 ")

	if err != nil {
		fmt.Println("err :", err.Error())
		return err
	}
	_, err = stmt.Exec(data.Tel, data.Img, data.Nameshop, data.Address, data.IDAddress, data.IDAmper, data.Id_district, data.Lat, data.Long, id)

	if err != nil {
		return err
	}

	
	return c.JSON(http.StatusOK, echo.Map{
		"status":  true,
		// "restult": data,
	})
}
