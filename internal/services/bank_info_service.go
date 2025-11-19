package services

import (
	"agen_edc/internal/models"
	"agen_edc/internal/repositories"
)

type BankInfoService struct {
	repo *repositories.BankInfoRepo
}

func NewBankInfoService(repo *repositories.BankInfoRepo) *BankInfoService {
	return &BankInfoService{repo: repo}
}

func (s *BankInfoService) Create(info *models.BankInfo) error {
	return s.repo.Create(info)
}

func (s *BankInfoService) GetByID(id uint) (*models.BankInfo, error) {
	return s.repo.GetByID(id)
}

func (s *BankInfoService) GetByAgentID(agentID uint) ([]models.BankInfo, error) {
	return s.repo.GetByAgentID(agentID)
}

func (s *BankInfoService) Update(info *models.BankInfo) error {
	return s.repo.Update(info)
}

func (s *BankInfoService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *BankInfoService) GetAll() ([]models.BankInfo, error) {
	return s.repo.GetAll()
}
