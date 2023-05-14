package hotel

import (
	connectdb "backend/connectDB"
	"backend/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func SelectRoom(c echo.Context) (err error) {

	rows, err := connectdb.SqlDB.Query("SELECT room.id , room.create_by_user,room.name_room,room.room_type_id,room.price ,room_type.detail FROM room AS room LEFT JOIN room_type AS room_type ON room_type.id = room.id ORDER BY room.id ASC")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	var CheckUse []models.Room

	for rows.Next() {
		var data models.Room
		err := rows.Scan(&data.Id, &data.Create_by_user, &data.Name_room, &data.Room_type_id, &data.Price, &data.Detail)
		if err != nil {
			log.Println("Scan failed:", err.Error())
		}
		CheckUse = append(CheckUse, data)

	}

	return c.JSON(http.StatusOK, CheckUse)
}
