package Controllers

import (
	"mahaqu/src/config"
	"mahaqu/src/helper"
	"mahaqu/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateTargetDonation(c *gin.Context) {
	param := c.Param("uuid")

	var targetDonation models.TargetDonation
	var targetDonationInput models.TargetDonationInput

	if err := config.DB.Where("uuid = ?", param).First(&targetDonation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&targetDonationInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	targetDonation.Name = targetDonationInput.Name
	targetDonation.TargetAmount = targetDonationInput.TargetAmount
	targetDonation.Description = targetDonationInput.Description
	targetDonation.CategoryDonationID = targetDonationInput.CategoryDonationID
	targetDonation.ImageUrl = targetDonationInput.ImageUrl
	if err := config.DB.Save(&targetDonation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	helper.Response(c, "MHQ0001", "updated data", targetDonation, nil)
}
