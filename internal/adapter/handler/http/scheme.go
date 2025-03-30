package http

import (
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/cxnub/fas-mgmt-system/internal/core/port"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type SchemeHandler struct {
	s port.SchemeService
}

func NewSchemeHandler(s port.SchemeService) *SchemeHandler {
	return &SchemeHandler{s: s}
}

type GetSchemeRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// GetScheme godoc
// @Summary	  Get Scheme by ID
// @Description  Retrieve a scheme using its unique identifier.
// @Tags		 schemes
// @Accept	   json
// @Produce	  json
// @Param		id   path	  string  true  "Scheme ID" format(uuid)
// @Success	  200  {object}  SchemeResponse  "Successfully retrieved scheme"
// @Failure	  400  {object}  ErrorResponse		  "Validation error occurred"
// @Failure	  404  {object}  ErrorResponse		  "Scheme not found"
// @Failure	  500  {object}  ErrorResponse		  "Internal server error"
// @Router	   /schemes/{id} [get]
func (h *SchemeHandler) GetScheme(ctx *gin.Context) {
	var req GetSchemeRequest

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

	scheme, err := h.s.GetSchemeById(ctx, id)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newSchemeResponse(*scheme)
	handleSuccess(ctx, http.StatusOK, "Successfully retrieved scheme.", rsp)
	return
}

// ListSchemes godoc
// @Summary	  List all schemes
// @Description  Retrieve a comprehensive list of all available schemes.
// @Tags		 schemes
// @Accept	   json
// @Produce	  json
// @Success	  200  {array}   SchemeResponse  "Successfully retrieved list of schemes"
// @Failure	  500  {object}  ErrorResponse			"Internal server error"
// @Router	   /schemes [get]
func (h *SchemeHandler) ListSchemes(ctx *gin.Context) {
	result, err := h.s.ListSchemes(ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newSchemesResponse(result)

	handleSuccess(ctx, http.StatusOK, "", rsp)
}

// ListApplicantAvailableSchemes godoc
// @Summary	  List Applicant Available Schemes
// @Description  Retrieve a list of schemes available for a specific applicant using their unique identifier.
// @Tags		 schemes
// @Accept	   json
// @Produce	  json
// @Param		applicant query   string  true  "Applicant ID" format(uuid)
// @Success	  200	   {array}  SchemeResponse  "Successfully retrieved available schemes"
// @Failure	  400	   {object} ErrorResponse		   "Validation error occurred"
// @Failure	  404	   {object} ErrorResponse		   "Applicant not found"
// @Failure	  500	   {object} ErrorResponse		   "Internal server error"
// @Router	   /schemes/applicant [get]
func (h *SchemeHandler) ListApplicantAvailableSchemes(ctx *gin.Context) {

	id, err := uuid.Parse(ctx.Query("applicant"))
	if err != nil {
		handleError(ctx, domain.ApplicantNotFoundError)
		return
	}

	result, err := h.s.ListApplicantAvailableSchemes(ctx, id)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newSchemesResponse(result)

	handleSuccess(ctx, http.StatusOK, "", rsp)
	return
}

type CreateSchemeRequest struct {
	Name string `json:"name" binding:"required"`
}

// CreateScheme godoc
// @Summary	  Create a new scheme
// @Description  Add a new scheme with the provided details.
// @Tags		 schemes
// @Accept	   json
// @Produce	  json
// @Param		body  body	  CreateSchemeRequest  true  "JSON object containing new scheme details"
// @Success	  201  {object}  SchemeResponse  "Successfully created scheme"
// @Failure	  400  {object}  ErrorResponse		  "Validation error occurred"
// @Failure	  500  {object}  ErrorResponse		  "Internal server error"
// @Router	   /schemes [post]
func (h *SchemeHandler) CreateScheme(ctx *gin.Context) {
	var req CreateSchemeRequest
	var newScheme *domain.Scheme

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		validationError(ctx, err, req)
		return
	}

	scheme := domain.Scheme{
		Name: &req.Name,
	}

	newScheme, err = h.s.CreateScheme(ctx, &scheme)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newSchemeResponse(*newScheme)

	handleSuccess(ctx, http.StatusCreated, "Successfully created scheme.", rsp)
	return
}

type UpdateSchemeRequest struct {
	ID   string  `uri:"id" binding:"required,uuid"`
	Name *string `json:"name"`
}

// UpdateScheme godoc
// @Summary	  Update an existing scheme
// @Description  Modify the details of an existing scheme using its unique identifier.
// @Tags		  schemes
// @Accept		  json
// @Produce	  json
// @Param		  id	   path	string				   true  "Scheme ID" format(uuid)
// @Param		  body	 body	UpdateSchemeRequest	  true  "JSON object with updates to the scheme"
// @Success	  200	  {object} SchemeResponse	"Successfully updated scheme"
// @Failure	  400	  {object} ErrorResponse			"Validation error occurred"
// @Failure	  404	  {object} ErrorResponse			"Scheme not found"
// @Failure	  500	  {object} ErrorResponse			"Internal server error"
// @Router		  /schemes/{id} [put]
func (h *SchemeHandler) UpdateScheme(ctx *gin.Context) {
	var req UpdateSchemeRequest

	err := ctx.ShouldBindUri(&req)
	id, err := uuid.Parse(req.ID)
	if err != nil {
		handleError(ctx, domain.InvalidSchemeError)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		validationError(ctx, err, req)
		return
	}

	existingScheme, err := h.s.GetSchemeById(ctx, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	if existingScheme == nil {
		handleError(ctx, domain.SchemeNotFoundError)
		return
	}

	newSchemeValues := domain.Scheme{
		ID:   &id,
		Name: req.Name,
	}

	updatedScheme, err := h.s.UpdateScheme(ctx, &newSchemeValues)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newSchemeResponse(*updatedScheme)

	handleSuccess(ctx, http.StatusOK, "Successfully updated scheme.", rsp)
	return
}

type DeleteSchemeRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// DeleteScheme godoc
// @Summary	  Delete a scheme
// @Description  Remove a scheme from the system using its unique identifier.
// @Tags		 schemes
// @Accept	   json
// @Produce	  json
// @Param		id  path  string  true  "Scheme ID" format(uuid)
// @Success	  200  {object}  Response  "Successfully deleted scheme"
// @Failure	  400  {object}  ErrorResponse	"Validation error occurred"
// @Failure	  404  {object}  ErrorResponse	"Scheme not found"
// @Failure	  500  {object}  ErrorResponse	"Internal server error"
// @Router	   /schemes/{id} [delete]
func (h *SchemeHandler) DeleteScheme(ctx *gin.Context) {
	var req DeleteSchemeRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		validationError(ctx, err, req)
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		handleError(ctx, domain.InvalidSchemeError)
		return
	}

	err = h.s.DeleteScheme(ctx, id)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, http.StatusOK, "Successfully deleted scheme.", nil)
	return
}
