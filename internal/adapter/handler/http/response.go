package http

import (
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
)

// Response represents a response body format
type Response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

// newResponse is a helper function to create a response body
func newResponse(success bool, message string, data any) Response {
	return Response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

// ApplicantResponse represents the response containing applicant's personal and status information.
type ApplicantResponse struct {
	ID               string `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	Name             string `json:"name" example:"John Doe"`
	EmploymentStatus string `json:"employment_status" example:"employed"`
	MaritalStatus    string `json:"marital_status" example:"married"`
	Sex              string `json:"sex" example:"male"`
	DateOfBirth      string `json:"date_of_birth" example:"2000-01-01"`
	CreatedAt        string `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt        string `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}

func newApplicantResponse(applicant domain.Applicant) ApplicantResponse {
	return ApplicantResponse{
		ID:               applicant.ID.String(),
		Name:             *applicant.Name,
		EmploymentStatus: string(*applicant.EmploymentStatus),
		MaritalStatus:    string(*applicant.MaritalStatus),
		Sex:              string(*applicant.Sex),
		DateOfBirth:      applicant.DateOfBirth.String(),
		CreatedAt:        applicant.CreatedAt.String(),
		UpdatedAt:        applicant.UpdatedAt.String(),
	}
}

// ApplicantsResponse represents a collection of applicant responses.
type ApplicantsResponse struct {
	Applicants []ApplicantResponse `json:"applicants" example:"applicants"`
}

func newApplicantsResponse(applicants []domain.Applicant) ApplicantsResponse {
	var applicantResponses []ApplicantResponse
	for _, a := range applicants {
		applicantResponses = append(applicantResponses, newApplicantResponse(a))
	}
	return ApplicantsResponse{
		Applicants: applicantResponses,
	}
}

// SchemeBenefitResponse represents a response structure containing benefit details like name and amount for a scheme.
type SchemeBenefitResponse struct {
	Name   string  `json:"name" example:"CDC Vouchers"`
	Amount float64 `json:"amount" example:"1000000"`
}

func newSchemeBenefitResponse(benefit []domain.Benefit) []SchemeBenefitResponse {
	var schemeBenefitResponses []SchemeBenefitResponse

	for _, b := range benefit {
		schemeBenefitResponses = append(schemeBenefitResponses, SchemeBenefitResponse{
			Name:   *b.Name,
			Amount: *b.Amount,
		})
	}

	return schemeBenefitResponses
}

// SchemeCriteriaResponse represents a response containing a criterion's name and value associated with a scheme.
type SchemeCriteriaResponse struct {
	Name  string `json:"name" example:"employment_status"`
	Value string `json:"value" example:"unemployed"`
}

func newSchemeCriteriaResponse(criteria []domain.SchemeCriteria) []SchemeCriteriaResponse {
	var schemeCriteriaResponses []SchemeCriteriaResponse

	for _, sc := range criteria {
		schemeCriteriaResponses = append(schemeCriteriaResponses, SchemeCriteriaResponse{
			Name:  *sc.Name,
			Value: *sc.Value,
		})
	}

	return schemeCriteriaResponses
}

// SchemeResponse represents the response structure containing details of a scheme, including ID, name, criteria, and benefits.
type SchemeResponse struct {
	ID       string                   `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	Name     string                   `json:"name" example:"Retrenchment Assistance Scheme"`
	Criteria []SchemeCriteriaResponse `json:"criteria"`
	Benefits []SchemeBenefitResponse  `json:"benefits"`
}

func newSchemeResponse(scheme domain.Scheme) SchemeResponse {
	return SchemeResponse{
		ID:       scheme.ID.String(),
		Name:     *scheme.Name,
		Criteria: newSchemeCriteriaResponse(*scheme.Criteria),
		Benefits: newSchemeBenefitResponse(*scheme.Benefits),
	}
}

// SchemesResponse represents the response structure containing a list of schemes with their respective details.
type SchemesResponse struct {
	Schemes []SchemeResponse `json:"schemes"`
}

func newSchemesResponse(schemes []domain.Scheme) SchemesResponse {
	var schemeResponses []SchemeResponse
	for _, s := range schemes {
		schemeResponses = append(schemeResponses, newSchemeResponse(s))
	}
	return SchemesResponse{
		Schemes: schemeResponses,
	}
}

// ApplicationResponse represents the response structure containing application details.
type ApplicationResponse struct {
	ID          string `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	ApplicantID string `json:"applicant_id" example:"00000000-0000-0000-0000-000000000000"`
	SchemeID    string `json:"scheme_id" example:"00000000-0000-0000-0000-000000000000"`
	CreatedAt   string `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt   string `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}

func newApplicationResponse(application domain.Application) ApplicationResponse {
	return ApplicationResponse{
		ID:          application.ID.String(),
		ApplicantID: application.ApplicantID.String(),
		SchemeID:    application.SchemeID.String(),
		CreatedAt:   application.CreatedAt.String(),
		UpdatedAt:   application.UpdatedAt.String(),
	}
}

// ApplicationsResponse represents a collection of application responses.
type ApplicationsResponse struct {
	Applications []ApplicationResponse `json:"applications"`
}

func newApplicationsResponse(applications []domain.Application) ApplicationsResponse {
	var applicationResponses []ApplicationResponse
	for _, a := range applications {
		applicationResponses = append(applicationResponses, newApplicationResponse(a))
	}
	return ApplicationsResponse{
		Applications: applicationResponses,
	}
}

// ErrorResponse represents a generic error response body format
type ErrorResponse struct {
	Success bool              `json:"success" example:"false"`
	Message string            `json:"message" example:"Error message"`
	Errors  map[string]string `json:"errors,omitempty" example:"field:error description"`
}

// newErrorResponse is a helper function to create an error response body
func newErrorResponse(msg string) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: msg,
	}
}

// newErrorResponse creates and returns a ValidationErrorResponse with the provided message and an empty errors map.
func newValidationErrorResponse(errorMap map[string]string) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: "Validation error",
		Errors:  errorMap,
	}
}
