package models

import (
	"time"
)

type TargetDonation struct {
	BaseModel
	CategoryDonationID int        `json:"category_donation_id"`
	Name               string     `json:"name"`
	Description        string     `json:"description" gorm:"type:text"`
	TargetAmount       float64    `json:"target_amount"`
	ExpiredDate        time.Time  `json:"expired_date"`
	ImageUrl           string     `json:"image_url"`
	CurrentAmount      int        `json:"current_amount"`
	Donation           []Donation `gorm:"Foreignkey:target_donation_id;association_foreignkey:ID;" json:"donation"`
}

type TargetDonationInput struct {
	Name               string  `json:"name"`
	Description        string  `json:"description" gorm:"type:text"`
	TargetAmount       float64 `json:"target_amount"`
	ExpiredDate        string  `json:"expired_date"`
	ImageUrl           string  `json:"image_url"`
	CategoryDonationID int     `json:"category_donation_id"`
}
