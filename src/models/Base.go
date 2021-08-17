package models

import "time"

type BaseModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	UUID      string     `json:"uuid"`
}
