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
	return s.SchemeRepository.GetSchemeByID(ctx, id)
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

func (s *SchemeService) AddSchemeBenefit(ctx context.Context, benefit *domain.Benefit) (newBenefit *domain.Benefit, err error) {
	// Check if scheme exists
	_, err = s.SchemeRepository.GetSchemeByID(ctx, *benefit.SchemeID)
	if err != nil {
		return nil, err
	}

	return s.SchemeRepository.AddSchemeBenefit(ctx, benefit)
}

func (s *SchemeService) UpdateSchemeBenefit(ctx context.Context, benefit *domain.Benefit) (newBenefit *domain.Benefit, err error) {
	// Check if scheme exists
	_, err = s.SchemeRepository.GetSchemeByID(ctx, *benefit.SchemeID)
	if err != nil {
		return nil, err
	}

	// Check if benefit exists
	_, err = s.SchemeRepository.GetBenefitByID(ctx, *benefit.ID)
	if err != nil {
		return nil, err
	}

	return s.SchemeRepository.UpdateSchemeBenefit(ctx, benefit)
}

func (s *SchemeService) DeleteSchemeBenefit(ctx context.Context, benefitID uuid.UUID) error {
	// Check if benefit exists
	_, err := s.SchemeRepository.GetBenefitByID(ctx, benefitID)
	if err != nil {
		return err
	}

	return s.SchemeRepository.DeleteSchemeBenefit(ctx, benefitID)
}

func (s *SchemeService) AddSchemeCriteria(ctx context.Context, criteria *domain.SchemeCriteria) (newCriteria *domain.SchemeCriteria, err error) {
	// Check if criteria is valid
	invalidCriteriaErr := util.IsValidCriteria(criteria)
	if invalidCriteriaErr != nil {
		return nil, *invalidCriteriaErr
	}

	// Check if scheme exists
	_, err = s.SchemeRepository.GetSchemeByID(ctx, *criteria.SchemeID)
	if err != nil {
		return nil, err
	}

	return s.SchemeRepository.AddSchemeCriteria(ctx, criteria)
}

func (s *SchemeService) UpdateSchemeCriteria(ctx context.Context, criteria *domain.SchemeCriteria) (newCriteria *domain.SchemeCriteria, err error) {
	// Check if criteria is valid
	invalidCriteriaErr := util.IsValidCriteria(criteria)
	if invalidCriteriaErr != nil {
		return nil, *invalidCriteriaErr
	}

	// Check if scheme exists
	_, err = s.SchemeRepository.GetSchemeByID(ctx, *criteria.SchemeID)
	if err != nil {
		return nil, err
	}

	// Check if criteria exists
	_, err = s.SchemeRepository.GetSchemeCriteriaByID(ctx, *criteria.ID)
	if err != nil {
		return nil, err
	}

	return s.SchemeRepository.UpdateSchemeCriteria(ctx, criteria)
}

func (s *SchemeService) DeleteSchemeCriteria(ctx context.Context, criteriaID uuid.UUID) error {
	// Check if criteria exists
	_, err := s.SchemeRepository.GetSchemeCriteriaByID(ctx, criteriaID)
	if err != nil {
		return err
	}

	return s.SchemeRepository.DeleteSchemeCriteria(ctx, criteriaID)
}
