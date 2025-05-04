package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitiatePhoneCall(ctx *gin.Context) {
	// Extract the phone number from the request body
	var requestBody struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(400, map[string]interface{}{
			"status": 400,
			"message": "Invalid request body",
			"data": nil,
		})
		return
	}

	fmt.Printf("Initiating phone call to %s\n", requestBody.PhoneNumber)
	// Here you would typically call a service to initiate the phone call

	ctx.JSON(200, map[string]interface{}{
		"status": 200,
		"message": "Phone call initiated successfully",
		"data": nil,
	})
}