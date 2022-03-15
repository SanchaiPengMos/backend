package register

import (
	"database/sql"
	"fmt"
	"net/http"
	"server/key"
	"server/models/register"

	"github.com/labstack/echo"
)

func Privet(g *echo.Group) {
	g.POST("/register", Register )
}

func Register(c echo.Context) error {

	data := new(register.Register)

	if err := c.Bind(data); err != nil {
		fmt.Println("err ->", err)
		return c.JSON(http.StatusOK, echo.Map{
			"status": false,
		})
	}

	db, err := sql.Open("mysql", key.Database)

	if err != nil {
		fmt.Println("error ->", err)
	}

	defer db.Close()

	query := "INSERT INTO tb_shop(id_location, tel,image,nameshop,address,id_address,id_amper,id_district,create_date,last_update) VALUES ('$1', $2 ,$3,$4,$5,$6,$7,$8,$9,$10 )"

	res, err := db.Exec(query)

	if err != nil {
		fmt.Println("err", err)
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		fmt.Println("err", err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"result": affectedRows,
	})

	//INSERT INTO `tb_shop` (`id`, `id_location`, `tel`, `image`, `nameshop`, `address`, `id_address`, `create_date`, `last_update`) VALUES (NULL, '1', '0948381309', 'Sanchai', 'Pengboot', '240/7', '1', '2022-03-15', '2022-03-15');

}
