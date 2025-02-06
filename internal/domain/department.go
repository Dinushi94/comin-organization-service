// internal/domain/department.go
package domain

import (
	"github.com/google/uuid"
)

type Department struct {
	Base
	OrganizationID uuid.UUID  `json:"organization_id" gorm:"type:uuid;not null"`
	Name           string     `json:"name" gorm:"not null"`
	Description    string     `json:"description"`
	ParentID       *uuid.UUID `json:"parent_id,omitempty" gorm:"type:uuid"`
	Status         string     `json:"status" gorm:"default:'active'"`
	ManagerID      *uuid.UUID `json:"manager_id,omitempty" gorm:"type:uuid"`
}

type DepartmentRequest struct {
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description"`
	ParentID    *uuid.UUID `json:"parent_id,omitempty"`
	ManagerID   *uuid.UUID `json:"manager_id,omitempty"`
}
