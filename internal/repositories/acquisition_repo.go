package repositories

import (
	"agen_edc/internal/models"

	"gorm.io/gorm"
)

type AcquisitionRepo struct {
	db *gorm.DB
}

func NewAcquisitionRepo(db *gorm.DB) *AcquisitionRepo {
	return &AcquisitionRepo{db: db}
}

func (r *AcquisitionRepo) Create(acquisition *models.AcquisitionInfo) error {
	return r.db.Create(acquisition).Error
}

func (r *AcquisitionRepo) GetByID(id uint) (*models.AcquisitionInfo, error) {
	var acquisition models.AcquisitionInfo
	err := r.db.First(&acquisition, id).Error
	return &acquisition, err
}

func (r *AcquisitionRepo) GetByAgentID(agentID uint) (*models.AcquisitionInfo, error) {
	var acquisition models.AcquisitionInfo
	err := r.db.Where("agent_id = ?", agentID).First(&acquisition).Error
	return &acquisition, err
}

func (r *AcquisitionRepo) Update(acquisition *models.AcquisitionInfo) error {
	return r.db.Save(acquisition).Error
}

func (r *AcquisitionRepo) Delete(id uint) error {
	return r.db.Delete(&models.AcquisitionInfo{}, id).Error
}

func (r *AcquisitionRepo) GetAll() ([]models.AcquisitionInfo, error) {
	var acquisitions []models.AcquisitionInfo
	err := r.db.Find(&acquisitions).Error
	return acquisitions, err
}
