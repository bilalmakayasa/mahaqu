package models

type Params struct {
	Limit            int `form:"limit"`
	Offset           int `form:"offset"`
	TargetDonationID int `form:"target_donation_id"`
	StatusID         int `form:"status_id"`
}
