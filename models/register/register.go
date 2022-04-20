package register

// type Register struct {
// 	ID          uint
// 	Id_Location int
// 	Tel         string
// 	Image       string
// 	NameShop    string
// 	Address     string
// 	Id_address  int
// 	Id_amper    int
// 	Id_district int
// 	Create_date time.Time
// 	Last_Update time.Time
// }

type RegisterUser struct {
	ID       uint   `sql:"primary_key" json:"id" form:"id" query:"id"`
	Fname    string `json:"fname" form:"fname" query:"fname"`
	LName    string `json:"lname" form:"lname" query:"lname"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
	Role_id  uint   `json:"role_id" form:"role_id" query:"role_id"`
	// Create_date time.Time `json:"created_date" form:"created_date" query:"created_date"`
}

type RegisterUserA struct {
	ID       uint   `sql:"primary_key" json:"id" form:"id" query:"id"`
	Fname    string `json:"fname" form:"fname" query:"fname"`
	LName    string `json:"lname" form:"lname" query:"lname"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
	Role_id  uint   `json:"role_id" form:"role_id" query:"role_id"`
	// Create_date time.Time `json:"created_date" form:"created_date" query:"created_date"`
}

type RegisterShop struct {
	ID          int     `sql:"primary_key" json:"ç" form:"id" query:"id"`
	Tel         string  `json:"tel" form:"tel" query:"tel"`
	Img         string  `json:"image" form:"image" query:"image"`
	Nameshop    string  `json:"nameshop" form:"nameshop" query:"nameshop"`
	Address     string  `json:"address" form:"address" query:"address"`
	IDAddress   string  `json:"province" form:"province" query:"province"`
	IDAmper     string  `json:"amper" form:"amper" query:"amper"`
	Id_district string  `json:"district" form:"district" query:"district"`
	Create_date string  `json:"create_date" form:"create_date" query:"create_date"`
	Last_update string  `json:"last_update" form:"last_update" query:"last_update"`
	Lat         float64 `json:"lat" form:"lat" query:"lat"`
	Long        float64 `json:"lng" form:"lng" query:"lng"`
	Id_User     float64 `json:"id_user" form:"id_user" query:"id_user"`
	Is_Open     string  `json:"is_open" form:"is_open" query:"is_open"`
}
