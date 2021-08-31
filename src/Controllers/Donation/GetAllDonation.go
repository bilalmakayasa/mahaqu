package Controllers

import (
	"mahaqu/src/config"
	"mahaqu/src/helper"
	"mahaqu/src/models"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllDonation(c *gin.Context) {
	var param models.Params
	var donation []models.Donation
	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	baseQuerry := config.DB
	if err := baseQuerry.Limit(param.Limit).Offset(param.Offset).Find(&donation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var result []models.DonationResponse
	now := time.Now()
	for _, item := range donation {

		data := models.DonationResponse{
			ID:               int(item.ID),
			TargetDonationID: item.TargetDonationID,
			PaymentTypeID:    item.PaymentTypeID,
			StatusID:         item.StatusID,
			Name:             item.Name,
			Amount:           item.Amount,
			Email:            item.Email,
			IsVisible:        item.IsVisible,
			Message:          item.Message,
			UUID:             item.UUID,
		}
		test := now.Sub(*item.CreatedAt)

		if test.Hours() > 24 {
			data.TimeValue = int(math.Floor(test.Hours() / 24))
			data.TimeType = "days"
		} else if test.Hours() >= 1 {
			data.TimeValue = int(test.Hours())
			data.TimeType = "hours"
		} else {
			data.TimeValue = int(test.Minutes())
			data.TimeType = "minutes"
		}

		result = append(result, data)
	}
	helper.Response(c, "MHQ0002", "fetch data success", result, nil)
}
