package models

type Newstopic struct {
	Id     int    `sql:"primary_key" json:"id" form:"id" query:"id"`
	Topic  string `json:"topic" form:"topic" query:"topic"`
	Detail string `json:"detail" form:"detail" query:"detail"`
}
