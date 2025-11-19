package services

import (
	"agen_edc/internal/models"
	"agen_edc/internal/repositories"
)

type OwnerService struct {
	repo *repositories.OwnerRepo
}

func NewOwnerService(repo *repositories.OwnerRepo) *OwnerService {
	return &OwnerService{repo: repo}
}

func (s *OwnerService) Create(owner *models.Owner) error {
	return s.repo.Create(owner)
}

func (s *OwnerService) GetByID(id uint) (*models.Owner, error) {
	return s.repo.GetByID(id)
}

func (s *OwnerService) GetByAgentID(agentID uint) ([]models.Owner, error) {
	return s.repo.GetByAgentID(agentID)
}

func (s *OwnerService) Update(owner *models.Owner) error {
	return s.repo.Update(owner)
}

func (s *OwnerService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *OwnerService) GetAll() ([]models.Owner, error) {
	return s.repo.GetAll()
}
