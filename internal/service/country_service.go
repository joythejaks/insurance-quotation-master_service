package service

import (
	"github.com/jordisetiawan/insurance-master-service/internal/repository"
)

type CountryService struct {
	repo repository.CountryRepository
}

func NewCountryService(
	repo repository.CountryRepository,
) *CountryService {
	return &CountryService{
		repo: repo,
	}
}
