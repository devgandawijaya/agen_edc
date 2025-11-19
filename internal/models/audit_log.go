package models

import (
	"time"

	"gorm.io/datatypes"
)

type AuditLog struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	AgentID   uint           `json:"agent_id"`
	Action    string         `gorm:"size:255;not null" json:"action"`
	Actor     string         `gorm:"size:255" json:"actor"`
	Metadata  datatypes.JSON `json:"metadata"`
	CreatedAt time.Time      `json:"created_at"`
}
