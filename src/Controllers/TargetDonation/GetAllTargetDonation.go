package Controllers

import (
	"mahaqu/src/config"
	"mahaqu/src/helper"
	"mahaqu/src/models"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllTargetDonation(c *gin.Context) {
	var param models.Params
	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var targetDonation []models.TargetDonation

	baseQuerry := config.DB
	queryResult := baseQuerry.Limit(param.Limit).Offset(param.Offset).Find(&targetDonation)

	if queryResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var count int64

	if err := config.DB.Model(&targetDonation).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var result []models.TargetDonationAllResponse
	today := time.Now()
	var isActive bool
	for _, item := range targetDonation {
		estDayLeft := int(item.ExpiredDate.Sub(today).Hours()) / 24
		if estDayLeft > 0 {
			isActive = true
		} else {
			isActive = false
		}
		data := models.TargetDonationAllResponse{
			ID:                 int(item.ID),
			Name:               item.Name,
			UUID:               item.UUID,
			CategoryDonationID: item.CategoryDonationID,
			ExpiredDate:        item.ExpiredDate,
			ExpiredDaysLeft:    estDayLeft,
			IsActive:           isActive,
			Description:        item.Description,
			TargetAmount:       item.TargetAmount,
			DonationCount:      item.DonationCount,
			CurrentAmount:      item.CurrentAmount,
			ImageUrl:           item.ImageUrl,
			DonationPercentage: (float64(item.CurrentAmount) / float64(item.TargetAmount)) * 100,
		}
		result = append(result, data)
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].ExpiredDaysLeft < result[j].ExpiredDaysLeft
	})

	resultData := models.GetAllResponse{
		Count: count,
		Data:  result,
	}
	helper.Response(c, "MHQ0002", "fetch data success", resultData, nil)
}
