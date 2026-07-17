package service

import (
	"github.com/google/uuid"
	"github.com/jordisetiawan/insurance-master-service/internal/dto"
	"github.com/jordisetiawan/insurance-master-service/internal/model"
	"github.com/jordisetiawan/insurance-master-service/internal/repository"
)

type RelationshipService interface {
	GetRelationships(page, limit int, search string) ([]model.Relationship, int64, error)
	GetRelationship(id string) (*model.Relationship, error)
	CreateRelationship(req dto.CreateRelationshipRequest) (*model.Relationship, error)
	UpdateRelationship(id string, req dto.UpdateRelationshipRequest) (*model.Relationship, error)
	DeleteRelationship(id string) error
}

type relationshipService struct {
	repo repository.RelationshipRepository
}

func NewRelationshipService(
	repo repository.RelationshipRepository,
) RelationshipService {
	return &relationshipService{
		repo: repo,
	}
}

func (s *relationshipService) GetRelationships(
	page,
	limit int,
	search string,
) ([]model.Relationship, int64, error) {
	return s.repo.FindAll(page, limit, search)
}

func (s *relationshipService) GetRelationship(
	id string,
) (*model.Relationship, error) {
	return s.repo.FindByID(id)
}

func (s *relationshipService) CreateRelationship(
	req dto.CreateRelationshipRequest,
) (*model.Relationship, error) {

	relationship := &model.Relationship{
		ID:       uuid.New(),
		Code:     req.Code,
		Name:     req.Name,
		IsActive: true,
	}

	if err := s.repo.Create(relationship); err != nil {
		return nil, err
	}

	return relationship, nil
}

func (s *relationshipService) UpdateRelationship(
	id string,
	req dto.UpdateRelationshipRequest,
) (*model.Relationship, error) {

	relationship, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		relationship.Name = req.Name
	}

	if req.IsActive != nil {
		relationship.IsActive = *req.IsActive
	}

	if err := s.repo.Update(relationship); err != nil {
		return nil, err
	}

	return relationship, nil
}

func (s *relationshipService) DeleteRelationship(
	id string,
) error {
	return s.repo.Delete(id)
}
