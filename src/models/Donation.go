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
	Message          string  `json:"message" gorm:"type:text"`
	UUID             string  `json:"uuid"`
}
type DonationResponse struct {
	ID               int     `json:"id"`
	TargetDonationID int     `json:"target_donation_id"`
	PaymentTypeID    int     `json:"payment_type_id"`
	StatusID         int     `json:"status_id"`
	Name             string  `json:"name"`
	Amount           float64 `json:"amount"`
	Email            string  `json:"email"`
	IsVisible        bool    `json:"is_visible"`
	Message          string  `json:"message" gorm:"type:text"`
	UUID             string  `json:"uuid"`
	TimeValue        int     `json:"time_value"`
	TimeType         string  `json:"time_type"`
}

type DonationInput struct {
	PaymentTypeID      int     `json:"payment_type_id"`
	Name               string  `json:"name"`
	Amount             float64 `json:"amount"`
	Email              string  `json:"email"`
	IsVisible          bool    `json:"is_visible"`
	Message            string  `json:"message" gorm:"type:text"`
	TargetDonationUUID string  `json:"target_donation_uuid"`
}
