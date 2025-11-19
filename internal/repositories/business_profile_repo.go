package repositories

import (
	"agen_edc/internal/models"

	"gorm.io/gorm"
)

type BusinessProfileRepo struct {
	db *gorm.DB
}

func NewBusinessProfileRepo(db *gorm.DB) *BusinessProfileRepo {
	return &BusinessProfileRepo{db: db}
}

func (r *BusinessProfileRepo) Create(profile *models.BusinessProfile) error {
	return r.db.Create(profile).Error
}

func (r *BusinessProfileRepo) GetByID(id uint) (*models.BusinessProfile, error) {
	var profile models.BusinessProfile
	err := r.db.First(&profile, id).Error
	return &profile, err
}

func (r *BusinessProfileRepo) GetByAgentID(agentID uint) ([]models.BusinessProfile, error) {
	var profiles []models.BusinessProfile
	err := r.db.Where("agent_id = ?", agentID).Find(&profiles).Error
	return profiles, err
}

func (r *BusinessProfileRepo) Update(profile *models.BusinessProfile) error {
	return r.db.Save(profile).Error
}

func (r *BusinessProfileRepo) Delete(id uint) error {
	return r.db.Delete(&models.BusinessProfile{}, id).Error
}

func (r *BusinessProfileRepo) GetAll() ([]models.BusinessProfile, error) {
	var profiles []models.BusinessProfile
	err := r.db.Find(&profiles).Error
	return profiles, err
}
