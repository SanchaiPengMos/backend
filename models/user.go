package models

type Userhotel struct {
	Id            uint   `sql:"primary_key" json:"id" form:"id" query:"id"`
	Email         string `json:"email" form:"email" query:"email"`
	Password      string `json:"password" form:"password" query:"password"`
	Firstname     string `json:"firstname" form:"firstname" query:"firstname"`
	Lastname      string `json:"lastname" form:"lastname" query:"lastname"`
	Birthday      string `json:"birthday" form:"birthday" query:"birthday"`
	Tel           string `json:"tel" form:"tel" query:"tel"`
	Create_by     int    `json:"create_by" form:"create_by" query:"create_by"`
	Update_by     int    `json:"update_by" form:"update_by" query:"update_by"`
	Userpermis_id int    `json:"userpermis_id" form:"userpermis_id" query:"userpermis_id"`
}

type UserCheck struct {
	Id            uint   `sql:"primary_key" json:"id" form:"id" query:"id"`
	Email         string `json:"email" form:"email" query:"email"`
	Firstname     string `json:"firstname" form:"firstname" query:"firstname"`
	Tel           string `json:"tel" form:"tel" query:"tel"`
	Lastname      string `json:"lastname" form:"lastname" query:"lastname"`
	Birthday      string `json:"birthday" form:"birthday" query:"birthday"`
	Userpermis_id int    `json:"userpermis_id" form:"userpermis_id" query:"userpermis_id"`
}