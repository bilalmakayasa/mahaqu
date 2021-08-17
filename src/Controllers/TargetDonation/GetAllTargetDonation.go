package Controllers

import (
	"mahaqu/src/config"
	"mahaqu/src/helper"
	"mahaqu/src/models"

	"github.com/gin-gonic/gin"
)

func GetAllTargetDonation(c *gin.Context) {
	var targetDonation []models.TargetDonation

	baseQuerry := config.DB
	baseQuerry.Find(&targetDonation)

	helper.Response(c, "MHQ0002", "fetch data success", targetDonation, nil)
}
