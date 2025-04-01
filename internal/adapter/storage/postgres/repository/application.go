package repository

import (
	"context"
	"errors"
	"github.com/cxnub/fas-mgmt-system/internal/adapter/storage/postgres"
	pg "github.com/cxnub/fas-mgmt-system/internal/adapter/storage/postgres/sqlc"
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// ApplicationRepository provides methods for managing applications in the database. It encapsulates database access logic.
type ApplicationRepository struct {
	db *postgres.DB
	q  pg.Querier
}

// NewApplicationRepository creates a new instance of ApplicationRepository with the provided database and querier dependencies.
func NewApplicationRepository(db *postgres.DB, q pg.Querier) *ApplicationRepository {
	return &ApplicationRepository{db: db, q: q}
}

// ListApplications retrieves a list of all applications from the database and converts them into domain entities.
func (r *ApplicationRepository) ListApplications(ctx context.Context) ([]domain.Application, error) {
	dbApplications, err := r.q.ListApplications(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ApplicationNotFoundError
		}

		return nil, err
	}

	dbApplicationsEntities := make([]domain.Application, len(dbApplications))
	for i, dbApplication := range dbApplications {
		dbApplicationsEntities[i] = *dbApplication.ToEntity()
	}

	return dbApplicationsEntities, nil
}

// GetApplicationById retrieves an Application entity by its unique identifier from the database.
// Returns domain.ApplicationNotFoundError if no matching record is found.
func (r *ApplicationRepository) GetApplicationById(ctx context.Context, id uuid.UUID) (*domain.Application, error) {
	dbApplication, err := r.q.GetApplication(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ApplicationNotFoundError
		}
	}

	return dbApplication.ToEntity(), nil
}

// CreateApplication inserts a new application into the database and returns the created application entity or an error.
func (r *ApplicationRepository) CreateApplication(ctx context.Context, application *domain.Application) (*domain.Application, error) {
	dbApplication := pg.ApplicationFromEntity(application)

	params := pg.CreateApplicationParams{
		ApplicantID: dbApplication.ApplicantID,
		SchemeID:    dbApplication.SchemeID,
	}

	a, err := r.q.CreateApplication(ctx, params)
	if err != nil {
		return nil, err
	}

	return a.ToEntity(), nil
}

// UpdateApplication updates an existing application in the database with the provided fields and returns the updated application.
func (r *ApplicationRepository) UpdateApplication(ctx context.Context, application *domain.Application) (*domain.Application, error) {

	setFields := false

	query := r.db.QueryBuilder.Update("applications")

	if application.SchemeID != nil {
		query = query.Set("scheme_id", application.SchemeID)
		setFields = true
	}

	if application.SchemeID != nil {
		query = query.Set("scheme_id", application.SchemeID)
		setFields = true
	}

	if !setFields {
		return nil, domain.NoUpdateFieldsError
	}

	query = query.Where("id = ?", application.ID)

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	a, err := r.q.GetApplication(ctx, *application.ID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ApplicantNotFoundError
		}
		return nil, err
	}

	return a.ToEntity(), nil
}

// DeleteApplication removes an application from the database by its unique identifier. Returns an error if delete fails.
func (r *ApplicationRepository) DeleteApplication(ctx context.Context, id uuid.UUID) error {
	err := r.q.DeleteApplication(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ApplicationNotFoundError
		}
		return err
	}

	return nil
}
