package models

import "time"

type UploadedDocument struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	AgentID      uint      `gorm:"index" json:"agent_id"`
	DocumentType string    `gorm:"size:50;not null" json:"document_type"`
	FilePath     string    `gorm:"type:text;not null" json:"file_path"`
	FileName     string    `gorm:"size:255" json:"file_name"`
	FileSize     int64     `json:"file_size"`
	MimeType     string    `gorm:"size:100" json:"mime_type"`
	UploadedBy   string    `gorm:"size:100" json:"uploaded_by"`
	UploadedAt   time.Time `json:"uploaded_at"`
}
