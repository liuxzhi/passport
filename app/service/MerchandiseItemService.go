package service

import (
	"passport/app/request"
	"passport/database"
	"passport/models"
)

type MerchandiseItemService struct{}

func (M *MerchandiseItemService) GetMerchandiseItemList(param request.MerchandiseGetRequest) []models.MerchandiseItem {
	MerchandiseItems := make([]models.MerchandiseItem, 0)
	err := database.MasterDB.Where("merchandise_id = ? ", param.Id).Find(&MerchandiseItems)
	if err != nil {

	}
	return MerchandiseItems
}
