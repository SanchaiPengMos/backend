package address

type District struct {
	Id       uint   `sql:"primary_key" json:"id" form:"id" query:"id"`
	Name     string `json:"name" form:"name" query:"name"`
	Id_Amper int    `json:"id_amper" form:"id_amper" query:"id_amper"`
	Zipcode  int    `json:"zipcode" form:"zipcode" query:"zipcode"`
}

type Amper struct {
	Id          uint   `sql:"primary_key" json:"id" form:"id" query:"id"`
	Name        string `json:"name" form:"name" query:"name"`
	Id_Province int    `json:"id_province" form:"id_province" query:"id_province"`
}

type Province struct {
	Id   uint   `sql:"primary_key" json:"id" form:"id" query:"id"`
	Name string `json:"name" form:"name" query:"name"`
}
