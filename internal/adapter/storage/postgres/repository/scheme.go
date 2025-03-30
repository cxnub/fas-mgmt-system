package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/cxnub/fas-mgmt-system/internal/adapter/storage/postgres"
	pg "github.com/cxnub/fas-mgmt-system/internal/adapter/storage/postgres/sqlc"
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type SchemeRepository struct {
	db *postgres.DB
	q  pg.Querier
}

// NewSchemeRepository creates a new instance of SchemeRepository using the provided database connection and querier.
func NewSchemeRepository(db *postgres.DB, q pg.Querier) *SchemeRepository {
	return &SchemeRepository{db: db, q: q}
}

// fetchBenefitsForSchemes fetches benefits for a specific scheme
func (r *SchemeRepository) fetchBenefitsForSchemes(ctx context.Context, schemeMap map[uuid.UUID]*domain.Scheme, schemeID *uuid.UUID) error {
	var err error
	var benefitArray []pg.Benefit

	if schemeID != nil {
		// Query for a specific benefit
		benefitArray, err = r.q.GetBenefitsByScheme(ctx, *schemeID)
	} else {
		// Query for all benefits
		benefitArray, err = r.q.ListBenefits(ctx)
	}

	if err != nil {
		return err
	}

	for _, benefit := range benefitArray {
		schemeBenefit := benefit.ToEntity()

		if scheme, exists := schemeMap[benefit.SchemeID]; exists {
			*scheme.Benefits = append(*scheme.Benefits, *schemeBenefit)
		}
	}

	return nil
}

// fetchCriteriaForSchemes fetches criteria for specified schemes
func (r *SchemeRepository) fetchCriteriaForSchemes(ctx context.Context, schemeMap map[uuid.UUID]*domain.Scheme, schemeID *uuid.UUID) error {
	var err error
	var criteriaArray []pg.SchemeCriterium

	if schemeID != nil {
		// Query for a specific scheme
		criteriaArray, err = r.q.GetSchemeCriteria(ctx, *schemeID)
	} else {
		// Query for all schemes
		criteriaArray, err = r.q.ListSchemeCriteria(ctx)
	}

	if err != nil {
		return err
	}

	for _, criteria := range criteriaArray {
		schemeCriteria := criteria.ToEntity()

		if scheme, exists := schemeMap[criteria.SchemeID]; exists {
			*scheme.Criteria = append(*scheme.Criteria, *schemeCriteria)
		}
	}

	return nil
}

// GetSchemeById retrieves a scheme by its ID, including its benefits and criteria, or returns an error if not found.
func (r *SchemeRepository) GetSchemeById(ctx context.Context, id uuid.UUID) (*domain.Scheme, error) {
	// Get the scheme by ID
	schemesQuery := r.db.QueryBuilder.
		Select("id", "name", "created_at", "updated_at").
		From("schemes").
		Where("id = ? AND deleted_at IS NULL", id)

	sql, args, err := schemesQuery.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build scheme query: %w", err)
	}

	row := r.db.QueryRow(ctx, sql, args...)
	var scheme domain.Scheme
	err = row.Scan(&scheme.ID, &scheme.Name, &scheme.CreatedAt, &scheme.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NotFoundError
		}
		return nil, err
	}
	scheme.Benefits = &[]domain.Benefit{}
	scheme.Criteria = &[]domain.SchemeCriteria{}

	// Create a temporary map with just this scheme
	schemeMap := map[uuid.UUID]*domain.Scheme{*scheme.ID: &scheme}

	// Fetch benefits for the scheme
	if err := r.fetchBenefitsForSchemes(ctx, schemeMap, scheme.ID); err != nil {
		return nil, err
	}

	// Fetch criteria for the scheme
	if err := r.fetchCriteriaForSchemes(ctx, schemeMap, scheme.ID); err != nil {
		return nil, err
	}

	return &scheme, nil
}

