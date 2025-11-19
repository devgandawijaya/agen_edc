package repositories

import (
	"agen_edc/internal/models"

	"gorm.io/gorm"
)

type BankInfoRepo struct {
	db *gorm.DB
}

func NewBankInfoRepo(db *gorm.DB) *BankInfoRepo {
	return &BankInfoRepo{db: db}
}

func (r *BankInfoRepo) Create(bank *models.BankInfo) error {
	return r.db.Create(bank).Error
}

func (r *BankInfoRepo) GetByID(id uint) (*models.BankInfo, error) {
	var bank models.BankInfo
	err := r.db.First(&bank, id).Error
	return &bank, err
}

func (r *BankInfoRepo) GetByAgentID(agentID uint) ([]models.BankInfo, error) {
	var banks []models.BankInfo
	err := r.db.Where("agent_id = ?", agentID).Find(&banks).Error
	return banks, err
}

func (r *BankInfoRepo) Update(bank *models.BankInfo) error {
	return r.db.Save(bank).Error
}

func (r *BankInfoRepo) Delete(id uint) error {
	return r.db.Delete(&models.BankInfo{}, id).Error
}

func (r *BankInfoRepo) GetAll() ([]models.BankInfo, error) {
	var banks []models.BankInfo
	err := r.db.Find(&banks).Error
	return banks, err
}
