package register

import (
	"database/sql"
	"fmt"
	"net/http"
	"server/key"
	"server/models/register"

	"github.com/labstack/echo"
)

func Pubilc(g *echo.Group) {
	g.POST("/register", Register)
}

func Register(c echo.Context) error {

	db, err := sql.Open("mysql", key.Database)

	if err != nil {
		fmt.Println("error ->", err)
	}

	defer db.Close()

	data := new(register.RegisterUser)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": false,
		})
	}
	stmt, err := db.Prepare("INSERT INTO tb_user(id,fname,lname,email,password,role_id) VALUES(?,?,?,?,?,?)")

	if err != nil {
		fmt.Println(err)
	}

	_, err = stmt.Exec(data.ID, data.Fname, data.Fname, data.Email, data.Password, data.Role_id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  false,
			"message": "ชื่อผู้ใช้งานนี้มีอยู่ในระบบแล้ว",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
	})
}
