package repository

import (
	"github.com/jordisetiawan/insurance-master-service/internal/model"
	"gorm.io/gorm"
)

type CurrencyRepository interface {
	Create(currency *model.Currency) error
	FindAll(page, limit int, search string) ([]model.Currency, int64, error)
	FindByID(id string) (*model.Currency, error)
	Update(currency *model.Currency) error
	Delete(id string) error
}

type currencyRepository struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) CurrencyRepository {
	return &currencyRepository{db: db}
}

func (r *currencyRepository) Create(currency *model.Currency) error {
	return r.db.Create(currency).Error
}

func (r *currencyRepository) FindAll(page, limit int, search string) ([]model.Currency, int64, error) {
	var currencies []model.Currency
	var count int64
	query := r.db.Model(&model.Currency{})
	if search != "" {
		query = query.Where("name ILIKE ? OR code ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	query.Count(&count)
	err := query.Offset((page - 1) * limit).Limit(limit).Find(&currencies).Error
	return currencies, count, err
}

func (r *currencyRepository) FindByID(id string) (*model.Currency, error) {
	var currency model.Currency
	if err := r.db.Where("id = ?", id).First(&currency).Error; err != nil {
		return nil, err
	}
	return &currency, nil
}

func (r *currencyRepository) Update(currency *model.Currency) error {
	return r.db.Save(currency).Error
}

func (r *currencyRepository) Delete(id string) error {
	return r.db.Delete(&model.Currency{}, "id = ?", id).Error
}
