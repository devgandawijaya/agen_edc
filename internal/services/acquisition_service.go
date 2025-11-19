package services

import (
	"agen_edc/internal/models"
	"agen_edc/internal/repositories"
)

type AcquisitionService struct {
	repo *repositories.AcquisitionRepo
}

func NewAcquisitionService(repo *repositories.AcquisitionRepo) *AcquisitionService {
	return &AcquisitionService{repo: repo}
}

func (s *AcquisitionService) Create(acquisition *models.AcquisitionInfo) error {
	return s.repo.Create(acquisition)
}

func (s *AcquisitionService) GetByID(id uint) (*models.AcquisitionInfo, error) {
	return s.repo.GetByID(id)
}

func (s *AcquisitionService) GetByAgentID(agentID uint) (*models.AcquisitionInfo, error) {
	return s.repo.GetByAgentID(agentID)
}

func (s *AcquisitionService) Update(acquisition *models.AcquisitionInfo) error {
	return s.repo.Update(acquisition)
}

func (s *AcquisitionService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *AcquisitionService) GetAll() ([]models.AcquisitionInfo, error) {
	return s.repo.GetAll()
}
