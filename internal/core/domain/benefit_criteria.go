package domain

import (
	"github.com/google/uuid"
	"time"
)

type BenefitCriteria struct {
	ID        *uuid.UUID
	BenefitID *uuid.UUID
	Name      *string
	Value     *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