// ListSchemes retrieves a list of schemes available.
func (r *SchemeRepository) ListSchemes(ctx context.Context) ([]domain.Scheme, error) {
	// Get all schemes
	schemesQuery := r.db.QueryBuilder.
		Select("id", "name", "created_at", "updated_at").
		From("schemes").
		Where("deleted_at IS NULL")

	sql, args, err := schemesQuery.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build schemes query: %w", err)
	}

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NotFoundError
		}
		return nil, err
	}
	defer rows.Close()

	// Store schemes in a map
	schemesMap := make(map[uuid.UUID]*domain.Scheme)

	for rows.Next() {
		var scheme domain.Scheme
		err = rows.Scan(&scheme.ID, &scheme.Name, &scheme.CreatedAt, &scheme.UpdatedAt)
		if err != nil {
			return nil, err
		}
		scheme.Benefits = &[]domain.Benefit{}
		scheme.Criteria = &[]domain.SchemeCriteria{}
		schemesMap[*scheme.ID] = &scheme
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	// Fetch benefits and map them to schemes
	if err := r.fetchBenefitsForSchemes(ctx, schemesMap, nil); err != nil {
		return nil, err
	}

	// Fetch criteria and map them to schemes
	if err := r.fetchCriteriaForSchemes(ctx, schemesMap, nil); err != nil {
		return nil, err
	}

	// Convert map to slice
	schemes := make([]domain.Scheme, 0, len(schemesMap))
	for _, scheme := range schemesMap {
		schemes = append(schemes, *scheme)
	}

	return schemes, nil
}

// CreateScheme inserts a new scheme into the database and returns the created scheme or an error if one occurs.
func (r *SchemeRepository) CreateScheme(ctx context.Context, scheme *domain.Scheme) (newScheme *domain.Scheme, err error) {
	dbScheme := pg.SchemeFromEntity(scheme)

	a, err := r.q.CreateScheme(ctx, dbScheme.Name)
	if err != nil {
		return nil, err
	}

	return a.ToEntity(), nil
}

// AddSchemeBenefit inserts a new benefit into a specific scheme and returns the created benefit or an error if one occurs.
func (r *SchemeRepository) AddSchemeBenefit(ctx context.Context, schemeID uuid.UUID, benefit *domain.Benefit) (newBenefit *domain.Benefit, err error) {
	dbBenefit := pg.BenefitFromEntity(benefit)

	params := pg.CreateBenefitParams{
		SchemeID: dbBenefit.SchemeID,
		Name:     dbBenefit.Name,
		Amount:   dbBenefit.Amount,
	}

	b, err := r.q.CreateBenefit(ctx, params)
	if err != nil {
		return nil, err
	}

	return b.ToEntity(), nil
}

// AddSchemeCriteria adds a new criteria to a specific scheme and returns the created criteria or an error if one occurs.
func (r *SchemeRepository) AddSchemeCriteria(ctx context.Context, schemeID uuid.UUID, criteria *domain.SchemeCriteria) (newCriteria *domain.SchemeCriteria, err error) {
	dbSchemeCriteria := pg.SchemeCriteriumFromEntity(criteria)

	params := pg.CreateSchemeCriteriaParams{
		SchemeID: dbSchemeCriteria.SchemeID,
		Name:     dbSchemeCriteria.Name,
		Value:    dbSchemeCriteria.Value,
	}

	c, err := r.q.CreateSchemeCriteria(ctx, params)
	if err != nil {
		return nil, err
	}

	return c.ToEntity(), nil
}

// UpdateScheme updates an existing scheme's details in the database and returns the updated scheme or an error.
func (r *SchemeRepository) UpdateScheme(ctx context.Context, scheme *domain.Scheme) (updatedScheme *domain.Scheme, err error) {
	var updatedDbScheme pg.Scheme

	if scheme == nil {
		return nil, fmt.Errorf("scheme cannot be nil")
	}

	if scheme.ID.String() == "" {
		return nil, fmt.Errorf("scheme id cannot be nil")
	}

	query := r.db.QueryBuilder.Update("schemes")

	setFields := false

	if scheme.Name != nil {
		query = query.Set("name", scheme.Name)
		setFields = true
	}

	if !setFields {
		return nil, domain.NoUpdateFieldsError
	}

	query = query.Where("id = ?", scheme.ID)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.SchemeNotFoundError
		}

		return nil, err
	}

	updatedDbScheme, err = r.q.GetScheme(ctx, *scheme.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.SchemeNotFoundError
		}

		return nil, err
	}

	return updatedDbScheme.ToEntity(), nil
}

// DeleteScheme deletes a scheme record from the database by their unique identifier. Returns an error if the operation fails.
func (r *SchemeRepository) DeleteScheme(ctx context.Context, id uuid.UUID) (err error) {
	err = r.q.DeleteScheme(ctx, id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.SchemeNotFoundError
		}

		return err
	}

	return nil
}
