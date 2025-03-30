package domain

import (
	"github.com/google/uuid"
	"time"
)

type Benefit struct {
	ID        *uuid.UUID
	SchemeID  *uuid.UUID
	Name      *string
	Amount    *float64
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
