package Controllers

import (
	"mahaqu/src/config"
	"mahaqu/src/helper"
	"mahaqu/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTargetDonation(c *gin.Context) {
	param := c.Param("uuid")

	var targetDonation models.TargetDonation

	if err := config.DB.Where("uuid = ?", param).First(&targetDonation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := config.DB.Where("uuid = ?", param).Delete(&targetDonation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	helper.Response(c, "MHQ0001", "deleted data", targetDonation, nil)
}
