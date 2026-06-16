package service

import (
	"github.com/google/uuid"
	"github.com/jordisetiawan/insurance-master-service/internal/dto"
	"github.com/jordisetiawan/insurance-master-service/internal/model"
	"github.com/jordisetiawan/insurance-master-service/internal/repository"
)

type CountryService interface {
	GetCountries(page, limit int, search string) ([]model.Country, int64, error)
	GetCountry(id string) (*model.Country, error)
	CreateCountry(req dto.CreateCountryRequest) (*model.Country, error)
	UpdateCountry(id string, req dto.UpdateCountryRequest) (*model.Country, error)
	DeleteCountry(id string) error
}

type countryService struct {
	repo repository.CountryRepository
}

func NewCountryService(
	repo repository.CountryRepository,
) CountryService {
	return &countryService{
		repo: repo,
	}
}

func (s *countryService) GetCountries(
	page,
	limit int,
	search string,
) ([]model.Country, int64, error) {
	return s.repo.FindAll(page, limit, search)
}

func (s *countryService) GetCountry(
	id string,
) (*model.Country, error) {
	return s.repo.FindByID(id)
}

func (s *countryService) CreateCountry(
	req dto.CreateCountryRequest,
) (*model.Country, error) {

	country := &model.Country{
		ID:       uuid.New(),
		Code:     req.Code,
		Name:     req.Name,
		IsActive: true,
	}

	if err := s.repo.Create(country); err != nil {
		return nil, err
	}

	return country, nil
}

func (s *countryService) UpdateCountry(
	id string,
	req dto.UpdateCountryRequest,
) (*model.Country, error) {

	country, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		country.Name = req.Name
	}

	country.IsActive = req.IsActive

	if err := s.repo.Update(country); err != nil {
		return nil, err
	}

	return country, nil
}

func (s *countryService) DeleteCountry(
	id string,
) error {
	return s.repo.Delete(id)
}
