package request

type MerchandiseGetRequest struct {
	Id string `form:"id" json:"id" xml:"id"  binding:"required"`
}
