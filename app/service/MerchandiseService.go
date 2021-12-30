package service

import (
	"fmt"
	"passport/database"
	"passport/models"
)



type MerchandiseService struct{}

func (M *MerchandiseService) Get(params map[string]interface{}) (category *models.Category) {
	fmt.Println("params", params)
	category = &models.Category{}
	database.MasterDB.ID(1).Get(category)
	return category
}
