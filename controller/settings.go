package controller

import (
	"github.com/AH-dark/random-donate/database"
	"github.com/AH-dark/random-donate/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBasicSettings(c *gin.Context) {
	settings := database.GetSettingByType([]string{"basic"})

	c.JSON(http.StatusOK, &model.ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    settings,
	})
}
