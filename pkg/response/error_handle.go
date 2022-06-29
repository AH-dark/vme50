package response

import (
	"github.com/AH-dark/random-donate/dataType"
	"github.com/AH-dark/random-donate/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandle(c *gin.Context, err error, code int) {
	utils.Log().Error(err.Error())
	c.JSON(code, &dataType.ApiResponse{
		Code:    code,
		Message: http.StatusText(code),
	})
	return
}

func ServerErrorHandle(c *gin.Context, err error) {
	utils.Log().Error(err.Error())
	c.JSON(http.StatusInternalServerError, &dataType.ApiResponse{
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
	})
	return
}
