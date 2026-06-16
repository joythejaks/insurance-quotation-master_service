package repository

import (
	"github.com/jordisetiawan/insurance-master-service/internal/model"
	"gorm.io/gorm"
)

type OccupationRepository interface {
	Create(occupation *model.Occupation) error
	FindAll(page, limit int, search string) ([]model.Occupation, int64, error)
	FindByID(id string) (*model.Occupation, error)
	Update(occupation *model.Occupation) error
	Delete(id string) error
}

type occupationRepository struct {
	db *gorm.DB
}

func NewOccupationRepository(db *gorm.DB) OccupationRepository {
	return &occupationRepository{db: db}
}

func (r *occupationRepository) Create(occupation *model.Occupation) error {
	return r.db.Create(occupation).Error
}

func (r *occupationRepository) FindAll(page, limit int, search string) ([]model.Occupation, int64, error) {
	var occupations []model.Occupation
	var count int64
	query := r.db.Model(&model.Occupation{})
	if search != "" {
		query = query.Where("name ILIKE ? OR code ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	query.Count(&count)
	err := query.Offset((page - 1) * limit).Limit(limit).Find(&occupations).Error
	return occupations, count, err
}

func (r *occupationRepository) FindByID(id string) (*model.Occupation, error) {
	var occupation model.Occupation
	if err := r.db.Where("id = ?", id).First(&occupation).Error; err != nil {
		return nil, err
	}
	return &occupation, nil
}

func (r *occupationRepository) Update(occupation *model.Occupation) error {
	return r.db.Save(occupation).Error
}

func (r *occupationRepository) Delete(id string) error {
	return r.db.Delete(&model.Occupation{}, "id = ?", id).Error
}
