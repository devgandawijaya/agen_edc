package services

import (
	"agen_edc/internal/models"
	"agen_edc/internal/repositories"
)

type AuditLogService struct {
	repo *repositories.AuditLogRepo
}

func NewAuditLogService(repo *repositories.AuditLogRepo) *AuditLogService {
	return &AuditLogService{repo: repo}
}

func (s *AuditLogService) Create(log *models.AuditLog) error {
	return s.repo.Create(log)
}

func (s *AuditLogService) GetByID(id uint) (*models.AuditLog, error) {
	return s.repo.GetByID(id)
}

func (s *AuditLogService) GetByAgentID(agentID uint) ([]models.AuditLog, error) {
	return s.repo.GetByAgentID(agentID)
}

func (s *AuditLogService) Update(log *models.AuditLog) error {
	return s.repo.Update(log)
}

func (s *AuditLogService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *AuditLogService) GetAll() ([]models.AuditLog, error) {
	return s.repo.GetAll()
}
