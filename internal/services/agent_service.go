package services

import (
	"agen_edc/internal/models"
	"agen_edc/internal/repositories"
	"encoding/json"

	"gorm.io/datatypes"
)

type AgentService struct {
	agentRepo    *repositories.AgentRepo
	auditLogRepo *repositories.AuditLogRepo
}

func NewAgentService(agentRepo *repositories.AgentRepo, auditLogRepo *repositories.AuditLogRepo) *AgentService {
	return &AgentService{
		agentRepo:    agentRepo,
		auditLogRepo: auditLogRepo,
	}
}

func (s *AgentService) Create(agent *models.Agent, actor string) error {
	err := s.agentRepo.Create(agent)
	if err != nil {
		return err
	}
	// Log audit
	metadata := map[string]interface{}{
		"agent_id": agent.ID,
		"action":   "create",
	}
	metadataJSON, _ := json.Marshal(metadata)
	auditLog := &models.AuditLog{
		AgentID:  agent.ID,
		Action:   "CREATE_AGENT",
		Actor:    actor,
		Metadata: datatypes.JSON(metadataJSON),
	}
	s.auditLogRepo.Create(auditLog)
	return nil
}

func (s *AgentService) GetByID(id uint) (*models.Agent, error) {
	return s.agentRepo.GetByID(id)
}

func (s *AgentService) Update(agent *models.Agent, actor string) error {
	err := s.agentRepo.Update(agent)
	if err != nil {
		return err
	}
	// Log audit
	metadata := map[string]interface{}{
		"agent_id": agent.ID,
		"action":   "update",
	}
	metadataJSON, _ := json.Marshal(metadata)
	auditLog := &models.AuditLog{
		AgentID:  agent.ID,
		Action:   "UPDATE_AGENT",
		Actor:    actor,
		Metadata: datatypes.JSON(metadataJSON),
	}
	s.auditLogRepo.Create(auditLog)
	return nil
}

func (s *AgentService) Delete(id uint, actor string) error {
	agent, err := s.agentRepo.GetByID(id)
	if err != nil {
		return err
	}
	err = s.agentRepo.Delete(id)
	if err != nil {
		return err
	}
	// Log audit
	metadata := map[string]interface{}{
		"agent_id": id,
		"action":   "delete",
	}
	metadataJSON, _ := json.Marshal(metadata)
	auditLog := &models.AuditLog{
		AgentID:  agent.ID,
		Action:   "DELETE_AGENT",
		Actor:    actor,
		Metadata: datatypes.JSON(metadataJSON),
	}
	s.auditLogRepo.Create(auditLog)
	return nil
}

func (s *AgentService) GetAll() ([]models.Agent, error) {
	return s.agentRepo.GetAll()
}

func (s *AgentService) GetFullByID(id uint) (*models.AgentFull, error) {
	return s.agentRepo.GetFullByID(id)
}

func (s *AgentService) Search(filters map[string]interface{}) ([]models.Agent, error) {
	return s.agentRepo.Search(filters)
}
