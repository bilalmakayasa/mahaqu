package Controllers

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetTargetDonationPicture(c *gin.Context) {
	fileName := c.Param("image_url")
	fileBase := filepath.Join("storage", "images", fileName)
	c.File(fileBase)
}
