package service

import (
	"context"
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/cxnub/fas-mgmt-system/internal/core/port"
	"github.com/google/uuid"
)

type ApplicantService struct {
	port.ApplicantRepository
}

func NewApplicantService(repo port.ApplicantRepository) *ApplicantService {
	return &ApplicantService{repo}
}
func (s *ApplicantService) GetApplicantById(ctx context.Context, id uuid.UUID) (*domain.Applicant, error) {
	return s.ApplicantRepository.GetApplicantById(ctx, id)
}

func (s *ApplicantService) ListApplicants(ctx context.Context) ([]domain.Applicant, error) {
	return s.ApplicantRepository.ListApplicants(ctx)
}

func (s *ApplicantService) CreateApplicant(ctx context.Context, applicant *domain.Applicant) (*domain.Applicant, error) {
	return s.ApplicantRepository.CreateApplicant(ctx, applicant)
}

func (s *ApplicantService) UpdateApplicant(ctx context.Context, applicant *domain.Applicant) (*domain.Applicant, error) {
	return s.ApplicantRepository.UpdateApplicant(ctx, applicant)
}

func (s *ApplicantService) DeleteApplicant(ctx context.Context, id uuid.UUID) error {
	return s.ApplicantRepository.DeleteApplicant(ctx, id)
}
