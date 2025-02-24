// internal/service/department_service.go
package service

import (
	"errors"

	"github.com/Axontik/comin-organization-service/internal/domain"
	"github.com/Axontik/comin-organization-service/internal/repository"
	"github.com/google/uuid"
)

type DepartmentService interface {
	Create(orgID uuid.UUID, req *domain.DepartmentRequest) (*domain.Department, error)
	GetByID(orgID, deptID uuid.UUID) (*domain.Department, error)
	Update(orgID, deptID uuid.UUID, req *domain.DepartmentRequest) (*domain.Department, error)
	Delete(orgID, deptID uuid.UUID) error
	List(orgID uuid.UUID) ([]domain.Department, error)
	ListByParentID(orgID, parentID uuid.UUID) ([]domain.Department, error)
}

type departmentService struct {
	deptRepo repository.DepartmentRepository
}

func NewDepartmentService(deptRepo repository.DepartmentRepository) DepartmentService {
	return &departmentService{
		deptRepo: deptRepo,
	}
}

func (s *departmentService) Create(orgID uuid.UUID, req *domain.DepartmentRequest) (*domain.Department, error) {
	dept := &domain.Department{
		OrganizationID: orgID,
		Name:           req.Name,
		Description:    req.Description,
		ParentID:       req.ParentID,
		ManagerID:      req.ManagerID,
		Status:         "active",
	}

	if err := s.deptRepo.Create(dept); err != nil {
		return nil, err
	}

	return dept, nil
}

func (s *departmentService) GetByID(orgID, deptID uuid.UUID) (*domain.Department, error) {
	dept, err := s.deptRepo.GetByID(deptID)
	if err != nil {
		return nil, err
	}

	// Verify department belongs to organization
	if dept.OrganizationID != orgID {
		return nil, errors.New("department not found in organization")
	}

	return dept, nil
}

func (s *departmentService) Update(orgID, deptID uuid.UUID, req *domain.DepartmentRequest) (*domain.Department, error) {
	dept, err := s.GetByID(orgID, deptID)
	if err != nil {
		return nil, err
	}

	dept.Name = req.Name
	dept.Description = req.Description
	dept.ParentID = req.ParentID
	dept.ManagerID = req.ManagerID

	if err := s.deptRepo.Update(dept); err != nil {
		return nil, err
	}

	return dept, nil
}
func (s *departmentService) Delete(orgID, deptID uuid.UUID) error {
	_, err := s.GetByID(orgID, deptID)
	if err != nil {
		return err
	}

	// Check if department has children
	children, err := s.ListByParentID(orgID, deptID)
	if err != nil {
		return err
	}
	if len(children) > 0 {
		return errors.New("cannot delete department with sub-departments")
	}

	return s.deptRepo.Delete(deptID)
}

func (s *departmentService) List(orgID uuid.UUID) ([]domain.Department, error) {
	return s.deptRepo.List(orgID)
}

func (s *departmentService) ListByParentID(orgID, parentID uuid.UUID) ([]domain.Department, error) {
	return s.deptRepo.ListByParentID(parentID)
}
