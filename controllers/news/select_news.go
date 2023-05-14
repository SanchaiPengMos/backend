package news

import (
	connectdb "backend/connectDB"
	"backend/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func NewsTopic(c echo.Context) (err error) {

	rows, err := connectdb.SqlDB.Query("SELECT id , topic , detail FROM news order by id ASC ")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	var CheckUse []models.Newstopic

	for rows.Next() {
		var data models.Newstopic
		err := rows.Scan(&data.Id, &data.Topic, &data.Detail)
		if err != nil {
			log.Println("Scan failed:", err.Error())
		}
		CheckUse = append(CheckUse, data)
	}

	return c.JSON(http.StatusOK, CheckUse)
}
