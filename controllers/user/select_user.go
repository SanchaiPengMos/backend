package user

// import (
// 	connectdb "backend/connectDB"
// 	"backend/models"
// 	"log"
// 	"net/http"

// 	"github.com/labstack/echo"
// )

// func UserBooking(c echo.Context) (err error) {

// 	id := c.Param("id")

// 	var data models.Userhotel

// 	rows, err := connectdb.SqlDB.Prepare("SELECT id , email,password,firstname,lastname,birthday,tel,created_by,update_by,userpermis_id FROM userhotel WHERE id = $1")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	defer rows.Close()

// 	var userhotels []models.Userhotel

// 	err = rows.QueryRow(id).Scan(&data.Id, &data.Email, &data.Password, &data.Firstname, &data.Lastname, &data.Birthday, &data.Tel, &data.Create_by, &data.Update_by, &data.Userpermis_id)
// 	if err != nil {
// 		log.Println("Scan failed:", err.Error())
// 	}

// 	userhotels = append(userhotels, data)

// 	return c.JSON(http.StatusOK, userhotels)
// }
