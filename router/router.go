package router

import (
	"github.com/gin-gonic/gin"
	"passport/app/handler"
)

func RegisterRoutes(s *gin.Engine) {
	new(handler.MerchandiseHandler).RegisterRoute(s)
}