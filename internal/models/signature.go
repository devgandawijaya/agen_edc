package models

import "time"

type Signature struct {
	ID                   uint       `gorm:"primaryKey" json:"id"`
	AgentID              uint       `gorm:"index" json:"agent_id"`
	OwnerSignaturePath   string     `gorm:"type:text" json:"owner_signature_path"`
	CompanySignaturePath string     `gorm:"type:text" json:"company_signature_path"`
	SignDate             *time.Time `json:"sign_date"`
	CreatedAt            time.Time  `json:"created_at"`
}
