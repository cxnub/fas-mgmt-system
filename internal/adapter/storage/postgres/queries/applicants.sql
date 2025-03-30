-- db/query/applicants.sql

-- name: GetApplicant :one
-- Used for GET /api/applicants/{id}
SELECT * FROM applicants
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListApplicants :many
-- Used for GET /api/applicants
SELECT * FROM applicants
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: CreateApplicant :one
-- Used for POST /api/applicants
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
RETURNING *;

-- name: UpdateApplicant :one
-- Used for PUT /api/applicants/{id}
UPDATE applicants
SET
    name = $2,
    employment_status = $3,
    marital_status = $4,
    sex = $5,
    date_of_birth = $6
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteApplicant :exec
-- Used for DELETE /api/applicants/{id}
UPDATE applicants
SET
    deleted_at = now()
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetApplicantWithFamily :many
-- Used for getting an applicant with their family members
SELECT
    a.*,
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
WHERE a.id = $1 AND a.deleted_at IS NULL;