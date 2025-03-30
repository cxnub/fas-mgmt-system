package repository

import (
	"context"
	sql2 "database/sql"
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

// fetchBenefitsForScheme fetches benefits for a specific scheme
func (r *SchemeRepository) fetchBenefitsForSchemes(ctx context.Context, schemeMap map[uuid.UUID]*domain.Scheme, schemeID *uuid.UUID) error {
	var benefitsQuery interface {
		ToSql() (string, []interface{}, error)
	}

	if schemeID != nil {
		// Query for a specific scheme
		benefitsQuery = r.db.QueryBuilder.
			Select("id", "name", "amount", "scheme_id").
			From("benefits").
			Where("scheme_id = ? AND deleted_at IS NULL", *schemeID)
	} else {
		// Query for all schemes
		benefitsQuery = r.db.QueryBuilder.
			Select("id", "name", "amount", "scheme_id").
			From("benefits").
			Where("deleted_at IS NULL")
	}

	sql, args, err := benefitsQuery.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build benefits query: %w", err)
	}

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var benefit domain.Benefit
		var schemeID uuid.UUID
		var amount sql2.NullFloat64

		err = rows.Scan(&benefit.ID, &benefit.Name, &amount, &schemeID)
		if err != nil {
			return err
		}

		benefit.Amount = &amount.Float64

		if scheme, exists := schemeMap[schemeID]; exists {
			*scheme.Benefits = append(*scheme.Benefits, benefit)
		}
	}

	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}

// fetchCriteriaForSchemes fetches criteria for specified schemes
func (r *SchemeRepository) fetchCriteriaForSchemes(ctx context.Context, schemeMap map[uuid.UUID]*domain.Scheme, schemeID *uuid.UUID) error {
	var criteriaQuery interface {
		ToSql() (string, []interface{}, error)
	}

	if schemeID != nil {
		// Query for a specific scheme
		criteriaQuery = r.db.QueryBuilder.
			Select("id", "name", "value", "scheme_id").
			From("scheme_criteria").
			Where("scheme_id = ? AND deleted_at IS NULL", *schemeID)
	} else {
		// Query for all schemes
		criteriaQuery = r.db.QueryBuilder.
			Select("id", "name", "value", "scheme_id").
			From("scheme_criteria").
			Where("deleted_at IS NULL")
	}

	sql, args, err := criteriaQuery.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build criteria query: %w", err)
	}

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var criterion domain.SchemeCriteria
		var schemeID uuid.UUID

		err = rows.Scan(&criterion.ID, &criterion.Name, &criterion.Value, &schemeID)
		if err != nil {
			return err
		}

		if scheme, exists := schemeMap[schemeID]; exists {
			*scheme.Criteria = append(*scheme.Criteria, criterion)
		}
	}

	if rows.Err() != nil {
		return rows.Err()
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
