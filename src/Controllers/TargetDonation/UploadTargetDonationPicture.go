package Controllers

import (
	"fmt"
	"mahaqu/src/helper"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadTargetDonationPicture(c *gin.Context) {
	file, err := c.FormFile("File")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	fileBase := filepath.Join("statics", "images", file.Filename)
	if err := c.SaveUploadedFile(file, fileBase); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	helper.Response(c, "MHQ0001", "created data", `success`, nil)
}
