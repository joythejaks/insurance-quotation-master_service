package repository

import (
	"github.com/jordisetiawan/insurance-master-service/internal/model"
	"gorm.io/gorm"
)

type CountryRepository interface {
	Create(country *model.Country) error
	FindAll(page, limit int, search string) ([]model.Country, int64, error)
	FindByID(id string) (*model.Country, error)
	Update(country *model.Country) error
	Delete(id string) error
}

type countryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) CountryRepository {
	return &countryRepository{db: db}
}

func (r *countryRepository) Create(country *model.Country) error {
	return r.db.Create(country).Error
}

func (r *countryRepository) FindAll(page, limit int, search string) ([]model.Country, int64, error) {
	var countries []model.Country
	var count int64
	query := r.db.Model(&model.Country{})
	if search != "" {
		query = query.Where("name ILIKE ? OR code ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	query.Count(&count)
	err := query.Offset((page - 1) * limit).Limit(limit).Find(&countries).Error
	return countries, count, err
}

func (r *countryRepository) FindByID(id string) (*model.Country, error) {
	var country model.Country
	if err := r.db.Where("id = ?", id).First(&country).Error; err != nil {
		return nil, err
	}
	return &country, nil
}

func (r *countryRepository) Update(country *model.Country) error {
	return r.db.Save(country).Error
}

func (r *countryRepository) Delete(id string) error {
	return r.db.Delete(&model.Country{}, "id = ?", id).Error
}
