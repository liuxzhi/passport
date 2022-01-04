package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"passport/app/request"
	"passport/app/service"
)

type MerchandiseHandler struct{}

func (m *MerchandiseHandler) RegisterRoute(s *gin.Engine) {
	s.POST("/merchandise/get", m.Get)
}

func (m *MerchandiseHandler) Get(c *gin.Context) {
	var reqInfo request.MerchandiseGetRequest

	if err := c.ShouldBindJSON(&reqInfo); err != nil {
		c.JSON(200, gin.H{
			"error": "",
		})
	}
	fmt.Println("reqInfo", reqInfo.Id)

	merchandise := new(service.MerchandiseService).Get(reqInfo)
	c.JSON(200, gin.H{
		"merchandise": merchandise,
	})
}
