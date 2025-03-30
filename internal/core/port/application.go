package port

import (
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type ApplicationRepository interface {
	GetApplicationById(ctx context.Context, id uuid.UUID) (*domain.Application, error)
	ListApplications(ctx context.Context) ([]domain.Application, error)
	CreateApplication(ctx context.Context, application *domain.Application) (*domain.Application, error)
	UpdateApplication(ctx context.Context, application *domain.Application) (*domain.Application, error)
	DeleteApplication(ctx context.Context, id uuid.UUID) error
}

type ApplicationService interface {
	GetApplicationById(ctx context.Context, id uuid.UUID) (*domain.Application, error)
	ListApplications(ctx context.Context) ([]domain.Application, error)
	CreateApplication(ctx context.Context, application *domain.Application) (*domain.Application, error)
	UpdateApplication(ctx context.Context, application *domain.Application) (*domain.Application, error)
	DeleteApplication(ctx context.Context, id uuid.UUID) error
}
