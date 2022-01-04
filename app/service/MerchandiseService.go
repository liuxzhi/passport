package service

import (
	"fmt"
	"passport/app/request"
	"passport/database"
	"passport/models"
)

type MerchandiseService struct{}

func (M *MerchandiseService) Get(param request.MerchandiseGetRequest) *models.Merchandise {
	merchandise := new(models.Merchandise)
	_, err := database.MasterDB.ID(param.Id).Get(merchandise)
	if err != nil {
		fmt.Println("MerchandiseService Get err", err)
	}
	return merchandise
}
