package user

type User struct {
	Id        uint   `sql:"primary_key" json:"id" form:"id" query:"id"`
	Firstname string `json:"f_name" form:"f_name" query:"f_name"`
	Lastname  string `json:"l_name" form:"l_name" query:"l_name"`
}