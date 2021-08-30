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
	UUID               string     `json:"uuid"`
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

type UrlResponse struct {
	ImageUrl string `json:"image_url"`
}

type TargetDonationAllResponse struct {
	ID                 int       `json:"id"`
	CategoryDonationID int       `json:"category_donation_id"`
	Name               string    `json:"name"`
	TargetAmount       float64   `json:"target_amount"`
	UUID               string    `json:"uuid"`
	ExpiredDate        time.Time `json:"expired_date"`
	CurrentAmount      int       `json:"current_amount"`
	ExpiredDaysLeft    int       `json:"expired_days_left"`
	IsActive           bool      `json:"is_active"`
	ImageUrl           string    `json:"image_url"`
	DonationPercentage float64   `json:"donation_percentage"`
}

type GetAllResponse struct {
	Count int64       `json:"total"`
	Data  interface{} `json:"data"`
}
type TargetDonationResponse struct {
	ID                 int       `json:"id"`
	CategoryDonationID int       `json:"category_donation_id"`
	Name               string    `json:"name"`
	TargetAmount       float64   `json:"target_amount"`
	Description        string    `json:"description" gorm:"type:text"`
	UUID               string    `json:"uuid"`
	ExpiredDate        time.Time `json:"expired_date"`
	CurrentAmount      int       `json:"current_amount"`
	ExpiredDaysLeft    int       `json:"expired_days_left"`
	ImageUrl           string    `json:"image_url"`
	IsActive           bool      `json:"is_active"`
	DonationPercentage float64   `json:"donation_percentage"`
}
