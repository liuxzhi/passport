package service

import (
	"fmt"
	"passport/database"
	"passport/models"
	"xorm.io/xorm"
)

type AttributeService struct{}

func (A *AttributeService) GetAttributeList(conditions map[string]string) []models.Attribute {
	attributeList := make([]models.Attribute, 0)
	ids, ok := conditions["ids"]
	var session *xorm.Session
	if ok {
		session = database.MasterDB.Where("id IN ("+ ids+")")
	}

	if session != nil {
		err := session.Find(&attributeList)
		if err != nil {
			fmt.Println("attribute_list_err", err)
		}
	}


	return attributeList
}
