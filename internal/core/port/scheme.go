package port

import (
	"context"
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/google/uuid"
)

type SchemeRepository interface {
	GetSchemeByID(ctx context.Context, id uuid.UUID) (*domain.Scheme, error)
	ListSchemes(ctx context.Context) ([]domain.Scheme, error)
	CreateScheme(ctx context.Context, scheme *domain.Scheme) (*domain.Scheme, error)
	UpdateScheme(ctx context.Context, scheme *domain.Scheme) (*domain.Scheme, error)
	DeleteScheme(ctx context.Context, id uuid.UUID) error

	GetBenefitByID(ctx context.Context, benefitID uuid.UUID) (*domain.Benefit, error)
	AddSchemeBenefit(ctx context.Context, benefit *domain.Benefit) (newBenefit *domain.Benefit, err error)
	UpdateSchemeBenefit(ctx context.Context, benefit *domain.Benefit) (newBenefit *domain.Benefit, err error)
	DeleteSchemeBenefit(ctx context.Context, benefitID uuid.UUID) error

	GetSchemeCriteriaByID(ctx context.Context, criteriaID uuid.UUID) (*domain.SchemeCriteria, error)
	AddSchemeCriteria(ctx context.Context, criteria *domain.SchemeCriteria) (newCriteria *domain.SchemeCriteria, err error)
	UpdateSchemeCriteria(ctx context.Context, criteria *domain.SchemeCriteria) (newCriteria *domain.SchemeCriteria, err error)
	DeleteSchemeCriteria(ctx context.Context, criteriaID uuid.UUID) error
}

type SchemeService interface {
	GetSchemeByID(ctx context.Context, id uuid.UUID) (*domain.Scheme, error)
	ListSchemes(ctx context.Context) ([]domain.Scheme, error)
	CreateScheme(ctx context.Context, scheme *domain.Scheme) (*domain.Scheme, error)
	UpdateScheme(ctx context.Context, scheme *domain.Scheme) (*domain.Scheme, error)
	DeleteScheme(ctx context.Context, id uuid.UUID) error
	ListApplicantAvailableSchemes(ctx context.Context, applicantID uuid.UUID) ([]domain.Scheme, error)

	AddSchemeBenefit(ctx context.Context, benefit *domain.Benefit) (newBenefit *domain.Benefit, err error)
	UpdateSchemeBenefit(ctx context.Context, benefit *domain.Benefit) (newBenefit *domain.Benefit, err error)
	DeleteSchemeBenefit(ctx context.Context, benefitID uuid.UUID) error

	AddSchemeCriteria(ctx context.Context, criteria *domain.SchemeCriteria) (newCriteria *domain.SchemeCriteria, err error)
	UpdateSchemeCriteria(ctx context.Context, criteria *domain.SchemeCriteria) (newCriteria *domain.SchemeCriteria, err error)
	DeleteSchemeCriteria(ctx context.Context, criteriaID uuid.UUID) error
}
