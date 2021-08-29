package Controllers

import (
	"log"
	"mahaqu/src/config"
	"mahaqu/src/helper"
	"mahaqu/src/models"
	"mahaqu/src/utility"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type DonationCount struct {
	Count float64
}

func RequestDonation(c *gin.Context) {
	var donationInput models.DonationInput

	if err := c.ShouldBindJSON(&donationInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var targetDonation models.TargetDonation

	if err := config.DB.Where("uuid = ?", donationInput.TargetDonationUUID).First(&targetDonation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "target donation not found!"})
		return
	}

	donationCount := &DonationCount{}
	err := helper.RedisInitialize().GetRedis("donation_count", donationCount)
	var donationCountData float64
	if err == redis.Nil {
		setValue := &DonationCount{Count: 1}
		err := helper.RedisInitialize().SetRedis("donation_count", setValue, 0)
		if err != nil {
			log.Fatalf("Error: %v", err.Error())
		}
		donationCountData = 1
	} else {
		donationCountData = donationCount.Count + 1
		setValue := &DonationCount{Count: donationCountData}
		err := helper.RedisInitialize().SetRedis("donation_count", setValue, 0)
		if err != nil {
			log.Fatalf("Error: %v", err.Error())
		}
	}

	data := models.Donation{
		TargetDonationID: int(targetDonation.ID),
		PaymentTypeID:    donationInput.PaymentTypeID,
		StatusID:         1,
		Name:             donationInput.Name,
		Amount:           donationInput.Amount + donationCountData,
		Email:            donationInput.Email,
		IsVisible:        donationInput.IsVisible,
		UUID:             utility.RandString(40),
	}

	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input database error!"})
		return
	}
	helper.Response(c, "MHQ0001", "created data", data, nil)
}
