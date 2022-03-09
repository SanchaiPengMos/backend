package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"server/key"
	"server/models/user"

	"github.com/labstack/echo"
)



func Privet(g *echo.Group) {

	g.GET("/user/getuser", GetUser)
}

func GetUser(c echo.Context) error {

	db, err := sql.Open("mysql", key.Database)

	if err != nil {
		fmt.Println("error ->", err)
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM User_test")

	if err != nil {
		fmt.Println("result error ->", err)
	}
	
	var UserLog []user.User
	for results.Next() {

		var UserT user.User

		err = results.Scan(&UserT.Id, &UserT.Firstname, &UserT.Lastname)

		if err != nil {
			fmt.Println("scan user error ->", err.Error())
		}

		UserLog = append(UserLog, UserT)

	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"result": UserLog,
	})

}
