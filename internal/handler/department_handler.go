// internal/handler/department_handler.go
package handler

import (
	"net/http"

	"github.com/Axontik/comin-organization-service/internal/domain"
	"github.com/Axontik/comin-organization-service/internal/errors"
	"github.com/Axontik/comin-organization-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DepartmentHandler struct {
	deptService service.DepartmentService
}

func NewDepartmentHandler(deptService service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{deptService: deptService}
}

// Create handles department creation
// @Summary Create department
// @Description Create a new department in an organization
// @Tags departments
// @Accept json
// @Produce json
// @Param org_id path string true "Organization ID"
// @Param department body domain.DepartmentRequest true "Department details"
// @Success 201 {object} domain.Department
// @Failure 400,404 {object} errors.AppError
// @Router /organizations/{org_id}/departments [post]
func (h *DepartmentHandler) Create(c *gin.Context) {
	orgID, err := uuid.Parse(c.Param("org_id"))
	if err != nil {
		c.Error(errors.NewBadRequestError("invalid organization id"))
		return
	}

	var req domain.DepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.NewBadRequestError(err.Error()))
		return
	}

	dept, err := h.deptService.Create(orgID, &req)
	if err != nil {
		c.Error(errors.NewInternalServerError("failed to create department"))
		return
	}

	c.JSON(http.StatusCreated, dept)
}

// GetByID handles fetching a single department
// @Summary Get department by ID
// @Description Get department details by ID
// @Tags departments
// @Produce json
// @Param org_id path string true "Organization ID"
// @Param id path string true "Department ID"
// @Success 200 {object} domain.Department
// @Failure 404 {object} errors.AppError
// @Router /organizations/{org_id}/departments/{id} [get]
func (h *DepartmentHandler) GetByID(c *gin.Context) {
	orgID, err := uuid.Parse(c.Param("org_id"))
	if err != nil {
		c.Error(errors.NewBadRequestError("invalid organization id"))
		return
	}

	deptID, err := uuid.Parse(c.Param("dept_id"))
	if err != nil {
		c.Error(errors.NewBadRequestError("invalid department id"))
		return
	}

	dept, err := h.deptService.GetByID(orgID, deptID)
	if err != nil {
		c.Error(errors.NewNotFoundError("department not found"))
		return
	}

	c.JSON(http.StatusOK, dept)
}

// Update handles department updates
// @Summary Update department
// @Description Update department details
// @Tags departments
// @Accept json
// @Produce json
// @Param org_id path string true "Organization ID"
// @Param id path string true "Department ID"
// @Param department body domain.DepartmentRequest true "Department details"
// @Success 200 {object} domain.Department
// @Failure 400,404 {object} errors.AppError
// @Router /organizations/{org_id}/departments/{id} [put]
func (h *DepartmentHandler) Update(c *gin.Context) {
	orgID, err := uuid.Parse(c.Param("org_id"))
	if err != nil {
		c.Error(errors.NewBadRequestError("invalid organization id"))
		return
	}

	deptID, err := uuid.Parse(c.Param("dept_id"))
	if err != nil {
		c.Error(errors.NewBadRequestError("invalid department id"))
		return
	}

	var req domain.DepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.NewBadRequestError(err.Error()))
		return
	}

	dept, err := h.deptService.Update(orgID, deptID, &req)
	if err != nil {
		c.Error(errors.NewInternalServerError("failed to update department"))
		return
	}

	c.JSON(http.StatusOK, dept)
}

// List handles fetching all departments in an organization
// @Summary List departments
// @Description Get all departments in an organization
// @Tags departments
// @Produce json
// @Param org_id path string true "Organization ID"
// @Success 200 {array} domain.Department
// @Failure 404 {object} errors.AppError
// @Router /organizations/{org_id}/departments [get]
func (h *DepartmentHandler) List(c *gin.Context) {
	orgID, err := uuid.Parse(c.Param("org_id"))
	if err != nil {
		c.Error(errors.NewBadRequestError("invalid organization id"))
		return
	}

	depts, err := h.deptService.List(orgID)
	if err != nil {
		c.Error(errors.NewInternalServerError("failed to fetch departments"))
		return
	}

	c.JSON(http.StatusOK, depts)
}

// Delete handles department deletion
// @Summary Delete department
// @Description Delete a department
// @Tags departments
// @Param org_id path string true "Organization ID"
// @Param id path string true "Department ID"
// @Success 204 "No Content"
// @Failure 400,404 {object} errors.AppError
// @Router /organizations/{org_id}/departments/{id} [delete]
func (h *DepartmentHandler) Delete(c *gin.Context) {
	orgID, err := uuid.Parse(c.Param("org_id"))
	if err != nil {
		c.Error(errors.NewBadRequestError("invalid organization id"))
		return
	}

	deptID, err := uuid.Parse(c.Param("dept_id"))
	if err != nil {
		c.Error(errors.NewBadRequestError("invalid department id"))
		return
	}

	if err := h.deptService.Delete(orgID, deptID); err != nil {
		c.Error(errors.NewInternalServerError("failed to delete department"))
		return
	}

	c.Status(http.StatusNoContent)
}
