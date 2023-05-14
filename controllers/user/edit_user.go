package user

import (
	connectdb "backend/connectDB"
	"backend/models"
	"net/http"

	"github.com/labstack/echo"
)

func EditUser(c echo.Context) (err error) {

	data := new(models.Userhotel)

	if err := c.Bind(data); err != nil {
		return err
	}

	id := c.Param("id")

	stmt, err := connectdb.SqlDB.Prepare("UPDATE userhotel SET firstname=$1, lastname=$2,tel=$3 WHERE id= $4 ")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(data.Firstname, data.Lastname, data.Tel, id)

	if err != nil {
		return err
	}

	defer stmt.Close()

	return c.JSON(http.StatusOK, data)
}
