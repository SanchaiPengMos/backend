package models

type Booking struct {
	ID          uint   `sql:"primary_key" json:"id" form:"id" query:"id"`
	User_id     int    `json:"User_id" form:"User_id" query:"User_id"`
	Create_from string `json:"create_from" form:"create_from" query:"create_from"`
	Create_to   string `json:"create_to" form:"create_to" query:"create_to"`
	Room_id     int    `json:"room_id" form:"room_id" query:"room_id"`
	Active_id   int    `json:"active_id" from:"active_id" query:"active_id"`
}