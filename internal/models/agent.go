package models

import (
	"time"

	"gorm.io/datatypes"
)

type Agent struct {
	ID                  uint           `gorm:"primaryKey" json:"id"`
	AgentType           string         `gorm:"size:20;not null" json:"agent_type"`
	BusinessPlaceStatus string         `gorm:"size:10;not null" json:"business_place_status"`
	EDCUsageActivity    datatypes.JSON `json:"edc_usage_activity"`
	CooperationType     string         `gorm:"size:20;not null" json:"cooperation_type"`
	TransactionFeatures datatypes.JSON `json:"transaction_features"`
	SubmissionStatus    string         `gorm:"size:30;default:'pending'" json:"submission_status"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
}

type AgentFull struct {
	ID                  uint           `json:"id"`
	AgentType           string         `json:"agent_type"`
	BusinessPlaceStatus string         `json:"business_place_status"`
	EDCUsageActivity    datatypes.JSON `json:"edc_usage_activity"`
	CooperationType     string         `json:"cooperation_type"`
	TransactionFeatures datatypes.JSON `json:"transaction_features"`
	SubmissionStatus    string         `json:"submission_status"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	// Acquisition fields
	AcquisitionID   *uint      `json:"acquisition_id"`
	AcquisitionDate *time.Time `json:"acquisition_date"`
	// Add other fields from related models as needed
	// For brevity, assuming flattened view
}
