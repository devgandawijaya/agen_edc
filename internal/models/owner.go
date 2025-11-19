package models

import "time"

type Owner struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	AgentID        uint       `gorm:"index" json:"agent_id"`
	FullName       string     `gorm:"size:255;not null" json:"full_name"`
	BirthPlace     string     `gorm:"size:255" json:"birth_place"`
	BirthDate      *time.Time `json:"birth_date"`
	Gender         string     `gorm:"size:20" json:"gender"`
	Religion       string     `gorm:"size:50" json:"religion"`
	Occupation     string     `gorm:"size:100" json:"occupation"`
	KTPAddress     string     `gorm:"type:text" json:"ktp_address"`
	City           string     `gorm:"size:100" json:"city"`
	Province       string     `gorm:"size:100" json:"province"`
	PostalCode     string     `gorm:"size:20" json:"postal_code"`
	Phone          string     `gorm:"size:50" json:"phone"`
	Email          string     `gorm:"size:255" json:"email"`
	IdentityNumber string     `gorm:"size:100" json:"identity_number"`
	TaxNumber      string     `gorm:"size:100" json:"tax_number"`
	CreatedAt      time.Time  `json:"created_at"`
}
