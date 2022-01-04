package request

type MerchandiseGetRequest struct {
	Id string `form:"id" json:"id" xml:"id"  binding:"required"`
}

func (m *MerchandiseGetRequest) Handle() {

}

type MerchandiseItemRequest struct {
	Id string `form:"id" json:"id" xml:"id"  binding:"required"`
}