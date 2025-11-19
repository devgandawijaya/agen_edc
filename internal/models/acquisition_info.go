package models

import "time"

type AcquisitionInfo struct {
	ID                     uint      `gorm:"primaryKey" json:"id"`
	AgentID                uint      `gorm:"index" json:"agent_id"`
	AcquisitionType        string    `gorm:"size:30;not null" json:"acquisition_type"`
	AcquisitionName        string    `gorm:"size:255" json:"acquisition_name"`
	AcquisitionNIK         string    `gorm:"size:100" json:"acquisition_nik"`
	AcquisitionCityOrEmail string    `gorm:"size:255" json:"acquisition_city_or_email"`
	CreatedAt              time.Time `json:"created_at"`
}
