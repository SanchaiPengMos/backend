package management

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"server/key"
	"server/models/shop"
	validatetoken "server/validateToken"
	"strings"

	"github.com/labstack/echo"
)

func Pubilc(g *echo.Group) {
	g.GET("/search/management/:id", SearchShop)
	
}

func Private(g *echo.Group) {
	g.PUT("/edit/shop", EditShop)
}

func ShopManagement(c echo.Context) error {
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

	id_user := string(decodeToken.Id)

	id := c.Param(id_user)

	stmt, err := db.Prepare("SELECT id ,id_location,tel,image,nameshop,address,id_address,id_amper,id_distric,create_date,last_update,lat,lng,id_user FROM tb_shop WHERE id = $1")
	if err != nil {
		log.Println(err)
	}
	var data shop.ShopManagement

	var shopmanagement []shop.ShopManagement

	err = stmt.QueryRow(id).Scan(
		&data.ID,
		// &data.Id_Location,
		&data.Tel,
		&data.Img,
		&data.Nameshop,
		&data.Address,
		&data.IDAddress,
		&data.IDAmper,
		&data.Id_district,
		&data.Create_date,
		&data.Last_update,
		&data.Lat,
		&data.Long,
		&data.Id_User,
	)
	if err != nil {
		log.Println("Scan failed:", err.Error())
	}

	shopmanagement = append(shopmanagement, data)

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"result": shopmanagement,
	})

}

func SearchShop(c echo.Context) error {
	db, err := sql.Open("mysql", key.Database)
	if err != nil {
		fmt.Println("error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": false,
		})
	}
	defer db.Close()

	id := c.Param("id")

	stmt, err := db.Query("SELECT tshop.id,tshop.tel,tshop.image,tshop.nameshop,tshop.address,tshop.province,tshop.amper,tshop.district,tshop.create_date,tshop.last_update,tshop.lat,tshop.lng,tshop.id_user,tshop.is_open FROM tb_shop as tshop WHERE id_user = ?", id)
	if err != nil {
		log.Println(err)
	}

	
	var OBJshopmanagement shop.ShopManagement

	for stmt.Next() {
		// var data  shop.ShopManagement
		err := stmt.Scan(
			&OBJshopmanagement.ID,
			&OBJshopmanagement.Tel,
			&OBJshopmanagement.Img,
			&OBJshopmanagement.Nameshop,
			&OBJshopmanagement.Address,
			&OBJshopmanagement.IDAddress,
			&OBJshopmanagement.IDAmper,
			&OBJshopmanagement.Id_district,
			&OBJshopmanagement.Create_date,
			&OBJshopmanagement.Last_update,
			&OBJshopmanagement.Lat,
			&OBJshopmanagement.Long,
			&OBJshopmanagement.Id_User,
			&OBJshopmanagement.Is_Open,
		)
		if err != nil {
			log.Println("Scan failed:", err.Error())
		}
	}

	fmt.Println("OBJshopmanagement ID :", OBJshopmanagement.ID )

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"result": OBJshopmanagement,
	})
}
