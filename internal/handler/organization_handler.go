// internal/handler/organization_handler.go
package handler

import (
	"net/http"
	"strconv"

	"github.com/Axontik/comin-organization-service/internal/domain"
	"github.com/Axontik/comin-organization-service/internal/errors"
	"github.com/Axontik/comin-organization-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationHandler struct {
	orgService service.OrganizationService
}

func NewOrganizationHandler(orgService service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{orgService: orgService}
}

// @Summary Create organization
// @Description Create a new organization
// @Tags organizations
// @Accept json
// @Produce json
// @Param organization body domain.OrganizationRequest true "Organization details"
// @Success 201 {object} domain.OrganizationResponse
// @Failure 400 {object} ErrorResponse
// @Router /organizations [post]
func (h *OrganizationHandler) Create(c *gin.Context) {
	var req domain.OrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.NewBadRequestError("Invalid request payload"))
		return
	}

	org, err := h.orgService.Create(&req)
	if err != nil {
		// Handle specific errors
		switch e := err.(type) {
		case *errors.AppError:
			c.Error(e)
		default:
			c.Error(errors.NewInternalServerError("Failed to create organization"))
		}
		return
	}

	c.JSON(http.StatusCreated, org)
}

// @Summary Get organization
// @Description Get organization by ID
// @Tags organizations
// @Produce json
// @Param org_id path string true "Organization ID"
// @Success 200 {object} domain.OrganizationResponse
// @Failure 404 {object} ErrorResponse
// @Router /organizations/{org_id} [get]
func (h *OrganizationHandler) GetByID(c *gin.Context) {
	org_id, err := uuid.Parse(c.Param("org_id"))
	if err != nil {
		c.Error(errors.NewBadRequestError("Invalid organization ID"))
		return
	}

	org, err := h.orgService.GetByID(org_id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.Error(errors.NewNotFoundError("Organization not found"))
		default:
			c.Error(errors.NewInternalServerError("Failed to fetch organization"))
		}
		return
	}

	c.JSON(http.StatusOK, org)
}

// @Summary Update organization
// @Description Update organization by ID
// @Tags organizations
// @Produce json
// @Param org_id path string true "Organization ID"
// @Param body body domain.OrganizationRequest true "Organization details"
// @Success 200 {object} domain.OrganizationResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /organizations/{org_id} [put]
func (h *OrganizationHandler) Update(c *gin.Context) {
	org_id, err := uuid.Parse(c.Param("org_id"))
	if err != nil {
		c.Error(errors.NewBadRequestError("Invalid organization ID"))
		return
	}

	var req domain.OrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.NewBadRequestError("Invalid request body"))
		return
	}

	org, err := h.orgService.Update(org_id, &req)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.Error(errors.NewNotFoundError("Organization not found"))
		default:
			c.Error(errors.NewInternalServerError("Failed to update organization"))
		}
		return
	}

	c.JSON(http.StatusOK, org)
}

// @Summary Delete organization
// @Description Delete organization by ID
// @Tags organizations
// @Produce json
// @Param org_id path string true "Organization ID"
// @Success 204 {object} nil
// @Failure 404 {object} ErrorResponse
// @Router /organizations/{org_id} [delete]
func (h *OrganizationHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("org_id"))
	if err != nil {
		c.Error(errors.NewBadRequestError("Invalid organization ID"))
		return
	}

	if err := h.orgService.Delete(id); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.Error(errors.NewNotFoundError("Organization not found"))
		default:
			c.Error(errors.NewInternalServerError("Failed to delete organization"))
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary List organizations
// @Description List organizations
// @Tags organizations
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} domain.OrganizationResponse
// @Failure 400 {object} ErrorResponse
// @Router /organizations [get]
func (h *OrganizationHandler) List(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.Error(errors.NewBadRequestError("Invalid limit"))
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.Error(errors.NewBadRequestError("Invalid offset"))
		return
	}

	orgs, err := h.orgService.List(limit, offset)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.Error(errors.NewNotFoundError("No organizations found"))
		default:
			c.Error(errors.NewInternalServerError("Failed to list organizations"))
		}
		return
	}

	c.JSON(http.StatusOK, orgs)
}
