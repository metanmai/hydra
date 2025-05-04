package response

import "github.com/gin-gonic/gin"

func GenericApiResponse(ctx *gin.Context, message string, statusCode int, data map[string]any) {
	ctx.JSON(statusCode, map[string]any{
		"status":  statusCode,
		"message": message,
		"data":    data,
	})
}