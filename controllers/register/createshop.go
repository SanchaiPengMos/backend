package register

import (
	"database/sql"
	"fmt"
	"server/key"

	"github.com/labstack/echo"
)

func createShop(c echo.Context, ID uint) {
	fmt.Println(ID)
	db, err := sql.Open("mysql", key.Database)

	if err != nil {
		fmt.Println("error ->", err)
	}

	defer db.Close()

	abc := 0
	is_OP := 0

	stmt, err := db.Prepare("INSERT INTO tb_shop(id,id_user,is_open) VALUES(?,?,?)")

	if err != nil {
		fmt.Println(err)
	}

	_, err = stmt.Exec(
		abc,
		ID,
		is_OP,
	)
	if err != nil {
		fmt.Println("err", err.Error())

	}
}
