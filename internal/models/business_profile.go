package models

import (
	"time"

	"gorm.io/gorm"
)

type BusinessProfile struct {
	ID                    uint      `gorm:"primaryKey" json:"id"`
	AgentID               uint      `gorm:"index" json:"agent_id"`
	BusinessName          string    `gorm:"size:255;not null" json:"business_name"`
	ContactPerson         string    `gorm:"size:255" json:"contact_person"`
	ContactPhone          string    `gorm:"size:50" json:"contact_phone"`
	Phone                 string    `gorm:"size:50" json:"phone"`
	BusinessNPWP          string    `gorm:"size:100" json:"business_npwp"`
	BusinessEmail         string    `gorm:"size:255" json:"business_email"`
	BusinessAddress       string    `gorm:"type:text" json:"business_address"`
	Latitude              float64   `json:"latitude"`
	Longitude             float64   `json:"longitude"`
	City                  string    `gorm:"size:100" json:"city"`
	Province              string    `gorm:"size:100" json:"province"`
	PostalCode            string    `gorm:"size:20" json:"postal_code"`
	ProductDescription    string    `gorm:"size:255" json:"product_description"`
	BusinessDuration      string    `gorm:"size:100" json:"business_duration"`
	MonthlyGrossProfit    string    `gorm:"size:100" json:"monthly_gross_profit"`
	MonthlyTransactionAvg string    `gorm:"size:100" json:"monthly_transaction_avg"`
	OperatingDays         string    `gorm:"size:100" json:"operating_days"`
	OperatingHours        string    `gorm:"size:100" json:"operating_hours"`
	Notes                 string    `gorm:"type:text" json:"notes"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

func (b *BusinessProfile) BeforeCreate(tx *gorm.DB) (err error) {
	// placeholder hook
	return
}
