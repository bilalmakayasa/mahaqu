package routes

import (
	Donation "mahaqu/src/Controllers/Donation"
	PaymentType "mahaqu/src/Controllers/PaymentType"
	TargetDonation "mahaqu/src/Controllers/TargetDonation"

	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine {
	r := gin.Default()

	targetDonation := r.Group("/target_donation")
	targetDonation.POST("", TargetDonation.CreateTargetDonation)
	targetDonation.GET("", TargetDonation.GetAllTargetDonation)
	targetDonation.GET("/:uuid", TargetDonation.GetTargetDonationUUID)
	targetDonation.PUT("/:uuid", TargetDonation.UpdateTargetDonation)
	targetDonation.DELETE("/:uuid", TargetDonation.DeleteTargetDonation)

	targetDonation.POST("/image", TargetDonation.UploadTargetDonationPicture)
	targetDonation.GET("/image/:image_url", TargetDonation.GetTargetDonationPicture)
	targetDonation.DELETE("/image/:image_url", TargetDonation.DeleteTargetDonation)

	paymentType := r.Group("/payment_type")
	paymentType.GET("", PaymentType.GetAllPaymentType)
	paymentType.GET("/:uuid", PaymentType.GetPaymentTypeUUID)

	donation := r.Group("/donation")
	donation.POST("/request", Donation.RequestDonation)
	donation.GET("/confirm/:uuid", Donation.ConfirmDonation)
	donation.GET("/validate/:uuid", Donation.ValidateDonation)
	donation.GET("", Donation.GetAllDonation)
	return r
}
