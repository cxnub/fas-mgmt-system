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
