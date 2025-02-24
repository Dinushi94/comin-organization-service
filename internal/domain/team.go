// internal/domain/team.go
package domain

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	Base
	OrganizationID uuid.UUID  `json:"organization_id" gorm:"type:uuid;not null"`
	DepartmentID   *uuid.UUID `json:"department_id,omitempty" gorm:"type:uuid"`
	Name           string     `json:"name" gorm:"not null"`
	Description    string     `json:"description"`
	LeadID         *uuid.UUID `json:"lead_id,omitempty" gorm:"type:uuid"`
	Status         string     `json:"status" gorm:"default:'active'"`
}

type TeamMember struct {
	Base
	TeamID   uuid.UUID `json:"team_id" gorm:"type:uuid;not null"`
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Role     string    `json:"role" gorm:"not null"`
	JoinedAt time.Time `json:"joined_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type TeamRequest struct {
	Name         string     `json:"name" binding:"required"`
	Description  string     `json:"description"`
	DepartmentID *uuid.UUID `json:"department_id,omitempty"`
	LeadID       *uuid.UUID `json:"lead_id,omitempty"`
}
