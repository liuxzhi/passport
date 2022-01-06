package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"passport/app/request"
	"passport/app/service"
	"passport/util"
	"strconv"
	"strings"
)

type MerchandiseHandler struct{}

func (m *MerchandiseHandler) RegisterRoute(s *gin.Engine) {
	s.POST("/merchandise/get", m.Get)
}

func (m *MerchandiseHandler) Get(c *gin.Context) {
	var reqInfo request.MerchandiseGetRequest
	if err := c.BindJSON(&reqInfo); err != nil {
		c.JSON(200, gin.H{
			"error": "请求参数类型错误",
		})
		return
	}

	merchandise := new(service.MerchandiseService).Get(reqInfo)
	merchandiseItems := new(service.MerchandiseItemService).GetMerchandiseItemList(reqInfo)
	merchandiseAttributeList := new(service.MerchandiseAttributeService).GetMerchandiseAttributeList(reqInfo)

	merchandiseAttributeIdTemp := make([]string, 0)
	for _, merchandiseAttribute := range merchandiseAttributeList {
		merchandiseAttributeIdTemp = append(merchandiseAttributeIdTemp, strconv.Itoa(merchandiseAttribute.AttributeId))
	}

	merchandiseAttributeIds := strings.Join(merchandiseAttributeIdTemp, ",")
	conditions4AttributeList := make(map[string]string)
	conditions4AttributeList["ids"] = merchandiseAttributeIds
	attributeList := new(service.AttributeService).GetAttributeList(conditions4AttributeList)

	merchandiseAttributeMapList := make([]map[string]interface{}, 0)

	for _, merchandiseAttribute := range merchandiseAttributeList {
		for _, attribute := range attributeList {
			if merchandiseAttribute.AttributeId == attribute.Id {
				var merchandiseAttributeMap = make(map[string]interface{})
				merchandiseAttributeMap["id"] = merchandiseAttribute.Id
				merchandiseAttributeMap["merchandise_id"] = merchandiseAttribute.MerchandiseId
				merchandiseAttributeMap["attribute_id"] = merchandiseAttribute.AttributeId
				merchandiseAttributeMap["is_prime"] = merchandiseAttribute.IsPrime
				merchandiseAttributeMap["name"] = attribute.Name
				merchandiseAttributeMapList = append(merchandiseAttributeMapList, merchandiseAttributeMap)
			}
		}
	}

	conditions4MerchandiseAttributeValueList := make(map[string]string)
	conditions4MerchandiseAttributeValueList["merchandise_id"] = reqInfo.Id
	merchandiseAttributeValueList := new(service.MerchandiseAttributeValueService).GetMerchandiseAttributeValueList(conditions4MerchandiseAttributeValueList)

	merchandiseAttributeValueIds := make([]string, 0)
	for _, merchandiseAttributeValue := range merchandiseAttributeValueList {
		merchandiseAttributeValueIds = append(merchandiseAttributeValueIds, strconv.Itoa(merchandiseAttributeValue.AttributeValueId))
	}

	conditions4AttributeValueList := make(map[string]string)
	attributeIds := strings.Join(merchandiseAttributeValueIds, ",")
	conditions4AttributeValueList["ids"] = attributeIds
	attributeValueList := new(service.AttributeValueService).GetAttributeValueList(conditions4AttributeValueList)

	merchandiseAttributeValueMapList := make([]map[string]interface{}, 0)

	for _, merchandiseAttributeValue := range merchandiseAttributeValueList {
		for _, attributeValue := range attributeValueList {
			if merchandiseAttributeValue.AttributeValueId == attributeValue.Id {
				var merchandiseAttributeValueMap = make(map[string]interface{}, 0)
				merchandiseAttributeValueMap["id"] = merchandiseAttributeValue.Id
				merchandiseAttributeValueMap["attribute_id"] = merchandiseAttributeValue.AttributeId
				merchandiseAttributeValueMap["attribute_value_id"] = merchandiseAttributeValue.AttributeValueId
				merchandiseAttributeValueMap["merchandise_id"] = merchandiseAttributeValue.MerchandiseId
				merchandiseAttributeValueMap["value"] = attributeValue.Value
				merchandiseAttributeValueMapList = append(merchandiseAttributeValueMapList, merchandiseAttributeValueMap)
			}
		}
	}

	for index, merchandiseAttributeMapItem := range merchandiseAttributeMapList {
		var temp = make([]interface{}, 0)
		for _, merchandiseAttributeValueMapItem := range merchandiseAttributeValueMapList {
			if merchandiseAttributeMapItem["attribute_id"] == merchandiseAttributeValueMapItem["attribute_id"] {
				temp = append(temp, merchandiseAttributeValueMapItem["attribute_value_id"])
			}
		}
		merchandiseAttributeMapList[index]["value_ids"] = temp
	}

	merchandiseAttributeValueAssociatedList := make([]map[string]interface{}, 0)
	for _, merchandiseAttributeValueMap := range merchandiseAttributeValueMapList {

		var merchandiseAttributeValueAssociatedData = make(map[string]interface{})
		var name = ""

		for _, merchandiseAttributeMapItem := range merchandiseAttributeMapList {
			if merchandiseAttributeValueMap["attribute_id"] == merchandiseAttributeMapItem["attribute_id"] {
				value, ok := merchandiseAttributeMapItem["name"].(string)
				if ok {
					name = value
				}
			}
		}
		merchandiseAttributeValueAssociatedData["name"] = name

		temp := make([]interface{}, 0)
		for _, value := range merchandiseAttributeValueMap {
			for _, attributeValue := range attributeValueList {
				if attributeValue.Id == value {
					temp = append(temp, attributeValue)
				}
			}
		}
		merchandiseAttributeValueAssociatedData["values"] = temp
		merchandiseAttributeValueAssociatedList = append(merchandiseAttributeValueAssociatedList, merchandiseAttributeValueAssociatedData)
	}
	fmt.Println("merchandiseAttributeValueAssociatedList", merchandiseAttributeValueAssociatedList)
	merchandiseInfo := util.Struct2Map(*merchandise)
	merchandiseInfo["item_id"] = "0"
	if reqInfo.ItemId != "" {
		merchandiseInfo["item_id"] = reqInfo.ItemId
	}

	merchandiseInfo["attribute_list"] = merchandiseAttributeMapList
	merchandiseInfo["attribute_value_list"] = merchandiseAttributeValueList
	merchandiseInfo["attribute_value_associated_list"] = merchandiseAttributeValueAssociatedList
	merchandiseInfo["item_list"] = merchandiseItems

	c.JSON(200, gin.H{
		"merchandise": merchandiseInfo,
	})
}
