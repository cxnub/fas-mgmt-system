package domain

import (
	"github.com/google/uuid"
	"time"
)

type SchemeCriteria struct {
	ID        *uuid.UUID
	SchemeID  *uuid.UUID
	Name      *string
	Value     *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
