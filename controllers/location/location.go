package location

import (
	"database/sql"
	"fmt"
	"net/http"
	"server/key"

	// "server/models/location"
	"server/models/shop"

	"github.com/labstack/echo"
)

func Privet(g *echo.Group) {
	g.GET("/location", GetLocation)
}

func GetLocation(c echo.Context) error {

	db, err := sql.Open("mysql", key.Database)
	if err != nil {
		fmt.Println("Location err ", err)
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM tb_shop WHERE is_open = 1 ")

	if err != nil {
		fmt.Println("Result err ", err)
	}

	var LocationA []shop.ShopManagement

	for result.Next() {

		// var Locat location.Location
		var ABCCC shop.ShopManagement

		err = result.Scan(
			&ABCCC.ID,
			// &ABCCC.Id_Location,
			&ABCCC.Tel,
			&ABCCC.Img,
			&ABCCC.Nameshop,
			&ABCCC.Address,
			&ABCCC.IDAddress,
			&ABCCC.IDAmper,
			&ABCCC.Id_district,
			&ABCCC.Create_date,
			&ABCCC.Last_update,
			&ABCCC.Lat,
			&ABCCC.Long,
			&ABCCC.Id_User,
			&ABCCC.Is_Open,
		)

		if err != nil {
			fmt.Println("scan location err ", err)
		}

		LocationA = append(LocationA, ABCCC)

	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"result": LocationA,
	})

}
