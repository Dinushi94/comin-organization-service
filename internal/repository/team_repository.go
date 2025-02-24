// internal/repository/team_repository.go
package repository

import (
	"github.com/Axontik/comin-organization-service/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamRepository interface {
	Create(team *domain.Team) error
	GetByID(id uuid.UUID) (*domain.Team, error)
	Update(team *domain.Team) error
	Delete(id uuid.UUID) error
	List(organizationID uuid.UUID) ([]domain.Team, error)
	ListByDepartment(departmentID uuid.UUID) ([]domain.Team, error)
	AddMember(member *domain.TeamMember) error
	UpdateMember(member *domain.TeamMember) error
	RemoveMember(teamID, userID uuid.UUID) error
	ListMembers(teamID uuid.UUID) ([]domain.TeamMember, error)
}

type teamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &teamRepository{db: db}
}

func (r *teamRepository) Create(team *domain.Team) error {
	return r.db.Create(team).Error
}

func (r *teamRepository) GetByID(id uuid.UUID) (*domain.Team, error) {
	var team domain.Team
	err := r.db.First(&team, "id = ?", id).Error
	return &team, err
}

func (r *teamRepository) Update(team *domain.Team) error {
	return r.db.Save(team).Error
}

func (r *teamRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Team{}, "id = ?", id).Error
}

func (r *teamRepository) List(organizationID uuid.UUID) ([]domain.Team, error) {
	var teams []domain.Team
	err := r.db.Where("organization_id = ?", organizationID).Find(&teams).Error
	return teams, err
}

func (r *teamRepository) ListByDepartment(departmentID uuid.UUID) ([]domain.Team, error) {
	var teams []domain.Team
	err := r.db.Where("department_id = ?", departmentID).Find(&teams).Error
	return teams, err
}

func (r *teamRepository) AddMember(member *domain.TeamMember) error {
	return r.db.Create(member).Error
}

func (r *teamRepository) UpdateMember(member *domain.TeamMember) error {
	return r.db.Save(member).Error
}

func (r *teamRepository) RemoveMember(teamID, userID uuid.UUID) error {
	return r.db.Where("team_id = ? AND user_id = ?", teamID, userID).Delete(&domain.TeamMember{}).Error
}

func (r *teamRepository) ListMembers(teamID uuid.UUID) ([]domain.TeamMember, error) {
	var members []domain.TeamMember
	err := r.db.Where("team_id = ?", teamID).Find(&members).Error
	return members, err
}
