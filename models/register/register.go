package register

import "time"

type Register struct {
	ID          uint
	Id_Location int
	Tel         string
	Image       string
	NameShop    string
	Address     string
	Id_address  int
	Id_amper    int
	Id_district int
	Create_date time.Time
	Last_Update time.Time
}
