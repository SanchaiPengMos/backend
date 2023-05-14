package register

import (
	connectdb "backend/connectDB"
	"backend/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) (err error) {

	data := new(models.Userhotel)

	if err = c.Bind(data); err != nil {

		return err
	}
	//เพิ่มข้อมูล user
	stmt, err := connectdb.SqlDB.Prepare("INSERT INTO userhotel(email,password,firstname,lastname,birthday,tel,userpermis_id) VALUES($1,$2,$3,$4,$5,$6,2)")

	if err != nil {

		log.Println("Prepare failed:", err.Error())

	}
	//แปลง
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {

		log.Println("Error hashedPassword : ", err)

	}

	//Export data user
	_, err = stmt.Exec(data.Email, hashedPassword, data.Firstname, data.Lastname, data.Birthday, data.Tel)

	if err != nil {
		log.Println("DATABASE INSERT Error :", err.Error())
	}

	defer stmt.Close()

	return c.JSON(http.StatusOK, data)

}
