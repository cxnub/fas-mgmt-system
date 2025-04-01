package util

import (
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"slices"
	"strconv"
	"strings"
	"time"
)

// Possible operators
var operators = []string{">=", "<=", ">", "<", "=="}

// CompareNumber checks if a given number satisfies the condition string (e.g., ">=65").
func CompareNumber(condition string, num int) (bool, error) {
	var operator string
	var valueStr string

	// Find which operator exists in the condition string
	for _, op := range operators {
		if strings.HasPrefix(condition, op) {
			operator = op
			valueStr = strings.TrimPrefix(condition, op)
			break
		}
	}

	if operator == "" {
		return false, domain.InvalidSchemeCriteriaAgeValueError
	}

	// Convert the number string to an integer
	value, err := strconv.Atoi(strings.TrimSpace(valueStr))

	if err != nil {
		return false, domain.InvalidSchemeCriteriaAgeValueError
	}

	// Perform the comparison
	switch operator {
	case ">=":
		return num >= value, nil
	case "<=":
		return num <= value, nil
	case ">":
		return num > value, nil
	case "<":
		return num < value, nil
	case "==":
		return num == value, nil
	default:
		return false, domain.InvalidSchemeCriteriaAgeValueError
	}
}

func CheckSchemeEligibility(scheme domain.Scheme, applicant *domain.Applicant, relationships map[domain.RelationshipType]*domain.Applicant) bool {
	if scheme.Criteria == nil {
		return true
	}

	relations := make([]domain.RelationshipType, 0, len(relationships))
	for k := range relationships {
		relations = append(relations, k)
	}

	for _, criterion := range *scheme.Criteria {
		criterionName := strings.ToLower(strings.TrimSpace(*criterion.Name))
		criterionValue := strings.ToLower(strings.TrimSpace(*criterion.Value))
		switch criterionName {
		case "employment_status":
			if domain.EmploymentStatus(criterionValue) != *applicant.EmploymentStatus {
				return false
			}
		case "marital_status":
			if domain.MaritalStatus(criterionValue) != *applicant.MaritalStatus {
				return false
			}
		case "has_children":
			if criterionValue == "true" && !slices.Contains(relations, domain.RelationshipTypeChild) {
				return false
			}
		case "age":
			if applicant.DateOfBirth != nil {
				age := time.Now().Year() - applicant.DateOfBirth.Year()
				if valid, _ := CompareNumber(criterionValue, age); !valid {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}

// IsValidCriteria checks if the given criteria is valid and can be used.
func IsValidCriteria(criterion *domain.SchemeCriteria) *error {
	if criterion == nil || criterion.Name == nil || criterion.Value == nil {
		return &domain.EmptySchemeCriteriaError
	}

	// Trim and convert the criterion name to lowercase for comparison
	criterionName := strings.ToLower(strings.TrimSpace(*criterion.Name))

	// Define a map of valid criteria names and their corresponding validation functions
	validCriteria := map[string]func(string) *error{
		"employment_status": func(value string) *error {
			if !domain.EmploymentStatus(value).IsValid() {
				return &domain.InvalidSchemeCriteriaEmploymentStatusValueError
			}
			return nil
		},
		"marital_status": func(value string) *error {
			if !domain.MaritalStatus(value).IsValid() {
				return &domain.InvalidSchemeCriteriaMaritalStatusValueError
			}
			return nil
		},
		"has_children": func(value string) *error {
			if value != "true" && value != "false" {
				return &domain.InvalidSchemeCriteriaHasChildrenValueError
			}
			return nil
		},
		"age": func(value string) *error {
			if _, err := CompareNumber(value, 0); err != nil {
				return &domain.InvalidSchemeCriteriaAgeValueError
			}
			return nil
		},
	}

	//
	// Retrieve the validation function for the given criteria name and check if it exists
	validate, exists := validCriteria[criterionName]
	if !exists {
		return &domain.InvalidSchemeCriteriaNameError
	}

	return validate(strings.ToLower(strings.TrimSpace(*criterion.Value)))
}
