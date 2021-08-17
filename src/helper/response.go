package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, code string, message string, data interface{}, meta interface{}) {
	if data == nil && meta == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    code,
			"message": message,
		})
	} else if meta == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": message,
			"data":    data,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": message,
			"meta":    meta,
			"data":    data,
		})
	}
}
