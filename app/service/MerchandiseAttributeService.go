package service

import (
	"passport/app/request"
	"passport/database"
	"passport/models"
)

type MerchandiseAttributeService struct{}

func (M *MerchandiseAttributeService) GetMerchandiseAttributeList(param request.MerchandiseGetRequest) []models.MerchandiseAttribute {
	MerchandiseAttributes := make([]models.MerchandiseAttribute, 0)
	err := database.MasterDB.Where("merchandise_id = ? ", param.Id).Find(&MerchandiseAttributes)
	if err != nil {

	}

	return MerchandiseAttributes
}
