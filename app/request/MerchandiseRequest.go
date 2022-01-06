package request

type MerchandiseGetRequest struct {
	Id string `form:"id" json:"id" xml:"id"  binding:"required"`
	ItemId string `form:"item_id" json:"item_id" xml:"item_id"`
}

func (m *MerchandiseGetRequest) Handle() {

}

type MerchandiseItemRequest struct {
	Id string `form:"id" json:"id" xml:"id"  binding:"required"`
}