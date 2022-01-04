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
	has, err := database.MasterDB.ID(param.Id).Get(merchandise)
	if has {
		fmt.Println("step I", has)
	}

	if err != nil {
		fmt.Println("step II", err)
	}
	return merchandise
}
