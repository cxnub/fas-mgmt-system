package domain

import (
	"github.com/google/uuid"
	"time"
)

type Scheme struct {
	ID        *uuid.UUID
	Name      *string
	Benefits  *[]Benefit
	Criteria  *[]SchemeCriteria
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
