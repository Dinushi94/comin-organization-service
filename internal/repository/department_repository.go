// internal/repository/department_repository.go
package repository

import (
	"github.com/Axontik/comin-organization-service/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	Create(dept *domain.Department) error
	GetByID(id uuid.UUID) (*domain.Department, error)
	Update(dept *domain.Department) error
	Delete(id uuid.UUID) error
	List(organizationID uuid.UUID) ([]domain.Department, error)
	ListByParentID(parentID uuid.UUID) ([]domain.Department, error)
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db: db}
}

func (r *departmentRepository) Create(dept *domain.Department) error {
	return r.db.Create(dept).Error
}

func (r *departmentRepository) GetByID(id uuid.UUID) (*domain.Department, error) {
	var dept domain.Department
	err := r.db.First(&dept, "id = ?", id).Error
	return &dept, err
}

func (r *departmentRepository) Update(dept *domain.Department) error {
	return r.db.Save(dept).Error
}

func (r *departmentRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Department{}, "id = ?", id).Error
}

func (r *departmentRepository) List(organizationID uuid.UUID) ([]domain.Department, error) {
	var depts []domain.Department
	err := r.db.Where("organization_id = ?", organizationID).Find(&depts).Error
	return depts, err
}

func (r *departmentRepository) ListByParentID(parentID uuid.UUID) ([]domain.Department, error) {
	var depts []domain.Department
	err := r.db.Where("parent_id = ?", parentID).Find(&depts).Error
	return depts, err
}
