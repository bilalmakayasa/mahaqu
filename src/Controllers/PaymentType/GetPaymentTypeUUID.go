package Controllers

import (
	"mahaqu/src/config"
	"mahaqu/src/helper"
	"mahaqu/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPaymentTypeUUID(c *gin.Context) {
	param := c.Param("uuid")
	var paymentType models.PaymentType

	if err := config.DB.Where("uuid = ?", param).First(&paymentType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// today := time.Now()
	// var isActive bool
	// estDayLeft := int(targetDonation.ExpiredDate.Sub(today).Hours()) / 24
	// if estDayLeft > 0 {
	// 	isActive = true
	// } else {
	// 	isActive = false
	// }
	// result := models.TargetDonationResponse{
	// 	ID:                 int(targetDonation.ID),
	// 	Name:               targetDonation.Name,
	// 	Description:        targetDonation.Description,
	// 	UUID:               targetDonation.UUID,
	// 	CategoryDonationID: targetDonation.CategoryDonationID,
	// 	ExpiredDate:        targetDonation.ExpiredDate,
	// 	ExpiredDaysLeft:    estDayLeft,
	// 	IsActive:           isActive,
	// 	TargetAmount:       targetDonation.TargetAmount,
	// 	CurrentAmount:      targetDonation.CurrentAmount,
	// 	DonationPercentage: (float64(targetDonation.CurrentAmount) / float64(targetDonation.TargetAmount)) * 100,
	// }
	helper.Response(c, "MHQ0002", "fetch data success", paymentType, nil)
}
