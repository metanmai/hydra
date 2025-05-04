package routes

import (
	"github.com/gin-gonic/gin"
	"hydra/controller" 
	"hydra/logger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	api.Use(logger.RequestLogger())

	api.POST("/initiate-phone-call", controller.InitiatePhoneCall)

	return r
}
