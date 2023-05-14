package models

type Room struct {
	Id             uint   `sql:"primary_key" json:"id" form:"id" query:"id"`
	Create_by_user int    `json:"create_by_user" form:"create_by_user" query:"create_by_user"`
	Name_room      string `json:"name_room" form:"name_room" query:"name_room"`
	Room_type_id   int    `json:"room_type_id" form:"room_type_id" query:"room_type_id"`
	Price          int    `json:"price" form:"price" query:"price"`
	Detail         string `json:"detail" form:"detail" query:"detail"`
}

type Total struct {
	Id    int `sql:"primary_key" json:"id" form:"id" query:"id"`
	Total int `json:"total" form:"total" query:"total"`
}
