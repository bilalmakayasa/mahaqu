package models

type CategoryDonation struct {
	BaseModel
	Name           string           `json:"name"`
	TargetDonation []TargetDonation `gorm:"Foreignkey:category_donation_id;association_foreignkey:ID;" json:"target_donation"`
}
