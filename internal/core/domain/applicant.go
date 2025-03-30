package domain

import (
	"github.com/google/uuid"
	"time"
)

type EmploymentStatus string

const (
	EmploymentStatusEmployed   EmploymentStatus = "employed"
	EmploymentStatusUnemployed EmploymentStatus = "unemployed"
)

func (es EmploymentStatus) IsValid() bool {
	switch es {
	case EmploymentStatusEmployed, EmploymentStatusUnemployed:
		return true
	default:
		return false
	}
}

type MaritalStatus string

const (
	MaritalStatusSingle  MaritalStatus = "single"
	MaritalStatusMarried MaritalStatus = "married"
	MaritalStatusWidowed MaritalStatus = "widowed"
	MaritalStatusDivorce MaritalStatus = "divorce"
)

func (ms MaritalStatus) IsValid() bool {
	switch ms {
	case MaritalStatusSingle, MaritalStatusMarried, MaritalStatusWidowed, MaritalStatusDivorce:
		return true
	default:
		return false
	}
}

type Sex string

const (
	SexMale   Sex = "male"
	SexFemale Sex = "female"
)

func (s Sex) IsValid() bool {
	switch s {
	case SexMale, SexFemale:
		return true
	default:
		return false
	}
}

type RelationshipType string

const (
	RelationshipTypeSpouse  RelationshipType = "spouse"
	RelationshipTypeChild   RelationshipType = "child"
	RelationshipTypeParent  RelationshipType = "parent"
	RelationshipTypeSibling RelationshipType = "sibling"
)

func (rt RelationshipType) IsValid() bool {
	switch rt {
	case RelationshipTypeSpouse, RelationshipTypeChild, RelationshipTypeParent, RelationshipTypeSibling:
		return true
	default:
		return false
	}
}

type Applicant struct {
	ID               *uuid.UUID
	Name             *string
	EmploymentStatus *EmploymentStatus
	MaritalStatus    *MaritalStatus
	Sex              *Sex
	DateOfBirth      *time.Time
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	Family           map[RelationshipType]*Applicant
}
