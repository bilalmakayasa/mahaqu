package models

type PaymentType struct {
	BaseModel
	Name          string     `json:"name"`
	AccountNumber string     `json:"account_number"`
	Icon          string     `json:"icon"`
	Donation      []Donation `gorm:"Foreignkey:payment_type_id;association_foreignkey:ID;" json:"donation"`
}
