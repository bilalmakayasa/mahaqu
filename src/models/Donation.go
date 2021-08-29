package models

type Donation struct {
	BaseModel
	TargetDonationID int     `json:"target_donation_id"`
	PaymentTypeID    int     `json:"payment_type_id"`
	StatusID         int     `json:"status_id"`
	Name             string  `json:"name"`
	Amount           float64 `json:"amount"`
	Email            string  `json:"email"`
	IsVisible        bool    `json:"is_visible"`
	UUID             string  `json:"uuid"`
}

type DonationInput struct {
	PaymentTypeID      int     `json:"payment_type_id"`
	Name               string  `json:"name"`
	Amount             float64 `json:"amount"`
	Email              string  `json:"email"`
	IsVisible          bool    `json:"is_visible"`
	TargetDonationUUID string  `json:"target_donation_uuid"`
}
