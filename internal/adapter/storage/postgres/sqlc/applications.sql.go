// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: applications.sql

package pg

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createApplication = `-- name: CreateApplication :one
INSERT INTO applications (
    id,
    created_at,
    applicant_id,
    scheme_id
) VALUES (
             gen_random_uuid(), now(), $1, $2
         )
RETURNING id, created_at, updated_at, deleted_at, applicant_id, scheme_id
`

type CreateApplicationParams struct {
	ApplicantID uuid.UUID
	SchemeID    uuid.UUID
}

// Used for POST /api/applications
func (q *Queries) CreateApplication(ctx context.Context, arg CreateApplicationParams) (Application, error) {
	row := q.db.QueryRow(ctx, createApplication, arg.ApplicantID, arg.SchemeID)
	var i Application
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.ApplicantID,
		&i.SchemeID,
	)
	return i, err
}

const deleteApplication = `-- name: DeleteApplication :exec
UPDATE applications
SET
    deleted_at = now()
WHERE id = $1 AND deleted_at IS NULL
`

// Used for DELETE /api/applications/{id}
func (q *Queries) DeleteApplication(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteApplication, id)
	return err
}

const getApplication = `-- name: GetApplication :one

SELECT id, created_at, updated_at, deleted_at, applicant_id, scheme_id FROM applications
WHERE id = $1 AND deleted_at IS NULL
`

// db/query/applications.sql
// Used for GET /api/applications/{id}
func (q *Queries) GetApplication(ctx context.Context, id uuid.UUID) (Application, error) {
	row := q.db.QueryRow(ctx, getApplication, id)
	var i Application
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.ApplicantID,
		&i.SchemeID,
	)
	return i, err
}

const getApplicationsByApplicant = `-- name: GetApplicationsByApplicant :many
SELECT id, created_at, updated_at, deleted_at, applicant_id, scheme_id FROM applications
WHERE applicant_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
`

// Used for getting applications for a specific applicant
func (q *Queries) GetApplicationsByApplicant(ctx context.Context, applicantID uuid.UUID) ([]Application, error) {
	rows, err := q.db.Query(ctx, getApplicationsByApplicant, applicantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Application
	for rows.Next() {
		var i Application
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.ApplicantID,
			&i.SchemeID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getApplicationsWithDetails = `-- name: GetApplicationsWithDetails :many
SELECT
    app.id, app.created_at, app.updated_at, app.deleted_at, app.applicant_id, app.scheme_id,
    a.name as applicant_name,
    a.employment_status as applicant_employment_status,
    s.name as scheme_name
FROM applications app
         JOIN applicants a ON app.applicant_id = a.id AND a.deleted_at IS NULL
         JOIN schemes s ON app.scheme_id = s.id AND s.deleted_at IS NULL
WHERE app.deleted_at IS NULL
ORDER BY app.created_at DESC
LIMIT $1 OFFSET $2
`

type GetApplicationsWithDetailsParams struct {
	Limit  int32
	Offset int32
}

type GetApplicationsWithDetailsRow struct {
	ID                        uuid.UUID
	CreatedAt                 pgtype.Timestamp
	UpdatedAt                 pgtype.Timestamp
	DeletedAt                 pgtype.Timestamp
	ApplicantID               uuid.UUID
	SchemeID                  uuid.UUID
	ApplicantName             string
	ApplicantEmploymentStatus EmploymentStatus
	SchemeName                string
}

// Used for getting applications with applicant and scheme details
func (q *Queries) GetApplicationsWithDetails(ctx context.Context, arg GetApplicationsWithDetailsParams) ([]GetApplicationsWithDetailsRow, error) {
	rows, err := q.db.Query(ctx, getApplicationsWithDetails, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetApplicationsWithDetailsRow
	for rows.Next() {
		var i GetApplicationsWithDetailsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.ApplicantID,
			&i.SchemeID,
			&i.ApplicantName,
			&i.ApplicantEmploymentStatus,
			&i.SchemeName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listApplications = `-- name: ListApplications :many
SELECT id, created_at, updated_at, deleted_at, applicant_id, scheme_id FROM applications
WHERE deleted_at IS NULL
ORDER BY created_at DESC
`

// Used for GET /api/applications
func (q *Queries) ListApplications(ctx context.Context) ([]Application, error) {
	rows, err := q.db.Query(ctx, listApplications)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Application
	for rows.Next() {
		var i Application
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.ApplicantID,
			&i.SchemeID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateApplication = `-- name: UpdateApplication :one
UPDATE applications
SET
    applicant_id = $2,
    scheme_id = $3
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, created_at, updated_at, deleted_at, applicant_id, scheme_id
`

type UpdateApplicationParams struct {
	ID          uuid.UUID
	ApplicantID uuid.UUID
	SchemeID    uuid.UUID
}

// Used for PUT /api/applications/{id}
func (q *Queries) UpdateApplication(ctx context.Context, arg UpdateApplicationParams) (Application, error) {
	row := q.db.QueryRow(ctx, updateApplication, arg.ID, arg.ApplicantID, arg.SchemeID)
	var i Application
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.ApplicantID,
		&i.SchemeID,
	)
	return i, err
}
