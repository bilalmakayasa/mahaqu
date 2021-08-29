package Controllers

import (
	"mahaqu/src/config"
	"mahaqu/src/helper"
	"mahaqu/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfirmDonation(c *gin.Context) {
	param := c.Param("uuid")

	var donation models.Donation

	if err := config.DB.Where("uuid = ?", param).First(&donation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	donation.StatusID = 2
	if err := config.DB.Save(&donation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	helper.Response(c, "MHQ0001", "updated data", donation, nil)
}