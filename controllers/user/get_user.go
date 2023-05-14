package user

import (
	connectdb "backend/connectDB"
	"backend/controllers/login"
	"backend/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func GetUser(c echo.Context) (err error) {

	token := c.Request().Header.Get("Authorization")

	decodeToken, err := login.ValidateToken(token)
	if err != nil {
		// err
		return c.JSON(http.StatusBadRequest, err)
	}

	userID := decodeToken.Id

	rows, err := connectdb.SqlDB.Query("SELECT id , email,firstname,lastname,birthday,tel,userpermis_id FROM userhotel WHERE id = $1", userID)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	var CheckUse []models.UserCheck

	for rows.Next() {
		var data models.UserCheck
		err := rows.Scan(&data.Id, &data.Email, &data.Firstname, &data.Lastname, &data.Birthday, &data.Tel, &data.Userpermis_id)
		if err != nil {
			log.Println("Scan failed:", err.Error())
		}
		CheckUse = append(CheckUse, data)
	}

	return c.JSON(http.StatusOK, CheckUse)
}
