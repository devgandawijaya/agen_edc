package repositories

import (
	"agen_edc/internal/models"

	"gorm.io/gorm"
)

type SignatureRepo struct {
	db *gorm.DB
}

func NewSignatureRepo(db *gorm.DB) *SignatureRepo {
	return &SignatureRepo{db: db}
}

func (r *SignatureRepo) Create(sig *models.Signature) error {
	return r.db.Create(sig).Error
}

func (r *SignatureRepo) GetByID(id uint) (*models.Signature, error) {
	var sig models.Signature
	err := r.db.First(&sig, id).Error
	return &sig, err
}

func (r *SignatureRepo) GetByAgentID(agentID uint) ([]models.Signature, error) {
	var sigs []models.Signature
	err := r.db.Where("agent_id = ?", agentID).Find(&sigs).Error
	return sigs, err
}

func (r *SignatureRepo) Update(sig *models.Signature) error {
	return r.db.Save(sig).Error
}

func (r *SignatureRepo) Delete(id uint) error {
	return r.db.Delete(&models.Signature{}, id).Error
}

func (r *SignatureRepo) GetAll() ([]models.Signature, error) {
	var sigs []models.Signature
	err := r.db.Find(&sigs).Error
	return sigs, err
}
