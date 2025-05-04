package routes

import (
	"github.com/gin-gonic/gin"
	"hydra/controller" 
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	api.POST("/initiate-phone-call", controller.InitiatePhoneCall)

	return r
}
