package http

import "github.com/cxnub/fas-mgmt-system/internal/core/domain"

// ===========================================
// ============ Applicant Routes =============
// ===========================================

// ApplicantRequestUri represents URI parameters for an applicant request containing a required ID as a UUID.
type ApplicantRequestUri struct {
	ID string `uri:"id" binding:"required,uuid" example:"b6c29c96-024b-4e70-834b-8e0dd2c66645"`
}

// CreateApplicantRequest represents the required information to create a new applicant in the system.
// The struct requires fields for name, employment status, sex, date of birth, and marital status with validation constraints.
type CreateApplicantRequest struct {
	Name             string                  `json:"name" binding:"required" example:"John Doe"`
	EmploymentStatus domain.EmploymentStatus `json:"employment_status" binding:"required,employment_status" example:"employed"`
	Sex              domain.Sex              `json:"sex" binding:"required,sex" example:"male"`
	DateOfBirth      string                  `json:"date_of_birth" binding:"required,date" example:"1990-01-01"`
	MaritalStatus    domain.MaritalStatus    `json:"marital_status" binding:"required,marital_status" example:"married"`
}

// UpdateApplicantRequest represents a request payload to update an applicant's details.
type UpdateApplicantRequest struct {
	Name             *string                  `json:"name" example:"John Doe" binding:"omitempty"`
	EmploymentStatus *domain.EmploymentStatus `json:"employment_status" binding:"omitempty,employment_status" example:"unemployed"`
	Sex              *domain.Sex              `json:"sex" binding:"omitempty,sex" example:"male"`
	DateOfBirth      *string                  `json:"date_of_birth" binding:"omitempty,date" example:"1990-01-01"`
	MaritalStatus    *domain.MaritalStatus    `json:"marital_status" binding:"omitempty,marital_status" example:"married"`
}

// ===========================================
// =========== Application Routes ============
// ===========================================

// ApplicationRequestUri represents the URI parameters required for an application request, containing a mandatory UUID ID.
type ApplicationRequestUri struct {
	ID string `uri:"id" binding:"required,uuid" example:"fe897b4f-568b-4ea1-8d95-99a91c97faf2"`
}

// CreateApplicationRequest represents a request for creating a new application with mandatory applicant and scheme identifiers.
type CreateApplicationRequest struct {
	ApplicantID string `json:"applicant_id" binding:"required"`
	SchemeID    string `json:"scheme_id" binding:"required"`
}

// UpdateApplicationRequest represents a request to update an application by its ID, ApplicantID, or SchemeID.
type UpdateApplicationRequest struct {
	ApplicantID *string `json:"applicant_id"`
	SchemeID    *string `json:"scheme_id"`
}

// ===========================================
// ============== Scheme Routes ==============
// ===========================================

// SchemeRequestUri represents a structure to bind and validate a URI path parameter for a scheme request.
type SchemeRequestUri struct {
	ID string `uri:"scheme_id" binding:"required,uuid"`
}

// SchemeCriteriaRequestUri represents the request URI structure for scheme criteria, containing a mandatory UUID identifier.
type SchemeCriteriaRequestUri struct {
	ID string `uri:"scheme_criteria_id" binding:"required,uuid"`
}

// BenefitRequestUri represents the URI structure for identifying a specific benefit request by its unique ID.
type BenefitRequestUri struct {
	ID string `uri:"benefit_id" binding:"required,uuid"`
}

// CreateSchemeRequest represents a request payload for creating a new scheme with a mandatory name field.
type CreateSchemeRequest struct {
	Name string `json:"name" binding:"required"`
}

// UpdateSchemeRequest represents a request payload for updating an existing scheme.
type UpdateSchemeRequest struct {
	Name *string `json:"name"`
}

// DeleteSchemeRequest represents a request to delete a scheme.
type DeleteSchemeRequest struct {
}

// AddSchemeBenefitRequest represents a request payload for adding a benefit to a scheme with required details.
type AddSchemeBenefitRequest struct {
	Name   string  `json:"name" binding:"required" example:"CDC Vouchers"`
	Amount float64 `json:"amount" binding:"required" example:"100"`
}

// UpdateSchemeBenefitRequest represents a request structure for updating a scheme benefit.
type UpdateSchemeBenefitRequest struct {
	Name     *string  `json:"name"`
	Amount   *float64 `json:"amount"`
	SchemeID *string  `json:"scheme_id"`
}

// AddSchemeCriteriaRequest represents the request to add a new criteria to an existing scheme.
type AddSchemeCriteriaRequest struct {
	Name  string `json:"name" binding:"required" example:"Age Limit"`
	Value string `json:"value" binding:"required" example:"18-50"`
}

// UpdateSchemeCriteriaRequest represents the payload for updating a scheme criteria.
type UpdateSchemeCriteriaRequest struct {
	Name     *string `json:"name"`
	Value    *string `json:"value"`
	SchemeID *string `json:"scheme_id"`
}
