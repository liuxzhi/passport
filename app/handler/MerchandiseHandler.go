package handler

import (
	"github.com/gin-gonic/gin"
	"passport/app/service"
)

type MerchandiseHandler struct{}

func (m *MerchandiseHandler) RegisterRoute(s *gin.Engine) {
	s.POST("/category/get", m.Get)
}

func (m *MerchandiseHandler) Get(c *gin.Context) {
	requestParamMap := make(map[string]interface{})
	c.BindJSON(&requestParamMap)
	category := new(service.MerchandiseService).Get(requestParamMap)
	c.JSON(200, gin.H{
		"id": category,
	})
}



