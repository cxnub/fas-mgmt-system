package http

import (
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/cxnub/fas-mgmt-system/internal/core/port"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// ApplicantHandler provides HTTP handler methods for managing applicants using an ApplicantService.
type ApplicantHandler struct {
	s port.ApplicantService
}

// NewApplicantHandler initializes a new ApplicantHandler with the provided ApplicantService.
func NewApplicantHandler(s port.ApplicantService) *ApplicantHandler {
	return &ApplicantHandler{s: s}
}

// GetApplicant godoc
// @Summary	  Retrieve Applicant by ID
// @Description  Retrieves the details of a single applicant using their unique identifier.
// @Tags		 Applicants
// @Accept	   json
// @Produce	  json
// @Param		id   path	  string  true  "Applicant ID"
// @Success	  200  {object}  ApplicantResponse  "Successfully retrieved applicant."
// @Failure	  400  {object}  ErrorResponse	  "Bad Request"
// @Failure	  404  {object}  ErrorResponse	  "Applicant Not Found"
// @Failure	  500  {object}  ErrorResponse	  "Internal Server Error"
// @Router	   /applicants/{id} [get]
func (h *ApplicantHandler) GetApplicant(ctx *gin.Context) {
	var req ApplicantRequestUri

	err := ctx.ShouldBindUri(&req)

	if err != nil {
		validationError(ctx, err, req)
		return
	}

	id, err := uuid.Parse(req.ID)

	if err != nil {
		handleError(ctx, err)
		return
	}

	applicant, err := h.s.GetApplicantById(ctx, id)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newApplicantResponse(*applicant)
	handleSuccess(ctx, http.StatusOK, "Successfully retrieved applicant.", rsp)
	return
}

// ListApplicants godoc
// @Summary		List All Applicants
// @Description	Retrieves and returns a list of all registered applicants.
// @Tags		   Applicants
// @Accept		 json
// @Produce		json
// @Success		200  {array}   ApplicantResponse "Successfully retrieved list of applicants."
// @Failure		500  {object}  ErrorResponse	 "Internal Server Error"
// @Router		 /applicants [get]
func (h *ApplicantHandler) ListApplicants(ctx *gin.Context) {
	applicants, err := h.s.ListApplicants(ctx)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newApplicantsResponse(applicants)
	handleSuccess(ctx, http.StatusOK, "Successfully retrieved applicants.", rsp)
	return
}

// CreateApplicant godoc
// @Summary	  Create a new Applicant
// @Description  Handles the creation of a new applicant by accepting necessary data and storing it in the system.
// @Tags		 Applicants
// @Accept	   json
// @Produce	  json
// @Param		CreateApplicantRequest  body	  CreateApplicantRequest  true  "Payload for creating a new applicant"
// @Success	  201   {object}  ApplicantResponse  "Successfully created applicant."
// @Failure	  400   {object}  ErrorResponse	  "Bad Request"
// @Failure	  500   {object}  ErrorResponse	  "Internal Server Error"
// @Router	   /applicants [post]
func (h *ApplicantHandler) CreateApplicant(ctx *gin.Context) {
	var req CreateApplicantRequest
	var newApplicant *domain.Applicant

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		validationError(ctx, err, req)
		return
	}

	dob, err := time.Parse("2006-01-02", req.DateOfBirth)

	applicant := domain.Applicant{
		Name:             &req.Name,
		EmploymentStatus: &req.EmploymentStatus,
		Sex:              &req.Sex,
		DateOfBirth:      &dob,
		MaritalStatus:    &req.MaritalStatus,
	}

	newApplicant, err = h.s.CreateApplicant(ctx, &applicant)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newApplicantResponse(*newApplicant)

	handleSuccess(ctx, http.StatusCreated, "Successfully created applicant.", rsp)
	return
}

// UpdateApplicant godoc
// @Summary	  Update an Applicant
// @Description  Updates the specified applicant's details based on the provided payload.
// @Tags		 Applicants
// @Accept	   json
// @Produce	  json
// @Param		id					  path	  string				  true   "Applicant ID"
// @Param		UpdateApplicantRequest  body	  UpdateApplicantRequest  true   "Payload for updating an applicant"
// @Success	  200					 {object}  ApplicantResponse  "Successfully updated applicant."
// @Failure	  400					 {object}  ErrorResponse	  "Bad Request"
// @Failure	  404					 {object}  ErrorResponse	  "Applicant Not Found"
// @Failure	  500					 {object}  ErrorResponse	  "Internal Server Error"
// @Router	   /applicants/{id}		[put]
func (h *ApplicantHandler) UpdateApplicant(ctx *gin.Context) {
	var reqUri ApplicantRequestUri
	var req UpdateApplicantRequest

	err := ctx.ShouldBindUri(&reqUri)
	id, err := uuid.Parse(reqUri.ID)
	if err != nil {
		handleError(ctx, domain.InvalidApplicantError)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		validationError(ctx, err, req)
		return
	}

	existingApplicant, err := h.s.GetApplicantById(ctx, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	if existingApplicant == nil {
		handleError(ctx, domain.ApplicantNotFoundError)
		return
	}

	newApplicantValues := domain.Applicant{
		ID:               &id,
		Sex:              req.Sex,
		MaritalStatus:    req.MaritalStatus,
		EmploymentStatus: req.EmploymentStatus,
		Name:             req.Name,
	}

	if req.DateOfBirth != nil {
		dob, parseErr := time.Parse("2006-01-02", *req.DateOfBirth)
		if parseErr == nil {
			newApplicantValues.DateOfBirth = &dob
		}
	}

	updatedApplicant, err := h.s.UpdateApplicant(ctx, &newApplicantValues)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newApplicantResponse(*updatedApplicant)

	handleSuccess(ctx, http.StatusOK, "Successfully updated applicant.", rsp)
	return
}

// DeleteApplicant godoc
// @Summary	  Delete an Applicant
// @Description  Deletes the applicant with the specified ID from the system.
// @Tags		 Applicants
// @Accept	   json
// @Produce	  json
// @Param		id   path	  string  true  "Applicant ID"
// @Success	  200  {object}  Response  "Successfully deleted applicant."
// @Failure	  400  {object}  ErrorResponse	"Bad Request"
// @Failure	  404  {object}  ErrorResponse	"Applicant Not Found"
// @Failure	  500  {object}  ErrorResponse	"Internal Server Error"
// @Router	   /applicants/{id} [delete]
func (h *ApplicantHandler) DeleteApplicant(ctx *gin.Context) {
	var req ApplicantRequestUri

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		validationError(ctx, err, req)
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		handleError(ctx, domain.InvalidApplicantError)
		return
	}

	err = h.s.DeleteApplicant(ctx, id)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, http.StatusOK, "Successfully deleted applicant.", nil)
	return
}
