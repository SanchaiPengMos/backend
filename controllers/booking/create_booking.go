package booking

import (
	connectdb "backend/connectDB"
	"backend/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Booking(c echo.Context) (err error) {

	data := new(models.Booking)

	if err = c.Bind(data); err != nil {
		return err
	}

	stmt, err := connectdb.SqlDB.Prepare("INSERT INTO booking(create_from,create_to,room_id,user_id) VALUES($1,$2,$3,$4)")

	if err != nil {

		log.Println("Prepare failed:", err.Error())

	}
	//Export data user
	_, err = stmt.Exec(data.Create_from, data.Create_to, data.Room_id, data.User_id)

	if err != nil {
		log.Println("DATABASE INSERT Error :", err.Error())
	}

	defer stmt.Close()

	return c.JSON(http.StatusOK, data)
}
