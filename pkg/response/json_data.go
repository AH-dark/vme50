package response

import (
	"github.com/AH-dark/random-donate/dataType"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DataHandle(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &dataType.ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
	c.Done()
}
