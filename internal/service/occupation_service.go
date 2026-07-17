package service

import (
	"github.com/google/uuid"
	"github.com/jordisetiawan/insurance-master-service/internal/dto"
	"github.com/jordisetiawan/insurance-master-service/internal/model"
	"github.com/jordisetiawan/insurance-master-service/internal/repository"
)

type OccupationService interface {
	GetOccupations(page, limit int, search string) ([]model.Occupation, int64, error)
	GetOccupation(id string) (*model.Occupation, error)
	CreateOccupation(req dto.CreateOccupationRequest) (*model.Occupation, error)
	UpdateOccupation(id string, req dto.UpdateOccupationRequest) (*model.Occupation, error)
	DeleteOccupation(id string) error
}

type occupationService struct {
	repo repository.OccupationRepository
}

func NewOccupationService(
	repo repository.OccupationRepository,
) OccupationService {
	return &occupationService{
		repo: repo,
	}
}

func (s *occupationService) GetOccupations(
	page,
	limit int,
	search string,
) ([]model.Occupation, int64, error) {
	return s.repo.FindAll(page, limit, search)
}

func (s *occupationService) GetOccupation(
	id string,
) (*model.Occupation, error) {
	return s.repo.FindByID(id)
}

func (s *occupationService) CreateOccupation(
	req dto.CreateOccupationRequest,
) (*model.Occupation, error) {

	occupation := &model.Occupation{
		ID:       uuid.New(),
		Code:     req.Code,
		Name:     req.Name,
		IsActive: true,
	}

	if err := s.repo.Create(occupation); err != nil {
		return nil, err
	}

	return occupation, nil
}

func (s *occupationService) UpdateOccupation(
	id string,
	req dto.UpdateOccupationRequest,
) (*model.Occupation, error) {

	occupation, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		occupation.Name = req.Name
	}

	if req.IsActive != nil {
		occupation.IsActive = *req.IsActive
	}

	if err := s.repo.Update(occupation); err != nil {
		return nil, err
	}

	return occupation, nil
}

func (s *occupationService) DeleteOccupation(
	id string,
) error {
	return s.repo.Delete(id)
}
