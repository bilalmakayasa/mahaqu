package Controllers

import (
	"fmt"
	"mahaqu/src/helper"
	"mahaqu/src/models"
	"mahaqu/src/utility"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadTargetDonationPicture(c *gin.Context) {
	file, err := c.FormFile("File")
	fmt.Println("masuk siniiiii")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	fileBase := filepath.Join("storage", "images", file.Filename)
	if err := c.SaveUploadedFile(file, fileBase); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	url := fmt.Sprintf("%s/target_donation/image/%s", utility.GoDotEnvVariable("BASE_URL"), file.Filename)
	data := models.UrlResponse{
		ImageUrl: url,
	}
	helper.Response(c, "MHQ0001", "created data", data, nil)
}
