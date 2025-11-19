package services

import (
	"agen_edc/internal/models"
	"agen_edc/internal/repositories"
)

type SignatureService struct {
	repo *repositories.SignatureRepo
}

func NewSignatureService(repo *repositories.SignatureRepo) *SignatureService {
	return &SignatureService{repo: repo}
}

func (s *SignatureService) Create(sig *models.Signature) error {
	return s.repo.Create(sig)
}

func (s *SignatureService) GetByID(id uint) (*models.Signature, error) {
	return s.repo.GetByID(id)
}

func (s *SignatureService) GetByAgentID(agentID uint) ([]models.Signature, error) {
	return s.repo.GetByAgentID(agentID)
}

func (s *SignatureService) Update(sig *models.Signature) error {
	return s.repo.Update(sig)
}

func (s *SignatureService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *SignatureService) GetAll() ([]models.Signature, error) {
	return s.repo.GetAll()
}
