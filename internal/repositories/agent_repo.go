package repositories

import (
	"agen_edc/internal/models"

	"gorm.io/gorm"
)

type AgentRepo struct{ db *gorm.DB }

func NewAgentRepo(db *gorm.DB) *AgentRepo { return &AgentRepo{db: db} }

func (r *AgentRepo) Create(agent *models.Agent) error { return r.db.Create(agent).Error }
func (r *AgentRepo) GetByID(id uint) (*models.Agent, error) {
	var a models.Agent
	if err := r.db.First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AgentRepo) Update(agent *models.Agent) error {
	return r.db.Save(agent).Error
}

func (r *AgentRepo) Delete(id uint) error {
	return r.db.Delete(&models.Agent{}, id).Error
}

func (r *AgentRepo) GetAll() ([]models.Agent, error) {
	var agents []models.Agent
	err := r.db.Find(&agents).Error
	return agents, err
}

func (r *AgentRepo) GetFullByID(id uint) (*models.AgentFull, error) {
	var agentFull models.AgentFull
	err := r.db.Table("vw_agent_full").Where("id = ?", id).First(&agentFull).Error
	return &agentFull, err
}

func (r *AgentRepo) Search(filters map[string]interface{}) ([]models.Agent, error) {
	var agents []models.Agent
	query := r.db
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}
	err := query.Find(&agents).Error
	return agents, err
}
