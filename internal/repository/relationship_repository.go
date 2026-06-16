package repository

import (
	"github.com/jordisetiawan/insurance-master-service/internal/model"
	"gorm.io/gorm"
)

type RelationshipRepository interface {
	Create(relationship *model.Relationship) error
	FindAll(page, limit int, search string) ([]model.Relationship, int64, error)
	FindByID(id string) (*model.Relationship, error)
	Update(relationship *model.Relationship) error
	Delete(id string) error
}

type relationshipRepository struct {
	db *gorm.DB
}

func NewRelationshipRepository(db *gorm.DB) RelationshipRepository {
	return &relationshipRepository{db: db}
}

func (r *relationshipRepository) Create(relationship *model.Relationship) error {
	return r.db.Create(relationship).Error
}

func (r *relationshipRepository) FindAll(page, limit int, search string) ([]model.Relationship, int64, error) {
	var relationships []model.Relationship
	var count int64
	query := r.db.Model(&model.Relationship{})
	if search != "" {
		query = query.Where("name ILIKE ? OR code ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	query.Count(&count)
	err := query.Offset((page - 1) * limit).Limit(limit).Find(&relationships).Error
	return relationships, count, err
}

func (r *relationshipRepository) FindByID(id string) (*model.Relationship, error) {
	var relationship model.Relationship
	if err := r.db.Where("id = ?", id).First(&relationship).Error; err != nil {
		return nil, err
	}
	return &relationship, nil
}

func (r *relationshipRepository) Update(relationship *model.Relationship) error {
	return r.db.Save(relationship).Error
}

func (r *relationshipRepository) Delete(id string) error {
	return r.db.Delete(&model.Relationship{}, "id = ?", id).Error
}
