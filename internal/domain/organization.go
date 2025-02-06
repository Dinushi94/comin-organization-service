// internal/domain/organization.go
package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type OrganizationSettings struct {
	Theme          string   `json:"theme"`
	AllowedDomains []string `json:"allowed_domains"`
	MaxUsers       int      `json:"max_users"`
	Features       []string `json:"features"`
	WorkingDays    []string `json:"working_days"`
	WorkingHours   struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"working_hours"`
	TimeZone   string   `json:"timezone"`
	LeaveTypes []string `json:"leave_types"`
}

// Implement sql.Scanner interface
func (s *OrganizationSettings) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, &s)
}

// Implement driver.Valuer interface
func (s OrganizationSettings) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type Organization struct {
	Base
	Name           string               `json:"name" gorm:"not null"`
	Slug           string               `json:"slug" gorm:"unique;not null"`
	Domain         string               `json:"domain"`
	Description    string               `json:"description"`
	Logo           string               `json:"logo"`
	Industry       string               `json:"industry"`
	Size           string               `json:"size"`
	Address        string               `json:"address"`
	Country        string               `json:"country"`
	Phone          string               `json:"phone"`
	Website        string               `json:"website"`
	Settings       OrganizationSettings `json:"settings" gorm:"type:jsonb"`
	Status         string               `json:"status" gorm:"default:'active'"`
	SubscriptionID *uuid.UUID           `json:"subscription_id,omitempty" gorm:"type:uuid"`
	BillingEmail   string               `json:"billing_email"`
}

type OrganizationRequest struct {
	Name         string               `json:"name" binding:"required"`
	Domain       string               `json:"domain"`
	Description  string               `json:"description"`
	Industry     string               `json:"industry"`
	Size         string               `json:"size"`
	Address      string               `json:"address"`
	Country      string               `json:"country"`
	Phone        string               `json:"phone"`
	Website      string               `json:"website"`
	Settings     OrganizationSettings `json:"settings"`
	BillingEmail string               `json:"billing_email"`
}

type OrganizationResponse struct {
	ID          uuid.UUID            `json:"id"`
	Name        string               `json:"name"`
	Slug        string               `json:"slug"`
	Domain      string               `json:"domain"`
	Description string               `json:"description"`
	Logo        string               `json:"logo"`
	Industry    string               `json:"industry"`
	Size        string               `json:"size"`
	Settings    OrganizationSettings `json:"settings"`
	Status      string               `json:"status"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}
