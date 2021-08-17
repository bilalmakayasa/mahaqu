package models

type DonationStatus struct {
	BaseModel
	Name     string     `json:"name"`
	Donation []Donation `gorm:"Foreignkey:status_id;association_foreignkey:ID;" json:"donation"`
}
