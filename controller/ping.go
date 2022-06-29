package controller

import (
	"github.com/AH-dark/random-donate/dataType"
	"github.com/gin-gonic/gin"
)

func PingHandler(context *gin.Context) {
	context.JSON(200, &dataType.ApiResponse{
		Code:    200,
		Message: "success",
	})
}
