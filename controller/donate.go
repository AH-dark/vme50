package controller

import (
	"github.com/AH-dark/random-donate/database"
	"github.com/AH-dark/random-donate/model"
	"github.com/AH-dark/random-donate/pkg/utils"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strings"
)

func DonatePostHandler(context *gin.Context) {
	if context.Request.Method != http.MethodPost {
		context.JSON(http.StatusMethodNotAllowed, &model.ApiResponse{
			Code:    405,
			Message: "method not allowed",
		})
		return
	}

	data := &database.DonateInfo{}
	err := context.BindJSON(&data)
	if err != nil {
		utils.Log().Error(err.Error())
		context.JSON(http.StatusInternalServerError, &model.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "server error",
		})
		return
	}

	// check payment and url
	switch data.Payment {
	case "alipay":
		if !strings.HasPrefix(data.Url, "https://qr.alipay.com/") {
			context.JSON(http.StatusBadRequest, &model.ApiResponse{
				Code:    http.StatusBadRequest,
				Message: "url is not legal",
			})
			return
		}
	case "wechat":
		if !strings.HasPrefix(data.Url, "wxp://") {
			context.JSON(http.StatusBadRequest, &model.ApiResponse{
				Code:    http.StatusBadRequest,
				Message: "url is not legal",
			})
			return
		}
	default:
		context.JSON(http.StatusBadRequest, &model.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "payment is not legal",
		})
		return
	}

	// insert to database
	err = database.DB.Save(data).Error
	if err != nil {
		utils.Log().Error(err.Error())
		context.JSON(http.StatusInternalServerError, &model.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "server error",
		})
		return
	}

	context.JSON(http.StatusOK, &model.ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
	})
}

func DonateRandomGetHandler(c *gin.Context) {
	var count int64 = 0
	database.DB.Model(&database.DonateInfo{}).Count(&count)
	if count < 1 {
		utils.Log().Warning("Donate Info 表中无数据")
		c.JSON(http.StatusInternalServerError, &model.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "data not exist",
		})
		return
	}

	data := database.DonateInfo{}
	err := database.DB.Offset(rand.Intn(int(count - 1))).First(&data).Error
	if err != nil {
		utils.Log().Error(err.Error())
		c.JSON(http.StatusInternalServerError, &model.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "server error",
		})
		return
	}

	c.JSON(http.StatusOK, &model.ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}
