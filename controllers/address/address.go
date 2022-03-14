package address

import (
	"database/sql"
	"fmt"
	"net/http"
	"server/key"
	"server/models/address"

	"github.com/labstack/echo"
)

func Privet(g *echo.Group) {
	g.GET("/province", GetProvince)
	g.GET("/amper", GetAmper)
	g.GET("/district", GetDistrict )
}

func GetProvince(c echo.Context) error {

	db, err := sql.Open("mysql", key.Database)
	if err != nil {
		fmt.Println("get tb_province", err)
	}

	defer db.Close()

	resultes, err := db.Query("SELECT * FROM tb_province ")

	if err != nil {
		fmt.Println("result err", err)
	}

	var TBProvince []address.Province
	for resultes.Next() {
		var TBProvinceArr address.Province

		err = resultes.Scan(
			&TBProvinceArr.Id,
			&TBProvinceArr.Name,
		)

		if err != nil {
			fmt.Println("scan user error ->", err.Error())
		}

		TBProvince = append(TBProvince, TBProvinceArr)

	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"result": TBProvince,
	})

}

func GetAmper(c echo.Context) error {

	ProvinceID := c.FormValue("id")

	db, err := sql.Open("mysql", key.Database)
	if err != nil {
		fmt.Println("Database Amper Error", err)
	}
	defer db.Close()

	var amper_TB = new(address.Amper)

	err = db.QueryRow("SELECT * FROM tb_amper WHERE id_province=? ", ProvinceID).Scan(&amper_TB.Id, &amper_TB.Name, &amper_TB.Id_Province)

	if err != nil {
		fmt.Println("error", err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"result": amper_TB,
	})

}

func GetDistrict(c echo.Context) error {

	DistrictID := c.FormValue("id")

	db, err := sql.Open("mysql", key.Database)
	if err != nil {
		fmt.Println("Database district Error", err)
	}
	defer db.Close()

	var district_TB = new(address.District)

	err = db.QueryRow("SELECT * FROM tb_district WHERE id_amper=? ", DistrictID).Scan(&district_TB.Id, &district_TB.Name, &district_TB.Id_Amper,&district_TB.Zipcode)

	if err != nil {
		fmt.Println("error", err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"result": district_TB,
	})

}
