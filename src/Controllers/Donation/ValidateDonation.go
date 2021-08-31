package Controllers

import (
	"mahaqu/src/config"
	"mahaqu/src/helper"
	"mahaqu/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateDonation(c *gin.Context) {
	param := c.Param("uuid")

	var donation models.Donation

	if err := config.DB.Where("uuid = ?", param).First(&donation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var targetDonation models.TargetDonation

	if err := config.DB.Where("id = ?", donation.TargetDonationID).First(&targetDonation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if donation.StatusID == 3 {
		helper.Response(c, "MHQ0001", "the payment has been validate", nil, nil)
	}
	donation.StatusID = 3
	if err := config.DB.Save(&donation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	targetDonation.DonationCount += 1
	if err := config.DB.Save(&targetDonation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	helper.Response(c, "MHQ0001", "updated data", donation, nil)
}
