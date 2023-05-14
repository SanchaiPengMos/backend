package login

import (
	connectdb "backend/connectDB"
	"backend/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) (err error) {

	data := new(models.Userhotel)

	//get email, passwrod form request(body)
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(body, &data)

	email := data.Email
	password := data.Password

	//select password by email
	rows, err := connectdb.SqlDB.Prepare("SELECT id, email, password, userpermis_id,firstname  FROM userhotel WHERE email=$1  ")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	//for rows.Next() { //Next ดึงหลาย rows  ถ้าจะดึงrows เดียวใช้อันไหน

	//แสดงข้อมูลที่ selcet ออกมา
	err = rows.QueryRow(email).Scan(&data.Id, &data.Email, &data.Password, &data.Userpermis_id, &data.Firstname)
	if err != nil {
		log.Println("Scan failed:", err.Error())
	}

	//verify password by bcrypt
	// data.Password มาเช็คว่า password เป็นค่าเดียวกันหรือไม่
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password))
	if err != nil {
		return c.String(http.StatusUnauthorized, "Your Email or Password were wrong")
	}

	// check email and password against DB after hashing the password
	if data.Userpermis_id == 1 {
		log.Println(data.Firstname)
		// create jwt token
		token, err := createJwtToken(int(data.Id)) //ควรเก็บ id and userpermis
		if err != nil {
			log.Println("Error Creating JWT token", err)
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		cookie := new(http.Cookie)
		cookie.Name = "cookie_login"
		cookie.Value = token
		cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.HttpOnly = true

		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, map[string]string{
			"firstname": data.Firstname,
			"email":     data.Email,
			"type":      "Admin",
			"message":   "success",
			"token":     token,
			"tel":       data.Tel,
			"birthday":  data.Birthday,
		})

	} else if data.Userpermis_id == 2 {

		// create jwt token

		token, err := createJwtToken(int(data.Id)) //ควรเก็บ id and userpermis
		if err != nil {
			log.Println("Error Creating JWT token", err)
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		cookie := new(http.Cookie)
		cookie.Name = "cookie_login"
		cookie.Value = token
		cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.HttpOnly = true
		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, map[string]string{
			"firstname": data.Firstname,
			"email":     data.Email,
			"type":      "Admin",
			"message":   "success",
			"token":     token,
			"tel":       data.Tel,
			"birthday":  data.Birthday,
		})

	}

	return c.String(http.StatusUnauthorized, "Your Email or Password were wrong")
}
