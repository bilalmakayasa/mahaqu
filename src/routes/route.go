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
	targetDonation.GET("/:uuid", Controllers.GetTargetDonationUUID)
	targetDonation.PUT("/:uuid", Controllers.UpdateTargetDonation)
	targetDonation.DELETE("/:uuid", Controllers.DeleteTargetDonation)

	targetDonation.POST("/image", Controllers.UploadTargetDonationPicture)
	targetDonation.GET("/image/:image_url", Controllers.GetTargetDonationPicture)
	targetDonation.DELETE("/image/:image_url", Controllers.DeleteTargetDonationPicture)
	return r
}
