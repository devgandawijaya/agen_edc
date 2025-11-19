package repositories

import (
	"agen_edc/internal/models"

	"gorm.io/gorm"
)

type UploadedDocumentRepo struct {
	db *gorm.DB
}

func NewUploadedDocumentRepo(db *gorm.DB) *UploadedDocumentRepo {
	return &UploadedDocumentRepo{db: db}
}

func (r *UploadedDocumentRepo) Create(doc *models.UploadedDocument) error {
	return r.db.Create(doc).Error
}

func (r *UploadedDocumentRepo) GetByID(id uint) (*models.UploadedDocument, error) {
	var doc models.UploadedDocument
	err := r.db.First(&doc, id).Error
	return &doc, err
}

func (r *UploadedDocumentRepo) GetByAgentID(agentID uint) ([]models.UploadedDocument, error) {
	var docs []models.UploadedDocument
	err := r.db.Where("agent_id = ?", agentID).Find(&docs).Error
	return docs, err
}

func (r *UploadedDocumentRepo) Update(doc *models.UploadedDocument) error {
	return r.db.Save(doc).Error
}

func (r *UploadedDocumentRepo) Delete(id uint) error {
	return r.db.Delete(&models.UploadedDocument{}, id).Error
}

func (r *UploadedDocumentRepo) GetAll() ([]models.UploadedDocument, error) {
	var docs []models.UploadedDocument
	err := r.db.Find(&docs).Error
	return docs, err
}
