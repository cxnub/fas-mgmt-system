package port

import (
	"context"
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/google/uuid"
)

type SchemeRepository interface {
	GetSchemeById(ctx context.Context, id uuid.UUID) (*domain.Scheme, error)
	ListSchemes(ctx context.Context) ([]domain.Scheme, error)
	CreateScheme(ctx context.Context, scheme *domain.Scheme) (*domain.Scheme, error)
	UpdateScheme(ctx context.Context, scheme *domain.Scheme) (*domain.Scheme, error)
	DeleteScheme(ctx context.Context, id uuid.UUID) error
}

type SchemeService interface {
	GetSchemeById(ctx context.Context, id uuid.UUID) (*domain.Scheme, error)
	ListSchemes(ctx context.Context) ([]domain.Scheme, error)
	CreateScheme(ctx context.Context, scheme *domain.Scheme) (*domain.Scheme, error)
	UpdateScheme(ctx context.Context, scheme *domain.Scheme) (*domain.Scheme, error)
	DeleteScheme(ctx context.Context, id uuid.UUID) error
	ListApplicantAvailableSchemes(ctx context.Context, applicantID uuid.UUID) ([]domain.Scheme, error)
}
