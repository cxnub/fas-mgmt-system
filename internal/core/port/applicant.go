package port

import (
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type ApplicantRepository interface {
	GetApplicantById(ctx context.Context, id uuid.UUID) (*domain.Applicant, error)
	ListApplicants(ctx context.Context) ([]domain.Applicant, error)
	CreateApplicant(ctx context.Context, applicant *domain.Applicant) (*domain.Applicant, error)
	UpdateApplicant(ctx context.Context, applicant *domain.Applicant) (*domain.Applicant, error)
	DeleteApplicant(ctx context.Context, id uuid.UUID) error
	GetApplicantFamily(ctx context.Context, id uuid.UUID) (map[domain.RelationshipType]*domain.Applicant, error)
}

type ApplicantService interface {
	GetApplicantById(ctx context.Context, id uuid.UUID) (*domain.Applicant, error)
	ListApplicants(ctx context.Context) ([]domain.Applicant, error)
	CreateApplicant(ctx context.Context, applicant *domain.Applicant) (*domain.Applicant, error)
	UpdateApplicant(ctx context.Context, applicant *domain.Applicant) (*domain.Applicant, error)
	DeleteApplicant(ctx context.Context, id uuid.UUID) error
}
