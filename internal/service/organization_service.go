// internal/service/organization_service.go
package service

import (
	"errors"
	"strings"

	"github.com/Axontik/comin-organization-service/internal/domain"
	"github.com/Axontik/comin-organization-service/internal/repository"
	"github.com/google/uuid"
)

type OrganizationService interface {
	Create(req *domain.OrganizationRequest) (*domain.OrganizationResponse, error)
	GetByID(id uuid.UUID) (*domain.OrganizationResponse, error)
	Update(id uuid.UUID, req *domain.OrganizationRequest) (*domain.OrganizationResponse, error)
	Delete(id uuid.UUID) error
	List(limit, offset int) ([]domain.OrganizationResponse, error)
	UpdateSettings(id uuid.UUID, settings domain.OrganizationSettings) error
	UpdateStatus(id uuid.UUID, status string) error
}

type organizationService struct {
	repo repository.OrganizationRepository
}

func NewOrganizationService(repo repository.OrganizationRepository) OrganizationService {
	return &organizationService{repo: repo}
}

func (s *organizationService) Create(req *domain.OrganizationRequest) (*domain.OrganizationResponse, error) {
	org := &domain.Organization{
		Name:        req.Name,
		Slug:        generateSlug(req.Name),
		Domain:      req.Domain,
		Description: req.Description,
		Industry:    req.Industry,
		Size:        req.Size,
		Address:     req.Address,
		Country:     req.Country,
		Phone:       req.Phone,
		Website:     req.Website,
		Settings:    req.Settings,
		Status:      "active",
	}

	if err := s.repo.Create(org); err != nil {
		return nil, err
	}

	return toOrganizationResponse(org), nil
}

func (s *organizationService) GetByID(id uuid.UUID) (*domain.OrganizationResponse, error) {
	org, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return toOrganizationResponse(org), nil
}

func (s *organizationService) Update(id uuid.UUID, req *domain.OrganizationRequest) (*domain.OrganizationResponse, error) {
	org, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	org.Name = req.Name
	org.Domain = req.Domain
	org.Description = req.Description
	org.Industry = req.Industry
	org.Size = req.Size
	org.Address = req.Address
	org.Country = req.Country
	org.Phone = req.Phone
	org.Website = req.Website
	org.Settings = req.Settings

	if err := s.repo.Update(org); err != nil {
		return nil, err
	}

	return toOrganizationResponse(org), nil
}

func (s *organizationService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *organizationService) List(limit, offset int) ([]domain.OrganizationResponse, error) {
	orgs, err := s.repo.List(limit, offset)
	if err != nil {
		return nil, err
	}

	var responses []domain.OrganizationResponse
	for _, org := range orgs {
		responses = append(responses, *toOrganizationResponse(&org))
	}
	return responses, nil
}

func (s *organizationService) UpdateSettings(id uuid.UUID, settings domain.OrganizationSettings) error {
	return s.repo.UpdateSettings(id, settings)
}

func (s *organizationService) UpdateStatus(id uuid.UUID, status string) error {
	validStatuses := map[string]bool{
		"active":    true,
		"inactive":  true,
		"suspended": true,
	}

	if !validStatuses[status] {
		return errors.New("invalid status")
	}

	return s.repo.UpdateStatus(id, status)
}

func generateSlug(name string) string {
	// Convert to lowercase and replace spaces with hyphens
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove special characters
	slug = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			return r
		}
		return -1
	}, slug)

	return slug
}

func toOrganizationResponse(org *domain.Organization) *domain.OrganizationResponse {
	return &domain.OrganizationResponse{
		ID:          org.ID,
		Name:        org.Name,
		Slug:        org.Slug,
		Domain:      org.Domain,
		Description: org.Description,
		Logo:        org.Logo,
		Industry:    org.Industry,
		Size:        org.Size,
		Settings:    org.Settings,
		Status:      org.Status,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   org.UpdatedAt,
	}
}
