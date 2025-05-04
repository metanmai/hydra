package controller

import (
	"fmt"
	"hydra/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitiatePhoneCall(ctx *gin.Context) {
	var log *logrus.Entry
	var ok bool

	var requestBody struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		response.GenericApiResponse(ctx, "Invalid request body", 400, nil)
		return
	}

	logger, _ := ctx.Get("logger")
    if log, ok = logger.(*logrus.Entry); !ok {
        fmt.Println("Logger not found in context")
		return
    } 

	log.WithFields(logrus.Fields{
		"phone_number": requestBody.PhoneNumber,
	}).Info("Initiating phone call")


	response.GenericApiResponse(ctx, "Phone call initiated successfully", 200, nil)
}