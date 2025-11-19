package services

import (
	"agen_edc/internal/models"
	"agen_edc/internal/repositories"
)

type UploadedDocumentService struct {
	repo *repositories.UploadedDocumentRepo
}

func NewUploadedDocumentService(repo *repositories.UploadedDocumentRepo) *UploadedDocumentService {
	return &UploadedDocumentService{repo: repo}
}

func (s *UploadedDocumentService) Create(doc *models.UploadedDocument) error {
	return s.repo.Create(doc)
}

func (s *UploadedDocumentService) GetByID(id uint) (*models.UploadedDocument, error) {
	return s.repo.GetByID(id)
}

func (s *UploadedDocumentService) GetByAgentID(agentID uint) ([]models.UploadedDocument, error) {
	return s.repo.GetByAgentID(agentID)
}

func (s *UploadedDocumentService) Update(doc *models.UploadedDocument) error {
	return s.repo.Update(doc)
}

func (s *UploadedDocumentService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *UploadedDocumentService) GetAll() ([]models.UploadedDocument, error) {
	return s.repo.GetAll()
}
