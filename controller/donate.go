package controller

import (
	"fmt"
	"github.com/AH-dark/random-donate/dataType"
	"github.com/AH-dark/random-donate/dataType/payment"
	"github.com/AH-dark/random-donate/model"
	"github.com/AH-dark/random-donate/pkg/response"
	"github.com/AH-dark/random-donate/pkg/utils"
	"github.com/AH-dark/random-donate/service"
	"github.com/gin-gonic/gin"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"strings"
)

func DonatePostHandler(c *gin.Context) {
	var data dataType.DonateInfoReq
	err := c.ShouldBind(&data)
	if err != nil {
		response.ErrorHandle(c, err, http.StatusInternalServerError)
		return
	}

	url, err := utils.ParseQRCode(data.QRCode)
	if err != nil {
		response.ErrorHandle(c, err, http.StatusInternalServerError)
		return
	}

	// check payment and url
	switch data.Payment {
	case payment.Alipay:
		if !strings.HasPrefix(url, "https://qr.alipay.com/") {
			c.JSON(http.StatusBadRequest, &dataType.ApiResponse{
				Code:    http.StatusBadRequest,
				Message: "qrcode is not legal",
			})
			return
		}
	case payment.Wechat:
		if !strings.HasPrefix(url, "wxp://") {
			c.JSON(http.StatusBadRequest, &dataType.ApiResponse{
				Code:    http.StatusBadRequest,
				Message: "qrcode is not legal",
			})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, &dataType.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "payment is not legal",
		})
		return
	}

	// generate data
	dbData := model.DonateInfo{
		Name:    data.Name,
		Email:   data.Email,
		Payment: data.Payment,
		Url:     url,
	}

	// find if exist
	isExist, err := service.DonateInfoIsExist(&model.DonateInfo{Email: data.Email})
	if err != nil {
		response.ServerErrorHandle(c, err)
		return
	} else if !isExist {
		c.JSON(http.StatusBadRequest, &dataType.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "data not exist",
		})
		return
	}

	// insert to database
	err = service.DonateInfoSave(&dbData)
	if err != nil {
		response.ServerErrorHandle(c, err)
		return
	}

	// get full data
	dbData, err = service.DonateInfoFind(&dbData)
	if err != nil {
		response.ServerErrorHandle(c, err)
		return
	}

	c.JSON(http.StatusOK, &dataType.ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    dbData,
	})
}

func DonateRandomGetHandler(c *gin.Context) {
	// 取值处理
	sessPrevId := utils.GetSession(c, "random_donate_prev_id")
	prevId := "0"
	if sessPrevId != nil {
		prevId = fmt.Sprintf("%v", sessPrevId)
	}

	// 获取数据
	data, err := service.DonateInfoRandomGet(prevId)
	if err != nil {
		response.ServerErrorHandle(c, err)
		return
	}

	utils.SetSession(c, map[string]interface{}{
		"random_donate_prev_id": data.ID,
	})

	utils.Log().Debug("random donate info: prev: %v, new: %d", prevId, data.ID)

	c.JSON(http.StatusOK, &dataType.ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}
