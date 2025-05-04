package controller

import (
	"fmt"
	"hydra/response"

	"github.com/gin-gonic/gin"
)

func InitiatePhoneCall(ctx *gin.Context) {
	var requestBody struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		response.GenericApiResponse(ctx, "Invalid request body", 400, nil)
		return
	}

	fmt.Printf("Initiating phone call to %s\n", requestBody.PhoneNumber)

	response.GenericApiResponse(ctx, "Phone call initiated successfully", 200, nil)
}