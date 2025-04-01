package http

import (
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"net/http"
)

// errorStatusMap is a map of defined error messages and their corresponding http status codes
var errorMap = map[error]struct {
	StatusCode int
	Message    string
}{
	domain.InvalidApplicantError: {
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid applicant id.",
	},
	domain.InvalidSchemeError: {
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid scheme id.",
	},
	domain.InvalidBenefitError: {
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid benefit id.",
	},
	domain.InvalidSchemeCriteriaError: {
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid scheme criteria id.",
	},
	domain.InvalidApplicationError: {
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid application id.",
	},
	domain.NotFoundError: {
		StatusCode: http.StatusNotFound,
		Message:    "Data not found.",
	},
	domain.ApplicantNotFoundError: {
		StatusCode: http.StatusNotFound,
		Message:    "Applicant not found.",
	},
	domain.SchemeNotFoundError: {
		StatusCode: http.StatusNotFound,
		Message:    "Scheme not found.",
	},
	domain.SchemeCriteriaNotFoundError: {
		StatusCode: http.StatusNotFound,
		Message:    "Scheme criteria not found.",
	},
	domain.InvalidSchemeCriteriaNameError: {
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid scheme criteria name, only employment_status, marital_status, has_children, or age are allowed.",
	},
	domain.InvalidSchemeCriteriaAgeValueError: {
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid scheme criteria age value, must start with an operator and ends with a number. Valid operators are: >, >=, <, <=, and ==. (e.g. >25, >=25, <25, <=25, ==25).",
	},
	domain.InvalidSchemeCriteriaEmploymentStatusValueError: {
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid scheme criteria employment status value, must be either employed or unemployed.",
	},
	domain.InvalidSchemeCriteriaMaritalStatusValueError: {
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid scheme criteria marital status value, must be either single, married, widowed or divorced.",
	},
	domain.InvalidSchemeCriteriaHasChildrenValueError: {
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid scheme criteria has children value, must be either true or false.",
	},
	domain.InvalidSchemeCriteriaError: {
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid scheme criteria.",
	},
	domain.BenefitNotFoundError: {
		StatusCode: http.StatusNotFound,
		Message:    "Benefit not found.",
	},
	domain.ApplicationNotFoundError: {
		StatusCode: http.StatusNotFound,
		Message:    "Application not found.",
	},
	domain.NoUpdateFieldsError: {
		StatusCode: http.StatusBadRequest,
		Message:    "No fields to update.",
	},
	domain.SchemeNotEligibleError: {
		StatusCode: http.StatusBadRequest,
		Message:    "Applicant does not meet the eligibility criteria for the scheme.",
	},
}
