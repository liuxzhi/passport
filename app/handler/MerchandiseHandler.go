package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"passport/app/request"
	"passport/app/service"
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
				merchandiseAttributeMap["Id"] = merchandiseAttribute.Id
				merchandiseAttributeMap["MerchandiseId"] = merchandiseAttribute.MerchandiseId
				merchandiseAttributeMap["AttributeId"] = merchandiseAttribute.AttributeId
				merchandiseAttributeMap["IsPrime"] = merchandiseAttribute.IsPrime
				merchandiseAttributeMap["Name"] = attribute.Name
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
				merchandiseAttributeValueMap["Id"] = merchandiseAttributeValue.Id
				merchandiseAttributeValueMap["AttributeId"] = merchandiseAttributeValue.AttributeId
				merchandiseAttributeValueMap["AttributeValueId"] = merchandiseAttributeValue.AttributeValueId
				merchandiseAttributeValueMap["MerchandiseId"] = merchandiseAttributeValue.MerchandiseId
				merchandiseAttributeValueMap["Value"] = attributeValue.Value
				merchandiseAttributeValueMapList = append(merchandiseAttributeValueMapList, merchandiseAttributeValueMap)
			}
		}
	}

	for index, merchandiseAttributeMapItem := range merchandiseAttributeMapList {
		var temp = make([]interface{}, 0)
		for _, merchandiseAttributeValueMapItems := range merchandiseAttributeValueMapList {
			if merchandiseAttributeMapItem["AttributeId"] == merchandiseAttributeValueMapItems["AttributeId"] {
				temp = append(temp, merchandiseAttributeValueMapItems["AttributeValueId"])
			}
		}
		merchandiseAttributeMapList[index]["Value_ids"] = temp
	}

	fmt.Println("merchandiseAttributeMapList", merchandiseAttributeMapList)

	c.JSON(200, gin.H{
		"merchandise":                      merchandise,
		"merchandise_items":                merchandiseItems,
		"merchandise_attribute_list":       merchandiseAttributeList,
		"attribute_list":                   attributeList,
		"merchandise_attribute_map_list":   merchandiseAttributeMapList,
		"merchandise_attribute_value_list": merchandiseAttributeValueList,
		"attribute_value_list":             attributeValueList,
	})
}
