package domain

import "errors"

var (
	InvalidApplicantError    = errors.New("invalid applicant id")
	InvalidSchemeError       = errors.New("invalid scheme id")
	InvalidApplicationError  = errors.New("invalid application id")
	NotFoundError            = errors.New("data not found")
	NoUpdateFieldsError      = errors.New("no fields to update")
	ApplicantNotFoundError   = errors.New("applicant not found")
	SchemeNotFoundError      = errors.New("scheme not found")
	ApplicationNotFoundError = errors.New("application not found")
	SchemeNotEligibleError   = errors.New("scheme not eligible")
)
