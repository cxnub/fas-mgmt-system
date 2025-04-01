package domain

import "errors"

var (
	InvalidApplicantError                           = errors.New("invalid applicant id")
	InvalidSchemeError                              = errors.New("invalid scheme id")
	InvalidBenefitError                             = errors.New("invalid benefit id")
	InvalidSchemeCriteriaError                      = errors.New("invalid scheme criteria id")
	EmptySchemeCriteriaError                        = errors.New("empty scheme criteria")
	InvalidSchemeCriteriaNameError                  = errors.New("invalid scheme criteria name")
	InvalidSchemeCriteriaAgeValueError              = errors.New("invalid benefit criteria age value")
	InvalidSchemeCriteriaEmploymentStatusValueError = errors.New("invalid benefit criteria employment status value")
	InvalidSchemeCriteriaMaritalStatusValueError    = errors.New("invalid benefit criteria marital status value")
	InvalidSchemeCriteriaHasChildrenValueError      = errors.New("invalid benefit criteria has children value")
	InvalidApplicationError                         = errors.New("invalid application id")
	NotFoundError                                   = errors.New("data not found")
	NoUpdateFieldsError                             = errors.New("no fields to update")
	ApplicantNotFoundError                          = errors.New("applicant not found")
	SchemeNotFoundError                             = errors.New("scheme not found")
	ApplicationNotFoundError                        = errors.New("application not found")
	SchemeNotEligibleError                          = errors.New("scheme not eligible")
	BenefitNotFoundError                            = errors.New("benefit not found")
	SchemeCriteriaNotFoundError                     = errors.New("scheme criteria not found")
)
