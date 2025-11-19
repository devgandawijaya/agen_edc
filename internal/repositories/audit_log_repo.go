package repositories

import (
	"agen_edc/internal/models"

	"gorm.io/gorm"
)

type AuditLogRepo struct {
	db *gorm.DB
}

func NewAuditLogRepo(db *gorm.DB) *AuditLogRepo {
	return &AuditLogRepo{db: db}
}

func (r *AuditLogRepo) Create(auditLog *models.AuditLog) error {
	return r.db.Create(auditLog).Error
}

func (r *AuditLogRepo) GetByID(id uint) (*models.AuditLog, error) {
	var auditLog models.AuditLog
	err := r.db.First(&auditLog, id).Error
	return &auditLog, err
}

func (r *AuditLogRepo) GetByAgentID(agentID uint) ([]models.AuditLog, error) {
	var auditLogs []models.AuditLog
	err := r.db.Where("agent_id = ?", agentID).Find(&auditLogs).Error
	return auditLogs, err
}

func (r *AuditLogRepo) Update(auditLog *models.AuditLog) error {
	return r.db.Save(auditLog).Error
}

func (r *AuditLogRepo) Delete(id uint) error {
	return r.db.Delete(&models.AuditLog{}, id).Error
}

func (r *AuditLogRepo) GetAll() ([]models.AuditLog, error) {
	var auditLogs []models.AuditLog
	err := r.db.Find(&auditLogs).Error
	return auditLogs, err
}
