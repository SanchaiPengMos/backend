package booking

import (
	connectdb "backend/connectDB"
	"backend/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func CountBooking(c echo.Context) (err error) {

	var data models.Total

	rows, err := connectdb.SqlDB.Prepare("SELECT booking.id,(DATE_PART('day', create_to::date) - DATE_PART('day', create_from::date))*room.price as total FROM booking AS booking LEFT JOIN room AS room ON room.id = booking.room_id ORDER BY booking.id DESC LIMIT 1")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	var userhotels []models.Total

	err = rows.QueryRow().Scan(&data.Id, &data.Total)
	if err != nil {
		log.Println("Scan failed:", err.Error())
	}

	userhotels = append(userhotels, data)

	return c.JSON(http.StatusOK, userhotels)
}
