package Controllers

import (
	"fmt"
	"mahaqu/src/helper"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DeleteTargetDonationPicture(c *gin.Context) {
	fileName := c.Param("image_url")
	fileBase := filepath.Join("statics", "images", fileName)
	err := os.Remove(fileBase)

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	helper.Response(c, "MHQ0004", "delete data", `success`, nil)
}
