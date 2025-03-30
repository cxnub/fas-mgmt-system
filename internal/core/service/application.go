package service

import (
	"context"
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/cxnub/fas-mgmt-system/internal/core/port"
	"github.com/cxnub/fas-mgmt-system/internal/core/util"
	"github.com/google/uuid"
)

type ApplicationService struct {
	port.ApplicationRepository
	port.ApplicantRepository
	port.SchemeRepository
}

func NewApplicationService(applicationRepo port.ApplicationRepository, applicantRepo port.ApplicantRepository, schemeRepo port.SchemeRepository) *ApplicationService {
	return &ApplicationService{applicationRepo, applicantRepo, schemeRepo}
}
func (s *ApplicationService) GetApplicationById(ctx context.Context, id uuid.UUID) (*domain.Application, error) {
	return s.ApplicationRepository.GetApplicationById(ctx, id)
}

func (s *ApplicationService) ListApplications(ctx context.Context) ([]domain.Application, error) {
	return s.ApplicationRepository.ListApplications(ctx)
}

func (s *ApplicationService) checkApplicationValidity(ctx context.Context, application *domain.Application) error {
	applicant, err := s.ApplicantRepository.GetApplicantById(ctx, *application.ApplicantID)
	if err != nil {
		return err
	}

	scheme, err := s.SchemeRepository.GetSchemeByID(ctx, *application.SchemeID)
	if err != nil {
		return err
	}

	// Get applicant relations
	relationships, err := s.ApplicantRepository.GetApplicantFamily(ctx, *applicant.ID)

	if err != nil {
		return err
	}

	// Check applicant eligibility
	if !util.CheckSchemeEligibility(*scheme, applicant, relationships) {
		return domain.SchemeNotEligibleError
	}

	return nil
}

func (s *ApplicationService) CreateApplication(ctx context.Context, application *domain.Application) (*domain.Application, error) {
	if err := s.checkApplicationValidity(ctx, application); err != nil {
		return nil, err
	}

	return s.ApplicationRepository.CreateApplication(ctx, application)
}

func (s *ApplicationService) UpdateApplication(ctx context.Context, application *domain.Application) (*domain.Application, error) {
	if err := s.checkApplicationValidity(ctx, application); err != nil {
		return nil, err
	}

	return s.ApplicationRepository.UpdateApplication(ctx, application)
}

func (s *ApplicationService) DeleteApplication(ctx context.Context, id uuid.UUID) error {
	return s.ApplicationRepository.DeleteApplication(ctx, id)
}
