package shop

type ShopManagement struct {
	ID          int     `sql:"primary_key" json:"ç" form:"id" query:"id"`
	Tel         string  `json:"tel" form:"tel" query:"tel"`
	Img         string  `json:"image" form:"image" query:"image"`
	Nameshop    string  `json:"nameshop" form:"nameshop" query:"nameshop"`
	Address     string  `json:"address" form:"address" query:"address"`
	IDAddress   string     `json:"province" form:"province" query:"province"`
	IDAmper     string     `json:"amper" form:"amper" query:"amper"`
	Id_district string     `json:"district" form:"district" query:"district"`
	Create_date string  `json:"create_date" form:"create_date" query:"create_date"`
	Last_update string  `json:"last_update" form:"last_update" query:"last_update"`
	Lat         float64 `json:"lat" form:"lat" query:"lat"`
	Long        float64 `json:"lng" form:"lng" query:"lng"`
	Id_User     float64 `json:"id_user" form:"id_user" query:"id_user"`
	Is_Open     string  `json:"is_open" form:"is_open" query:"is_open"`
}
