// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: applicants.sql

package pg

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createApplicant = `-- name: CreateApplicant :one
INSERT INTO applicants (
    id,
    created_at,
    name,
    employment_status,
    marital_status,
    sex,
    date_of_birth
) VALUES (
             gen_random_uuid(), now(), $1, $2, $3, $4, $5
         )
RETURNING id, created_at, updated_at, deleted_at, name, employment_status, marital_status, sex, date_of_birth
`

type CreateApplicantParams struct {
	Name             string
	EmploymentStatus EmploymentStatus
	MaritalStatus    MaritalStatus
	Sex              Sex
	DateOfBirth      pgtype.Date
}

// Used for POST /api/applicants
func (q *Queries) CreateApplicant(ctx context.Context, arg CreateApplicantParams) (Applicant, error) {
	row := q.db.QueryRow(ctx, createApplicant,
		arg.Name,
		arg.EmploymentStatus,
		arg.MaritalStatus,
		arg.Sex,
		arg.DateOfBirth,
	)
	var i Applicant
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
		&i.EmploymentStatus,
		&i.MaritalStatus,
		&i.Sex,
		&i.DateOfBirth,
	)
	return i, err
}

const deleteApplicant = `-- name: DeleteApplicant :exec
UPDATE applicants
SET
    deleted_at = now()
WHERE id = $1 AND deleted_at IS NULL
`

// Used for DELETE /api/applicants/{id}
func (q *Queries) DeleteApplicant(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteApplicant, id)
	return err
}

const getApplicant = `-- name: GetApplicant :one

SELECT id, created_at, updated_at, deleted_at, name, employment_status, marital_status, sex, date_of_birth FROM applicants
WHERE id = $1 AND deleted_at IS NULL
`

// db/query/applicants.sql
// Used for GET /api/applicants/{id}
func (q *Queries) GetApplicant(ctx context.Context, id uuid.UUID) (Applicant, error) {
	row := q.db.QueryRow(ctx, getApplicant, id)
	var i Applicant
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
		&i.EmploymentStatus,
		&i.MaritalStatus,
		&i.Sex,
		&i.DateOfBirth,
	)
	return i, err
}

const getApplicantWithFamily = `-- name: GetApplicantWithFamily :many
SELECT
    a.id, a.created_at, a.updated_at, a.deleted_at, a.name, a.employment_status, a.marital_status, a.sex, a.date_of_birth,
    r.relationship_type,
    family.id as family_member_id,
    family.name as family_member_name,
    family.employment_status as family_member_employment_status,
    family.marital_status as family_member_marital_status,
    family.sex as family_member_sex,
    family.date_of_birth as family_member_date_of_birth
FROM applicants a
         LEFT JOIN relationships r ON a.id = r.applicant_a_id AND r.deleted_at IS NULL
         LEFT JOIN applicants family ON r.applicant_b_id = family.id AND family.deleted_at IS NULL
WHERE a.id = $1 AND a.deleted_at IS NULL
`

type GetApplicantWithFamilyRow struct {
	ID                           uuid.UUID
	CreatedAt                    pgtype.Timestamp
	UpdatedAt                    pgtype.Timestamp
	DeletedAt                    pgtype.Timestamp
	Name                         string
	EmploymentStatus             EmploymentStatus
	MaritalStatus                MaritalStatus
	Sex                          Sex
	DateOfBirth                  pgtype.Date
	RelationshipType             NullRelationshipType
	FamilyMemberID               pgtype.UUID
	FamilyMemberName             pgtype.Text
	FamilyMemberEmploymentStatus NullEmploymentStatus
	FamilyMemberMaritalStatus    NullMaritalStatus
	FamilyMemberSex              NullSex
	FamilyMemberDateOfBirth      pgtype.Date
}

// Used for getting an applicant with their family members
func (q *Queries) GetApplicantWithFamily(ctx context.Context, id uuid.UUID) ([]GetApplicantWithFamilyRow, error) {
	rows, err := q.db.Query(ctx, getApplicantWithFamily, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetApplicantWithFamilyRow
	for rows.Next() {
		var i GetApplicantWithFamilyRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Name,
			&i.EmploymentStatus,
			&i.MaritalStatus,
			&i.Sex,
			&i.DateOfBirth,
			&i.RelationshipType,
			&i.FamilyMemberID,
			&i.FamilyMemberName,
			&i.FamilyMemberEmploymentStatus,
			&i.FamilyMemberMaritalStatus,
			&i.FamilyMemberSex,
			&i.FamilyMemberDateOfBirth,
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

const listApplicants = `-- name: ListApplicants :many
SELECT id, created_at, updated_at, deleted_at, name, employment_status, marital_status, sex, date_of_birth FROM applicants
WHERE deleted_at IS NULL
ORDER BY created_at DESC
`

// Used for GET /api/applicants
func (q *Queries) ListApplicants(ctx context.Context) ([]Applicant, error) {
	rows, err := q.db.Query(ctx, listApplicants)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Applicant
	for rows.Next() {
		var i Applicant
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Name,
			&i.EmploymentStatus,
			&i.MaritalStatus,
			&i.Sex,
			&i.DateOfBirth,
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

const updateApplicant = `-- name: UpdateApplicant :one
UPDATE applicants
SET
    name = $2,
    employment_status = $3,
    marital_status = $4,
    sex = $5,
    date_of_birth = $6
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, created_at, updated_at, deleted_at, name, employment_status, marital_status, sex, date_of_birth
`

type UpdateApplicantParams struct {
	ID               uuid.UUID
	Name             string
	EmploymentStatus EmploymentStatus
	MaritalStatus    MaritalStatus
	Sex              Sex
	DateOfBirth      pgtype.Date
}

// Used for PUT /api/applicants/{id}
func (q *Queries) UpdateApplicant(ctx context.Context, arg UpdateApplicantParams) (Applicant, error) {
	row := q.db.QueryRow(ctx, updateApplicant,
		arg.ID,
		arg.Name,
		arg.EmploymentStatus,
		arg.MaritalStatus,
		arg.Sex,
		arg.DateOfBirth,
	)
	var i Applicant
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
		&i.EmploymentStatus,
		&i.MaritalStatus,
		&i.Sex,
		&i.DateOfBirth,
	)
	return i, err
}
