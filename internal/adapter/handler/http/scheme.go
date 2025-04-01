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

// GetScheme godoc
// @Summary	  Get Scheme by ID
// @Description  Retrieve a scheme using its unique identifier.
// @Tags		 schemes
// @Accept	   json
// @Produce	  json
// @Param	  scheme_id   path	  string  true  "Scheme ID" format(uuid)
// @Success	  200  {object}  SchemeResponse  "Successfully retrieved scheme"
// @Failure	  400  {object}  ErrorResponse		  "Validation error occurred"
// @Failure	  404  {object}  ErrorResponse		  "Scheme not found"
// @Failure	  500  {object}  ErrorResponse		  "Internal server error"
// @Router	   /schemes/{scheme_id} [get]
func (h *SchemeHandler) GetScheme(ctx *gin.Context) {
	var req SchemeRequestUri

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

	scheme, err := h.s.GetSchemeByID(ctx, id)

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
// @Router	   /schemes/eligible [get]
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

// CreateScheme godoc
// @Summary	  Create a new scheme
// @Description  Add a new scheme with the provided details.
// @Tags		 schemes
// @Accept	   json
// @Produce	  json
// @Param		CreateSchemeRequest  body	  CreateSchemeRequest  true  "JSON object containing new scheme details"
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

// UpdateScheme godoc
// @Summary	  Update an existing scheme
// @Description  Modify the details of an existing scheme using its unique identifier.
// @Tags		  schemes
// @Accept		  json
// @Produce	  json
// @Param		  scheme_id	   path	string				   true  "Scheme ID" format(uuid)
// @Param		  body	 body	UpdateSchemeRequest	  true  "JSON object with updates to the scheme"
// @Success	  200	  {object} SchemeResponse	"Successfully updated scheme"
// @Failure	  400	  {object} ErrorResponse			"Validation error occurred"
// @Failure	  404	  {object} ErrorResponse			"Scheme not found"
// @Failure	  500	  {object} ErrorResponse			"Internal server error"
// @Router		  /schemes/{scheme_id} [put]
func (h *SchemeHandler) UpdateScheme(ctx *gin.Context) {
	var reqUri SchemeRequestUri
	var req UpdateSchemeRequest

	err := ctx.ShouldBindUri(&reqUri)
	id, err := uuid.Parse(reqUri.ID)
	if err != nil {
		handleError(ctx, domain.InvalidSchemeError)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		validationError(ctx, err, reqUri)
		return
	}

	existingScheme, err := h.s.GetSchemeByID(ctx, id)
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

// DeleteScheme godoc
// @Summary	  Delete a scheme
// @Description  Remove a scheme from the system using its unique identifier.
// @Tags		 schemes
// @Accept	   json
// @Produce	  json
// @Param		scheme_id  path  string  true  "Scheme ID" format(uuid)
// @Success	  200  {object}  Response  "Successfully deleted scheme"
// @Failure	  400  {object}  ErrorResponse	"Validation error occurred"
// @Failure	  404  {object}  ErrorResponse	"Scheme not found"
// @Failure	  500  {object}  ErrorResponse	"Internal server error"
// @Router	   /schemes/{scheme_id} [delete]
func (h *SchemeHandler) DeleteScheme(ctx *gin.Context) {
	var req SchemeRequestUri

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

// AddSchemeBenefit godoc
// @Summary	  Add a benefit to a scheme
// @Description  Attach a new benefit to an existing scheme by specifying the scheme ID.
// @Tags		  schemes
// @Accept		  json
// @Produce	  json
// @Param		  scheme_id  path	string					true  "Scheme ID" format(uuid)
// @Param		  AddSchemeBenefitRequest	   body	AddSchemeBenefitRequest  true  "JSON object with benefit details"
// @Success	  201	   {object}  SchemeBenefitResponse  "Successfully added benefit to scheme"
// @Failure	  400	   {object}  ErrorResponse			  "Validation error occurred"
// @Failure	  404	   {object}  ErrorResponse			  "Scheme not found"
// @Failure	  500	   {object}  ErrorResponse			  "Internal server error"
// @Router		  /schemes/{scheme_id}/benefits [post]
func (h *SchemeHandler) AddSchemeBenefit(ctx *gin.Context) {
	var reqUri SchemeRequestUri
	var req AddSchemeBenefitRequest

	err := ctx.ShouldBindUri(&reqUri)
	if err != nil {
		handleError(ctx, domain.InvalidSchemeError)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		validationError(ctx, err, req)
		return
	}

	schemeID, err := uuid.Parse(reqUri.ID)

	if err != nil {
		handleError(ctx, domain.InvalidSchemeError)
	}

	newBenefit := domain.Benefit{
		Name:     &req.Name,
		Amount:   &req.Amount,
		SchemeID: &schemeID,
	}

	benefit, err := h.s.AddSchemeBenefit(ctx, &newBenefit)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newSchemebenefitResponse(*benefit)
	handleSuccess(ctx, http.StatusCreated, "Successfully added benefit to scheme.", rsp)
}

// UpdateSchemeBenefit godoc
// @Summary	  Update a benefit of a scheme
// @Description  Modify an existing benefit of a scheme by specifying the scheme ID and benefit ID.
// @Tags		 schemes
// @Accept	   json
// @Produce	  json
// @Param		benefit_id		 path	  string				  true  "Benefit ID" format(uuid)
// @Param		UpdateSchemeBenefitRequest body UpdateSchemeBenefitRequest true "JSON object with updated benefit details"
// @Success	  200		{object}  SchemeBenefitResponse   "Successfully updated benefit"
// @Failure	  400		{object}  ErrorResponse		   "Validation error occurred"
// @Failure	  404		{object}  ErrorResponse		   "Benefit or Scheme not found"
// @Failure	  500		{object}  ErrorResponse		   "Internal server error"
// @Router	   /schemes/benefits/{benefit_id} [put]
func (h *SchemeHandler) UpdateSchemeBenefit(ctx *gin.Context) {
	var reqUri BenefitRequestUri
	var req UpdateSchemeBenefitRequest
	var benefit *domain.Benefit

	err := ctx.ShouldBindUri(&reqUri)
	if err != nil {
		handleError(ctx, domain.InvalidSchemeError)
		return
	}

	err = ctx.ShouldBindJSON(&req)

	if err != nil {
		validationError(ctx, err, req)
	}

	id, err := uuid.Parse(reqUri.ID)

	if err != nil {
		handleError(ctx, domain.InvalidBenefitError)
		return
	}

	schemeID, err := uuid.Parse(*req.SchemeID)

	if err != nil {
		handleError(ctx, domain.InvalidSchemeError)
	}

	newBenefit := domain.Benefit{
		ID:       &id,
		Name:     req.Name,
		Amount:   req.Amount,
		SchemeID: &schemeID,
	}

	benefit, err = h.s.UpdateSchemeBenefit(ctx, &newBenefit)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newSchemebenefitResponse(*benefit)
	handleSuccess(ctx, http.StatusOK, "Successfully updated benefit.", rsp)
}

// DeleteSchemeBenefit godoc
// @Summary	  Delete a benefit from a scheme
// @Description  Remove a benefit from a scheme using its unique identifier.
// @Tags		 schemes
// @Accept	   json
// @Produce	  json
// @Param		benefit_id  path  string  true  "Benefit ID" format(uuid)
// @Success	  200  {object}  Response  "Successfully deleted benefit"
// @Failure	  400  {object}  ErrorResponse "Validation error occurred"
// @Failure	  404  {object}  ErrorResponse "Benefit not found"
// @Failure	  500  {object}  ErrorResponse "Internal server error"
// @Router	   /schemes/benefits/{benefit_id} [delete]
func (h *SchemeHandler) DeleteSchemeBenefit(ctx *gin.Context) {
	var req BenefitRequestUri

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		handleError(ctx, domain.InvalidSchemeError)
		return
	}

	id, err := uuid.Parse(req.ID)

	if err != nil {
		handleError(ctx, domain.InvalidBenefitError)
	}

	err = h.s.DeleteSchemeBenefit(ctx, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, http.StatusOK, "Successfully deleted benefit.", nil)
	return
}

// AddSchemeCriteria godoc
// @Summary	  Add a criteria to a scheme
// @Description  Add a new criteria to an existing scheme by specifying the scheme ID.
// @Tags		  schemes
// @Accept		  json
// @Produce	  json
// @Param		  scheme_id  path	string					true  "Scheme ID" format(uuid)
// @Param		  AddSchemeCriteriaRequest	   body	AddSchemeCriteriaRequest  true  "JSON object with criteria details"
// @Success	  201	   {object}  SchemeCriteriaResponse  "Successfully added criteria to scheme"
// @Failure	  400	   {object}  ErrorResponse			  "Validation error occurred"
// @Failure	  404	   {object}  ErrorResponse			  "Scheme not found"
// @Failure	  500	   {object}  ErrorResponse			  "Internal server error"
// @Router		  /schemes/{scheme_id}/criteria [post]
func (h *SchemeHandler) AddSchemeCriteria(ctx *gin.Context) {
	var reqUri SchemeRequestUri
	var req AddSchemeCriteriaRequest

	err := ctx.ShouldBindUri(&reqUri)
	if err != nil {
		validationError(ctx, err, reqUri)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		validationError(ctx, err, req)
		return
	}

	schemeID, err := uuid.Parse(reqUri.ID)
	if err != nil {
		handleError(ctx, domain.InvalidSchemeError)
		return
	}

	newCriteria := domain.SchemeCriteria{
		Name:     &req.Name,
		Value:    &req.Value,
		SchemeID: &schemeID,
	}

	criteria, err := h.s.AddSchemeCriteria(ctx, &newCriteria)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newSchemeCriteriaResponse(*criteria)
	handleSuccess(ctx, http.StatusCreated, "Successfully added criteria to scheme.", rsp)
}

// UpdateSchemeCriteria godoc
// @Summary	  Update a criteria of a scheme
// @Description  Modify an existing criteria of a scheme by specifying the scheme ID and criteria ID.
// @Tags		  schemes
// @Accept		  json
// @Produce	  json
// @Param		  scheme_criteria_id			path	string					true  "Scheme Criteria ID" format(uuid)
// @Param		  UpdateSchemeCriteriaRequest	body	UpdateSchemeCriteriaRequest	true	"JSON object with updated criteria details"
// @Success	  200		{object}  SchemeCriteriaResponse  "Successfully updated criteria"
// @Failure	  400		{object}  ErrorResponse		  "Validation error occurred"
// @Failure	  404		{object}  ErrorResponse		  "Criteria or Scheme not found"
// @Failure	  500		{object}  ErrorResponse		  "Internal server error"
// @Router		  /schemes/criteria/{scheme_criteria_id} [put]
func (h *SchemeHandler) UpdateSchemeCriteria(ctx *gin.Context) {
	var reqUri SchemeCriteriaRequestUri
	var req UpdateSchemeCriteriaRequest

	err := ctx.ShouldBindUri(&reqUri)
	if err != nil {
		handleError(ctx, domain.InvalidSchemeError)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		validationError(ctx, err, req)
		return
	}

	schemeID, err := uuid.Parse(*req.SchemeID)
	if err != nil {
		handleError(ctx, domain.InvalidSchemeError)
		return
	}

	id, err := uuid.Parse(reqUri.ID)
	if err != nil {
		handleError(ctx, domain.InvalidSchemeCriteriaError)
		return
	}

	newCriteria := domain.SchemeCriteria{
		ID:       &id,
		Name:     req.Name,
		Value:    req.Value,
		SchemeID: &schemeID,
	}

	updatedCriteria, err := h.s.UpdateSchemeCriteria(ctx, &newCriteria)
	if err != nil {
		handleError(ctx, err)
		return
	}
	rsp := newSchemeCriteriaResponse(*updatedCriteria)
	handleSuccess(ctx, http.StatusOK, "Successfully updated criteria.", rsp)
}

// DeleteSchemeCriteria godoc
// @Summary	  Delete a criteria from a scheme
// @Description  Remove a criteria from a scheme using its unique identifier.
// @Tags		  schemes
// @Accept		  json
// @Produce	  json
// @Param		  scheme_criteria_id  path  string  true  "Criteria ID" format(uuid)
// @Success	  200  {object}  Response  "Successfully deleted criteria"
// @Failure	  400  {object}  ErrorResponse "Validation error occurred"
// @Failure	  404  {object}  ErrorResponse "Criteria not found"
// @Failure	  500  {object}  ErrorResponse "Internal server error"
// @Router		  /schemes/criteria/{scheme_criteria_id} [delete]
func (h *SchemeHandler) DeleteSchemeCriteria(ctx *gin.Context) {
	var req SchemeCriteriaRequestUri

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		handleError(ctx, domain.InvalidSchemeCriteriaError)
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		handleError(ctx, domain.InvalidSchemeCriteriaError)
		return
	}

	err = h.s.DeleteSchemeCriteria(ctx, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, http.StatusOK, "Successfully deleted criteria.", nil)
}
