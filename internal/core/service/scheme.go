package service

import (
	"context"
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/cxnub/fas-mgmt-system/internal/core/port"
	"github.com/cxnub/fas-mgmt-system/internal/core/util"
	"github.com/google/uuid"
)

type SchemeService struct {
	port.SchemeRepository
	port.ApplicantRepository
}

func NewSchemeService(sr port.SchemeRepository, ar port.ApplicantRepository) *SchemeService {
	return &SchemeService{sr, ar}
}

func (s *SchemeService) GetSchemeById(ctx context.Context, id uuid.UUID) (*domain.Scheme, error) {
	return s.SchemeRepository.GetSchemeById(ctx, id)
}

func (s *SchemeService) ListSchemes(ctx context.Context) ([]domain.Scheme, error) {
	return s.SchemeRepository.ListSchemes(ctx)
}

func (s *SchemeService) CreateScheme(ctx context.Context, scheme *domain.Scheme) (*domain.Scheme, error) {
	return s.SchemeRepository.CreateScheme(ctx, scheme)
}

func (s *SchemeService) UpdateScheme(ctx context.Context, scheme *domain.Scheme) (*domain.Scheme, error) {
	return s.SchemeRepository.UpdateScheme(ctx, scheme)
}

func (s *SchemeService) DeleteScheme(ctx context.Context, id uuid.UUID) error {
	return s.SchemeRepository.DeleteScheme(ctx, id)
}

func (s *SchemeService) ListApplicantAvailableSchemes(ctx context.Context, applicantID uuid.UUID) ([]domain.Scheme, error) {
	// Get all schemes
	schemes, err := s.SchemeRepository.ListSchemes(ctx)

	if err != nil {
		return nil, err
	}

	// Get applicant
	applicant, err := s.ApplicantRepository.GetApplicantById(ctx, applicantID)

	if err != nil {
		return nil, err
	}

	// Get applicant relations
	relationships, err := s.ApplicantRepository.GetApplicantFamily(ctx, applicantID)

	if err != nil {
		return nil, err
	}

	result := make([]domain.Scheme, 0)

	for _, scheme := range schemes {
		if util.CheckSchemeEligibility(scheme, applicant, relationships) {
			result = append(result, scheme)
		}
	}

	return result, nil
}
