package models

import "time"

type BankInfo struct {
	ID                 uint       `gorm:"primaryKey" json:"id"`
	AgentID            uint       `gorm:"index" json:"agent_id"`
	BankAccountNumber  string     `gorm:"size:100;not null" json:"bank_account_number"`
	BankName           string     `gorm:"size:255;not null" json:"bank_name"`
	BankCode           string     `gorm:"size:50" json:"bank_code"`
	AccountHolderName  string     `gorm:"size:255;not null" json:"account_holder_name"`
	VerificationStatus string     `gorm:"size:30;default:'unverified'" json:"verification_status"`
	CreatedAt          time.Time  `json:"created_at"`
	VerifiedAt         *time.Time `json:"verified_at"`
}
