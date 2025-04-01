package http

import (
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/cxnub/fas-mgmt-system/internal/core/port"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// ApplicationHandler provides handlers for managing application-related operations through ApplicationService.
type ApplicationHandler struct {
	s port.ApplicationService
}

// NewApplicationHandler initializes a new ApplicationHandler with the provided ApplicationService.
func NewApplicationHandler(s port.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{s: s}
}

// GetApplication godoc
//
// @Summary Retrieve application by ID
// @Description Get details of an application by its unique ID.
// @Tags Applications
// @Accept json
// @Produce json
// @Param id path string true "Application ID"
// @Success 200 {object} ApplicationResponse "Application retrieved successfully."
// @Failure 400 {object} ErrorResponse "Invalid UUID or bad input."
// @Failure 404 {object} ErrorResponse "Application not found."
// @Router /applications/{id} [get]
func (h *ApplicationHandler) GetApplication(ctx *gin.Context) {
	var reqUri ApplicationRequestUri

	err := ctx.ShouldBindUri(&reqUri)

	if err != nil {
		validationError(ctx, err, reqUri)
		return
	}

	id, err := uuid.Parse(reqUri.ID)

	if err != nil {
		handleError(ctx, domain.InvalidApplicationError)
		return
	}

	application, err := h.s.GetApplicationById(ctx, id)

	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, application)
	return
}

// ListApplications godoc
//
// @Summary List all applications
// @Description Retrieve all applications present in the system.
// @Tags Applications
// @Accept json
// @Produce json
// @Success 200 {array} ApplicationsResponse "Applications retrieved successfully."
// @Failure 500 {object} ErrorResponse "Internal server error."
// @Router /applications [get]
func (h *ApplicationHandler) ListApplications(ctx *gin.Context) {
	applications, err := h.s.ListApplications(ctx)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newApplicationsResponse(applications)
	handleSuccess(ctx, http.StatusOK, "Successfully retrieved applications.", rsp)
	return
}

// CreateApplication godoc
//
// @Summary Create a new application
// @Description Create a new application using the specified applicant ID and scheme ID.
// @Tags Applications
// @Accept json
// @Produce json
// @Param CreateApplicationRequest body CreateApplicationRequest true "Application creation payload"
// @Success 201 {object} ApplicationResponse "Application created successfully."
// @Failure 400 {object} ErrorResponse "Invalid input data."
// @Failure 500 {object} ErrorResponse "Internal server error."
// @Router /applications [post]
func (h *ApplicationHandler) CreateApplication(ctx *gin.Context) {
	var req CreateApplicationRequest
	var newApplication *domain.Application

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		validationError(ctx, err, req)
		return
	}

	applicantID, err := uuid.Parse(req.ApplicantID)
	if err != nil {
		handleError(ctx, domain.InvalidApplicantError)
		return
	}

	schemeID, err := uuid.Parse(req.SchemeID)
	if err != nil {
		handleError(ctx, domain.InvalidSchemeError)
		return
	}

	application := domain.Application{
		ApplicantID: &applicantID,
		SchemeID:    &schemeID,
	}

	newApplication, err = h.s.CreateApplication(ctx, &application)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newApplicationResponse(*newApplication)
	handleSuccess(ctx, http.StatusCreated, "Successfully created application.", rsp)
	return
}

// UpdateApplication godoc
//
// @Summary Update an application by ID
// @Description Update an application's details such as ApplicantID or SchemeID.
// @Tags Applications
// @Accept json
// @Produce json
// @Param id path string true "Application ID"
// @Param UpdateApplicationRequest body UpdateApplicationRequest true "Application update payload"
// @Success 200 {object} ApplicationResponse "Application updated successfully."
// @Failure 400 {object} ErrorResponse "Invalid input data."
// @Failure 404 {object} ErrorResponse "Application not found."
// @Failure 500 {object} ErrorResponse "Internal server error."
// @Router /applications/{id} [put]
func (h *ApplicationHandler) UpdateApplication(ctx *gin.Context) {
	var reqUri ApplicationRequestUri
	var req UpdateApplicationRequest

	err := ctx.ShouldBindUri(&req)
	id, err := uuid.Parse(reqUri.ID)
	if err != nil {
		handleError(ctx, domain.InvalidApplicationError)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		validationError(ctx, err, req)
		return
	}

	existingApplication, err := h.s.GetApplicationById(ctx, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	if existingApplication == nil {
		handleError(ctx, domain.ApplicationNotFoundError)
		return
	}

	newApplicationValues := domain.Application{
		ID:       &id,
		SchemeID: existingApplication.SchemeID,
	}

	updatedApplication, err := h.s.UpdateApplication(ctx, &newApplicationValues)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newApplicationResponse(*updatedApplication)
	handleSuccess(ctx, http.StatusOK, "Successfully updated application.", rsp)
	return
}

// DeleteApplication godoc
//
// @Summary Delete an application by ID
// @Description Remove an application from the system using its unique ID.
// @Tags Applications
// @Accept json
// @Produce json
// @Param id path string true "Application ID"
// @Success 200 {object} Response "Application deleted successfully."
// @Failure 400 {object} ErrorResponse "Invalid UUID or bad input."
// @Failure 404 {object} ErrorResponse "Application not found."
// @Failure 500 {object} ErrorResponse "Internal server error."
// @Router /applications/{id} [delete]
func (h *ApplicationHandler) DeleteApplication(ctx *gin.Context) {
	var req ApplicationRequestUri

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		validationError(ctx, err, req)
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		handleError(ctx, domain.InvalidApplicationError)
		return
	}

	err = h.s.DeleteApplication(ctx, id)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, http.StatusOK, "Successfully deleted application.", nil)
	return
}
