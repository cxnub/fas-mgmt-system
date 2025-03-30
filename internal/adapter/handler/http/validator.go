package http

import (
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/go-playground/validator/v10"
	"time"
)

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "date":
		return "Invalid date format, please use the YYYY-MM-DD format."
	case "marital_status":
		return "Invalid marital status, must be either single, married, widowed or divorced."
	case "sex":
		return "Invalid sex, must be either male or female."
	case "relationship_type":
		return "Invalid relationship type, must be either spouse, child, parent or sibling."
	case "employment_status":
		return "Invalid employment status, must be either employed or unemployed."
	default:
		return "Invalid field input."
	}
}

func validateMaritalStatus(f1 validator.FieldLevel) bool {
	status, ok := f1.Field().Interface().(domain.MaritalStatus)
	return ok && status.IsValid()
}

func validateSex(f1 validator.FieldLevel) bool {
	status, ok := f1.Field().Interface().(domain.Sex)
	return ok && status.IsValid()
}

func validateRelationshipType(f1 validator.FieldLevel) bool {
	status, ok := f1.Field().Interface().(domain.RelationshipType)
	return ok && status.IsValid()
}

func validateEmploymentStatus(f1 validator.FieldLevel) bool {
	status, ok := f1.Field().Interface().(domain.EmploymentStatus)
	return ok && status.IsValid()
}

func validateDate(f1 validator.FieldLevel) bool {
	dateStr, ok := f1.Field().Interface().(string)
	if !ok {
		return false
	}
	_, err := time.Parse("2006-01-02", dateStr) // Assuming the date is in "YYYY-MM-DD" format
	return err == nil
}
