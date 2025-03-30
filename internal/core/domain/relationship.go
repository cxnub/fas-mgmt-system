package domain

import (
	"github.com/google/uuid"
	"time"
)

type Relationship struct {
	ID               *uuid.UUID
	ApplicantAID     *uuid.UUID
	ApplicantBID     *uuid.UUID
	RelationshipType *RelationshipType
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
}
