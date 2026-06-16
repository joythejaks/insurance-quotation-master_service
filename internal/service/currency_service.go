package service

import (
	"github.com/google/uuid"
	"github.com/jordisetiawan/insurance-master-service/internal/dto"
	"github.com/jordisetiawan/insurance-master-service/internal/model"
	"github.com/jordisetiawan/insurance-master-service/internal/repository"
)

type CurrencyService interface {
	GetCurrencies(page, limit int, search string) ([]model.Currency, int64, error)
	GetCurrency(id string) (*model.Currency, error)
	CreateCurrency(req dto.CreateCurrencyRequest) (*model.Currency, error)
	UpdateCurrency(id string, req dto.UpdateCurrencyRequest) (*model.Currency, error)
	DeleteCurrency(id string) error
}

type currencyService struct {
	repo repository.CurrencyRepository
}

func NewCurrencyService(
	repo repository.CurrencyRepository,
) CurrencyService {
	return &currencyService{
		repo: repo,
	}
}

func (s *currencyService) GetCurrencies(
	page,
	limit int,
	search string,
) ([]model.Currency, int64, error) {
	return s.repo.FindAll(page, limit, search)
}

func (s *currencyService) GetCurrency(
	id string,
) (*model.Currency, error) {
	return s.repo.FindByID(id)
}

func (s *currencyService) CreateCurrency(
	req dto.CreateCurrencyRequest,
) (*model.Currency, error) {

	currency := &model.Currency{
		ID:       uuid.New(),
		Code:     req.Code,
		Name:     req.Name,
		Symbol:   req.Symbol,
		IsActive: true,
	}

	if err := s.repo.Create(currency); err != nil {
		return nil, err
	}

	return currency, nil
}

func (s *currencyService) UpdateCurrency(
	id string,
	req dto.UpdateCurrencyRequest,
) (*model.Currency, error) {

	currency, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		currency.Name = req.Name
	}

	currency.Symbol = req.Symbol
	currency.IsActive = req.IsActive

	if err := s.repo.Update(currency); err != nil {
		return nil, err
	}

	return currency, nil
}

func (s *currencyService) DeleteCurrency(
	id string,
) error {
	return s.repo.Delete(id)
}
