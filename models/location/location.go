package location

type Location struct {
	Id   uint    `sql:"primary_key" json:"id" form:"id" query:"id"`
	Lat  float64 `json:"lat" form:"lat" query:"lat"`
	Long float64 `json:"lng" form:"lng" query:"lng"`
}
