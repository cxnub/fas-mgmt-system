package util

import (
	"fmt"
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"slices"
	"strconv"
	"strings"
	"time"
)

// CompareNumber checks if a given number satisfies the condition string (e.g., ">=65").
func CompareNumber(condition string, num int) (bool, error) {
	// Possible operators
	operators := []string{">=", "<=", ">", "<", "="}

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
		return false, fmt.Errorf("invalid condition format: %s", condition)
	}

	// Convert the number string to an integer
	value, err := strconv.Atoi(strings.TrimSpace(valueStr))
	if err != nil {
		return false, fmt.Errorf("invalid number in condition: %s", valueStr)
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
	case "=":
		return num == value, nil
	default:
		return false, fmt.Errorf("unsupported operator: %s", operator)
	}
}

func CheckSchemeEligibility(scheme domain.Scheme, applicant *domain.Applicant, relationships map[domain.RelationshipType]*domain.Applicant) bool {
	relations := make([]domain.RelationshipType, len(relationships))

	i := 0
	for k := range relationships {
		relations[i] = k
		i++
	}

	if scheme.Criteria == nil {
		return true
	}

	for _, criterion := range *scheme.Criteria {
		switch *criterion.Name {
		case "employment_status":
			if domain.EmploymentStatus(*criterion.Value) != *applicant.EmploymentStatus {
				return false
			}
		case "marital_status":
			if domain.MaritalStatus(*criterion.Value) != *applicant.MaritalStatus {
				return false
			}
		case "has_children":
			if !slices.Contains(relations, domain.RelationshipTypeChild) {
				return false
			}
		case "age":
			// Calculate applicant's age based on their date of birth
			if applicant.DateOfBirth != nil {
				currentYear := time.Now().Year()
				birthYear := applicant.DateOfBirth.Year()
				age := currentYear - birthYear
				if valid, _ := CompareNumber(*criterion.Value, age); !valid {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}
