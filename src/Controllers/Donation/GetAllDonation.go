package Controllers

import (
	"mahaqu/src/config"
	"mahaqu/src/helper"
	"mahaqu/src/models"
	"net/http"

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
	helper.Response(c, "MHQ0002", "fetch data success", donation, nil)
}
