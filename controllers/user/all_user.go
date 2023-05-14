package user

import (
	connectdb "backend/connectDB"
	"backend/models"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func AllUser(c echo.Context) (err error) {
	fmt.Println("test")
	rows, err := connectdb.SqlDB.Query("SELECT id , email,firstname,lastname,userpermis_id FROM userhotel order by id asc")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var CheckUse []models.UserCheck

	for rows.Next() {
		var data models.UserCheck
		err := rows.Scan(&data.Id, &data.Email, &data.Firstname, &data.Lastname, &data.Userpermis_id)
		if err != nil {
			log.Println("Scan failed:", err.Error())
		}
		CheckUse = append(CheckUse, data)
	}

	return c.JSON(http.StatusOK, CheckUse)
}
