package user

import "database/sql"

type User struct {
	Id         uint         `sql:"primary_key" json:"id" form:"id" query:"id"`
	Firstname  string       `json:"fname" form:"fname" query:"fname"`
	Lastname   string       `json:"lname" form:"lname" query:"lname"`
	Username   string       `json:"username" form:"username" query:"username"`
	Pass       string       `json:"password" form:"password" query:"password"`
	Shop       int          `json:"shop_id" form:"shop_id" query:"shop_id"`
	Role       int          `json:"role_id" form:"role_id" query:"role_id"`
	CreateDate sql.NullTime `json:"created_date" form:"created_date" query:"created_date"`
}
