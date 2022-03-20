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
