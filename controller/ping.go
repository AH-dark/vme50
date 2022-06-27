package controller

import (
	"github.com/AH-dark/random-donate/model"
	"github.com/gin-gonic/gin"
)

func PingHandler(context *gin.Context) {
	context.JSON(200, &model.ApiResponse{
		Code:    200,
		Message: "success",
	})
}
