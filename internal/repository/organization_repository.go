// internal/repository/organization_repository.go
package repository

import (
	"github.com/Axontik/comin-organization-service/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	Create(org *domain.Organization) error
	GetByID(id uuid.UUID) (*domain.Organization, error)
	GetBySlug(slug string) (*domain.Organization, error)
	GetByDomain(domain string) (*domain.Organization, error)
	Update(org *domain.Organization) error
	Delete(id uuid.UUID) error
	List(limit, offset int) ([]domain.Organization, error)
	UpdateSettings(id uuid.UUID, settings domain.OrganizationSettings) error
	UpdateStatus(id uuid.UUID, status string) error
}

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{db: db}
}

func (r *organizationRepository) Create(org *domain.Organization) error {
	return r.db.Create(org).Error
}

func (r *organizationRepository) GetByID(id uuid.UUID) (*domain.Organization, error) {
	var org domain.Organization
	err := r.db.First(&org, "id = ?", id).Error
	return &org, err
}

func (r *organizationRepository) GetBySlug(slug string) (*domain.Organization, error) {
	var org domain.Organization
	err := r.db.First(&org, "slug = ?", slug).Error
	return &org, err
}

func (r *organizationRepository) GetByDomain(domains string) (*domain.Organization, error) {
	var org domain.Organization
	err := r.db.First(&org, "domais = ?", domains).Error
	return &org, err
}

func (r *organizationRepository) Update(org *domain.Organization) error {
	return r.db.Save(org).Error
}

func (r *organizationRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Organization{}, "id = ?", id).Error
}

func (r *organizationRepository) List(limit, offset int) ([]domain.Organization, error) {
	var orgs []domain.Organization
	err := r.db.Limit(limit).Offset(offset).Find(&orgs).Error
	return orgs, err
}

func (r *organizationRepository) UpdateSettings(id uuid.UUID, settings domain.OrganizationSettings) error {
	return r.db.Model(&domain.Organization{}).Where("id = ?", id).Update("settings", settings).Error
}

func (r *organizationRepository) UpdateStatus(id uuid.UUID, status string) error {
	return r.db.Model(&domain.Organization{}).Where("id = ?", id).Update("status", status).Error
}
