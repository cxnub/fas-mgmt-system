package domain

import (
	"github.com/google/uuid"
	"time"
)

type Application struct {
	ID          *uuid.UUID
	ApplicantID *uuid.UUID
	SchemeID    *uuid.UUID
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
