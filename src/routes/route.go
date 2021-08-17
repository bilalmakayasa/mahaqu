package routes

import (
	Controllers "mahaqu/src/Controllers/TargetDonation"

	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine {
	r := gin.Default()

	targetDonation := r.Group("/target_donation")
	targetDonation.POST("", Controllers.CreateTargetDonation)
	targetDonation.GET("", Controllers.GetAllTargetDonation)
	targetDonation.PUT("/:uuid", Controllers.UpdateTargetDonation)
	targetDonation.DELETE("/:uuid", Controllers.DeleteTargetDonation)
	return r
}
