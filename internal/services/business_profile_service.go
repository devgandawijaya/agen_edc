package services

import (
	"agen_edc/internal/models"
	"agen_edc/internal/repositories"
)

type BusinessProfileService struct {
	repo *repositories.BusinessProfileRepo
}

func NewBusinessProfileService(repo *repositories.BusinessProfileRepo) *BusinessProfileService {
	return &BusinessProfileService{repo: repo}
}

func (s *BusinessProfileService) Create(profile *models.BusinessProfile) error {
	return s.repo.Create(profile)
}

func (s *BusinessProfileService) GetByID(id uint) (*models.BusinessProfile, error) {
	return s.repo.GetByID(id)
}

func (s *BusinessProfileService) GetByAgentID(agentID uint) ([]models.BusinessProfile, error) {
	return s.repo.GetByAgentID(agentID)
}

func (s *BusinessProfileService) Update(profile *models.BusinessProfile) error {
	return s.repo.Update(profile)
}

func (s *BusinessProfileService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *BusinessProfileService) GetAll() ([]models.BusinessProfile, error) {
	return s.repo.GetAll()
}
