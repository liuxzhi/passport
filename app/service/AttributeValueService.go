

package service

import (
"fmt"
"passport/database"
"passport/models"
"xorm.io/xorm"
)

type AttributeValueService struct{}

func (A *AttributeValueService) GetAttributeValueList(conditions map[string]string) []models.AttributeValue {
	attributeValueList := make([]models.AttributeValue, 0)
	ids, ok := conditions["ids"]
	var session *xorm.Session
	if ok {
		session = database.MasterDB.Where("id IN ("+ ids+")")
	}

	if session != nil {
		err := session.Find(&attributeValueList)
		if err != nil {
			fmt.Println("attribute_list_err", err)
		}
	}


	return attributeValueList
}
