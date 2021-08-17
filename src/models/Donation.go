package models

type Donation struct {
	BaseModel
	TargetDonationID int     `json:"target_donation_id"`
	PaymentTypeID    int     `json:"payment_type_id"`
	StatusID         int     `json:"status_id"`
	Name             string  `json:"name"`
	Amount           float64 `json:"amount"`
	Email            string  `json:"email"`
	IsVisible        string  `json:"is_visible"`
}
