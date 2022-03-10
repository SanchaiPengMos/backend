package location

import (
	"database/sql"
	"fmt"
	"net/http"
	"server/key"
	"server/models/location"

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

	result, err := db.Query("SELECT * FROM tb_location ")

	if err != nil {
		fmt.Println("Result err ", err)
	}

	var LocationA []location.Location

	for result.Next() {

		var Locat location.Location

		err = result.Scan(
			&Locat.Id,
			&Locat.Lat,
			&Locat.Long,
		)

		if err != nil {
			fmt.Println("scan location err ", err)
		}

		LocationA = append(LocationA, Locat)

	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"result": LocationA,
	})

}
