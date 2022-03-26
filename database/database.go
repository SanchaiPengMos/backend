package database

import (
	"database/sql"
	"fmt"
	"server/key"

	"github.com/labstack/echo"
)

func Database(c echo.Context) {
	db, err := sql.Open("mysql", key.Database)
	if err != nil {
		fmt.Println("error", err)
	}
	defer db.Close()

}
