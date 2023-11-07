package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccessmessage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "test successful",
	})
}

func CalculateSum(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "test successful",
		"sum":     4,
	})
}
