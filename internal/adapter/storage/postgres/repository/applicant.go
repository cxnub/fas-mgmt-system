package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cxnub/fas-mgmt-system/internal/adapter/storage/postgres"
	pg "github.com/cxnub/fas-mgmt-system/internal/adapter/storage/postgres/sqlc"
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"time"
)

// ApplicantRepository provides methods for interacting with applicant-related data in the database.
// It uses postgres.DB for database operations and pg.Querier for application-specific queries.
type ApplicantRepository struct {
	db *postgres.DB
	q  pg.Querier
}

// NewApplicantRepository creates a new instance of ApplicantRepository using the provided database connection and querier.
func NewApplicantRepository(db *postgres.DB, q pg.Querier) *ApplicantRepository {
	return &ApplicantRepository{db: db, q: q}
}

// GetApplicantById retrieves an applicant by their unique identifier from the database.
func (r *ApplicantRepository) GetApplicantById(ctx context.Context, id uuid.UUID) (applicant *domain.Applicant, err error) {
	dbApplicant, err := r.q.GetApplicant(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ApplicantNotFoundError
		}

		return nil, err
	}

	return dbApplicant.ToEntity(), nil
}

// GetApplicantFamily retrieves an applicant's family members by the applicant's ID from the database.
func (r *ApplicantRepository) GetApplicantFamily(ctx context.Context, id uuid.UUID) (map[domain.RelationshipType]*domain.Applicant, error) {
	query := r.db.QueryBuilder.
		Select(
			"r.relationship_type",
			"family.id AS family_member_id",
			"family.name AS family_member_name",
			"family.employment_status AS family_member_employment_status",
			"family.marital_status AS family_member_marital_status",
			"family.sex AS family_member_sex",
			"family.date_of_birth AS family_member_date_of_birth",
		).
		From("relationships r").
		LeftJoin("applicants family ON r.applicant_b_id = family.id AND family.deleted_at IS NULL").
		Where(squirrel.Eq{"r.applicant_a_id": id}).
		Where("r.deleted_at IS NULL")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	family := make(map[domain.RelationshipType]*domain.Applicant)

	for rows.Next() {
		var relationshipType *string
		var familyMemberID *uuid.UUID
		var familyMemberName, familyMemberEmploymentStatus, familyMemberMaritalStatus, familyMemberSex *string
		var familyMemberDateOfBirth *time.Time

		err = rows.Scan(
			&relationshipType, &familyMemberID, &familyMemberName, &familyMemberEmploymentStatus,
			&familyMemberMaritalStatus, &familyMemberSex, &familyMemberDateOfBirth,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Add family member if exists
		if familyMemberID != nil && relationshipType != nil {
			familyMember := &domain.Applicant{
				ID:               familyMemberID,
				Name:             familyMemberName,
				EmploymentStatus: (*domain.EmploymentStatus)(familyMemberEmploymentStatus),
				MaritalStatus:    (*domain.MaritalStatus)(familyMemberMaritalStatus),
				Sex:              (*domain.Sex)(familyMemberSex),
				DateOfBirth:      familyMemberDateOfBirth,
			}
			family[domain.RelationshipType(*relationshipType)] = familyMember
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %w", err)
	}

	return family, nil
}

// ListApplicants retrieves all applicants from the database and converts them to the domain entity representation.
func (r *ApplicantRepository) ListApplicants(ctx context.Context) (applicants []domain.Applicant, err error) {
	dbApplicants, err := r.q.ListApplicants(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ApplicantNotFoundError
		}

		return nil, err
	}

	dbApplicantsEntities := make([]domain.Applicant, len(dbApplicants))
	for i, dbApplicant := range dbApplicants {
		dbApplicantsEntities[i] = *dbApplicant.ToEntity()
	}

	return dbApplicantsEntities, nil
}

// CreateApplicant inserts a new applicant into the database and returns the created applicant or an error if one occurs.
func (r *ApplicantRepository) CreateApplicant(ctx context.Context, applicant *domain.Applicant) (newApplicant *domain.Applicant, err error) {
	dbApplicant := pg.ApplicantFromEntity(applicant)

	params := pg.CreateApplicantParams{
		Name:             dbApplicant.Name,
		EmploymentStatus: dbApplicant.EmploymentStatus,
		MaritalStatus:    dbApplicant.MaritalStatus,
		Sex:              dbApplicant.Sex,
		DateOfBirth:      dbApplicant.DateOfBirth,
	}
	a, err := r.q.CreateApplicant(ctx, params)
	if err != nil {
		return nil, err
	}

	return a.ToEntity(), nil
}

// UpdateApplicant updates an existing applicant's details in the database and returns the updated applicant or an error.
func (r *ApplicantRepository) UpdateApplicant(ctx context.Context, applicant *domain.Applicant) (updatedApplicant *domain.Applicant, err error) {
	var updatedDbApplicant pg.Applicant

	if applicant == nil {
		return nil, fmt.Errorf("applicant cannot be nil")
	}

	if applicant.ID.String() == "" {
		return nil, fmt.Errorf("applicant id cannot be nil")
	}

	query := r.db.QueryBuilder.Update("applicants")

	setFields := false

	if applicant.Name != nil {
		query = query.Set("name", applicant.Name)
		setFields = true
	}

	if applicant.EmploymentStatus != nil {
		query = query.Set("employment_status", applicant.EmploymentStatus)
		setFields = true
	}

	if applicant.MaritalStatus != nil {
		query = query.Set("marital_status", applicant.MaritalStatus)
		setFields = true
	}

	if applicant.Sex != nil {
		query = query.Set("sex", applicant.Sex)
		setFields = true
	}

	if applicant.DateOfBirth != nil {
		query = query.Set("date_of_birth", applicant.DateOfBirth)
		setFields = true
	}

	if !setFields {
		return nil, domain.NoUpdateFieldsError
	}

	query = query.Where("id = ?", applicant.ID)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ApplicantNotFoundError
		}

		return nil, err
	}

	updatedDbApplicant, err = r.q.GetApplicant(ctx, *applicant.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ApplicantNotFoundError
		}

		return nil, err
	}

	return updatedDbApplicant.ToEntity(), nil
}

// DeleteApplicant deletes an applicant record from the database by their unique identifier. Returns an error if the operation fails.
func (r *ApplicantRepository) DeleteApplicant(ctx context.Context, id uuid.UUID) (err error) {
	err = r.q.DeleteApplicant(ctx, id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ApplicantNotFoundError
		}

		return err
	}

	return nil
}
