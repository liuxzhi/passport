package service

import (
	"passport/database"
	"passport/models"
)

type MerchandiseAttributeValueService struct{}

func (M *MerchandiseAttributeValueService) GetMerchandiseAttributeValueList(conditions map[string]string) []models.MerchandiseAttributeValue {
	MerchandiseAttributeValues := make([]models.MerchandiseAttributeValue, 0)
	err := database.MasterDB.Where("merchandise_id = ? ", conditions["merchandise_id"]).Find(&MerchandiseAttributeValues)
	if err != nil {

	}
	return MerchandiseAttributeValues
}
