package Controllers

import (
	"mahaqu/src/config"
	"mahaqu/src/helper"
	"mahaqu/src/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTargetDonation(c *gin.Context) {
	var targetDonationInput models.TargetDonationInput

	if err := c.ShouldBindJSON(&targetDonationInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dateArray := strings.Split(targetDonationInput.ExpiredDate, "/")
	inputDate, _ := time.Parse("2006-01-02", dateArray[2]+"-"+dateArray[1]+"-"+dateArray[0])
	data := models.TargetDonation{
		Name:               targetDonationInput.Name,
		Description:        targetDonationInput.Description,
		TargetAmount:       targetDonationInput.TargetAmount,
		ExpiredDate:        inputDate,
		ImageUrl:           targetDonationInput.ImageUrl,
		CategoryDonationID: targetDonationInput.CategoryDonationID,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input database error!"})
		return
	}

	helper.Response(c, "MHQ0001", "created data", data, nil)
}
