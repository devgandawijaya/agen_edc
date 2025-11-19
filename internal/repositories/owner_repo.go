package repositories

import (
	"agen_edc/internal/models"

	"gorm.io/gorm"
)

type OwnerRepo struct {
	db *gorm.DB
}

func NewOwnerRepo(db *gorm.DB) *OwnerRepo {
	return &OwnerRepo{db: db}
}

func (r *OwnerRepo) Create(owner *models.Owner) error {
	return r.db.Create(owner).Error
}

func (r *OwnerRepo) GetByID(id uint) (*models.Owner, error) {
	var owner models.Owner
	err := r.db.First(&owner, id).Error
	return &owner, err
}

func (r *OwnerRepo) GetByAgentID(agentID uint) ([]models.Owner, error) {
	var owners []models.Owner
	err := r.db.Where("agent_id = ?", agentID).Find(&owners).Error
	return owners, err
}

func (r *OwnerRepo) Update(owner *models.Owner) error {
	return r.db.Save(owner).Error
}

func (r *OwnerRepo) Delete(id uint) error {
	return r.db.Delete(&models.Owner{}, id).Error
}

func (r *OwnerRepo) GetAll() ([]models.Owner, error) {
	var owners []models.Owner
	err := r.db.Find(&owners).Error
	return owners, err
}
